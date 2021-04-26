package internal

import (
	"fmt"
	"math"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
)

// getLastDeviceInfo 获取指定设备的最近一次的有效消息记录.
func getLastDeviceInfo(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (msg *trashcan.Trashcan, err error) {
	msg = &trashcan.Trashcan{}
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline := now.Add(d2)
	whereArgs := "device_sn = ? and created_at > ?"
	err = db.Where(whereArgs, deviceSN, deadline).Order("id desc").First(msg).Error
	return
}

// GetDevices 获取垃圾箱设备列表, 每个成员包含percent字段, 表示此设备当前的使用率
// 可以按照使用率(percent字段)排序, 使用表连接查询实现.
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (trashcanDeviceList *protos.TrashcanDeviceList, err error) {
	trashcanDeviceList = &protos.TrashcanDeviceList{
		List: []*protos.TrashcanDevice{},
	}
	trashcanRecords := []*trashcan.Trashcan{}
	// 查询所有设备最近的消息, 条件查询语句在后面拼接. 注意msg_type指定为1, 2, 其他类型的消息不包含有效数据.
	sqlStr := `
		select main_tbl.* from trashcans as main_tbl 
		inner join (
			select device_sn, max(created_at) as created_at from trashcans where msg_type in (1, 2) group by device_sn
		) as b on main_tbl.device_sn = b.device_sn and main_tbl.created_at = b.created_at %s 
	`
	countSQLStr := `
		select count(main_tbl.id) from trashcans as main_tbl 
		inner join (
			select device_sn, max(created_at) as created_at from trashcans where msg_type in (1, 2) group by device_sn
		) as b on main_tbl.device_sn = b.device_sn and main_tbl.created_at = b.created_at %s 
	`

	whereStr := " where 1 = 1 "
	if req.CustomerID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.customer_id = %d ", req.CustomerID)
	}
	if req.GroupID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.group_id = %d ", req.GroupID)
	}
	if req.SerialNumber != "" {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.device_sn = '%s' ", req.SerialNumber)
	}

	countSQLStr = fmt.Sprintf(countSQLStr, whereStr)
	log.Debugf("exec sql string: %s", countSQLStr)
	var count uint64
	err = db.Raw(countSQLStr).Count(&count).Error
	if err != nil {
		log.Errorf("query device message count failed in GetDevices(): %s", err.Error())
		return
	}
	sqlStr = fmt.Sprintf(sqlStr, whereStr)
	sqlStr = pagination.BuildPaginationQueryInString(sqlStr, req.Pagination)
	log.Debugf("execute sql in GetDevices(): %s", sqlStr)
	err = db.Raw(sqlStr).Scan(&trashcanRecords).Error
	if err != nil {
		log.Errorf("query device latest messages failed in GetDevices(): %s", err.Error())
		return
	}

	for _, trashcanRecord := range trashcanRecords {
		trashcanDevice, err := getDeviceInfo(db, log, trashcanRecord.DeviceSN)
		if err != nil {
			log.Errorf("get trashcan device failed in GetDevices(): %s", err.Error())
			return trashcanDeviceList, nil
		}
		// 先按桶高1.25m计算
		trashcanDevice.Percent = float64(int(math.Abs((conf.TrashcanHeight - trashcanRecord.Percent) / conf.TrashcanHeight * 100)))
		trashcanDeviceList.List = append(trashcanDeviceList.List, trashcanDevice)
	}
	trashcanDeviceList.Count = count
	trashcanDeviceList.TotalCount = count
	trashcanDeviceList.CurrentPage = req.Pagination.Page
	trashcanDeviceList.PageSize = req.Pagination.PageSize
	return
}

func getDeviceInfo(db *gorm.DB, log log.Logger, deviceSN string) (trashcanDevice *protos.TrashcanDevice, err error) {
	deviceRecord := &model.Device{}
	whereArgs := map[string]interface{}{
		"device_type_id": deviceTypeID,
		"serial_number":  deviceSN,
	}
	err = db.Model(&model.Device{}).Where(whereArgs).First(&deviceRecord).Error
	if err != nil {
		log.Errorf("find trashcan device failed in getDeviceInfo(): %s", err.Error())
		return
	}

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, deviceRecord.DeviceTypeID).Error
	if err != nil {
		log.Errorf("find trashcan device type failed: in getDeviceInfo(): %s", err.Error())
		return
	}

	deviceModelRecord := &model.DeviceModel{}
	err = db.First(deviceModelRecord, deviceRecord.DeviceModelID).Error
	if err != nil {
		log.Errorf("find device type failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}

	groupRecord := &model.Group{}
	err = db.First(groupRecord, deviceRecord.GroupID).Error
	if err != nil {
		log.Errorf("find device type failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}

	trashcanDevice = &protos.TrashcanDevice{
		ID:            deviceRecord.ID,
		Name:          deviceRecord.Name,
		SerialNumber:  deviceRecord.SerialNumber,
		Position:      deviceRecord.Position,
		GroupID:       deviceRecord.GroupID,
		Group:         groupRecord.Name,
		DeviceTypeID:  deviceRecord.DeviceTypeID,
		DeviceType:    deviceTypeRecord.Name,
		DeviceModelID: deviceRecord.DeviceModelID,
		DeviceModel:   deviceModelRecord.Name,
		CustomerID:    deviceRecord.CustomerID,
		Latitude:      deviceRecord.Latitude,
		Longitude:     deviceRecord.Longitude,
	}
	return
}

// GetDevice ...
func GetDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (trashcanDevice *protos.TrashcanDevice, err error) {
	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			PageSize: 1,
		},
		CustomerID:   customerID,
		SerialNumber: deviceSN,
		DeviceTypeID: deviceTypeID,
	}
	deviceList, err := GetDevices(db, log, req)
	if err != nil {
		return
	}
	if deviceList.Count > 0 {
		trashcanDevice = deviceList.List[0]
	}
	return
}

// GetDeviceGroups 获取所有垃圾箱分组列表, 每个成员包含当前分组下的设备总量, 在线数量和离线数量信息.
func GetDeviceGroups(db *gorm.DB, log log.Logger, req *protos.GetTrashCanDeviceGroupsRequest) (resp *protos.GetTrashCanDeviceGroupsResponse, err error) {
	resp = &protos.GetTrashCanDeviceGroupsResponse{
		List: []*protos.TrashCanDeviceGroup{},
	}

	query := db.Model(&model.Group{})
	// 这里是条件查询
	whereArgs := map[string]interface{}{
		"customer_id":    req.CustomerID,
		"device_type_id": deviceTypeID,
	}

	query = query.Where(whereArgs)
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		log.Errorf("get customer: %d's trashcan device count failed in GetDeviceGroups(): %s", req.CustomerID, err.Error())
		err = nil
		return
	}

	query = pagination.BuildPaginationQuery(query, req.Pagination)
	groupRecords := []*model.Group{}
	err = query.Find(&groupRecords).Error
	if err != nil {
		log.Errorf("find group of trashcan failed in GetDeviceGroups(): %s", err.Error())
		err = nil
		return
	}
	// 接下来查询每个group里设备状态汇总情况.
	for _, group := range groupRecords {
		total, on, off, err := getDeviceGroup(db, log, group)
		if err != nil {
			log.Errorf("get device status in group: %d failed in GetDeviceGroups(): %s", group.ID, err.Error())
			err = nil
			continue
		}
		_group := &protos.TrashCanDeviceGroup{
			ID:          group.ID,
			Name:        group.Name,
			CustomerID:  group.CustomerID,
			DeviceTotal: total,
			DeviceOn:    on,
			DeviceOff:   off,
		}
		resp.List = append(resp.List, _group)
	}
	resp.Count = count
	resp.TotalCount = count
	resp.CurrentPage = req.Pagination.Page
	resp.PageSize = req.Pagination.PageSize
	return
}

// getDeviceGroup 获取单个分组的设备状态汇总, 包括当前分组下设备总量, 在线数量及离线数量.
func getDeviceGroup(db *gorm.DB, log log.Logger, group *model.Group) (total uint64, on uint64, off uint64, err error) {
	deviceRecords := []*model.Device{}
	query := db.Model(&model.Device{}).Where("group_id = ? and customer_id = ?", group.ID, group.CustomerID)
	_ = query.Count(&total).Error
	err = query.Find(&deviceRecords).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("find trashcan count failed in getDeviceGroup(): %s", err.Error())
		}
		return
	}
	for _, deviceRecord := range deviceRecords {
		status, err := GetDeviceStatus(db, log, deviceRecord.CustomerID, deviceRecord.SerialNumber)
		if err != nil {
			log.Errorf("get device %d status failed in getDeviceGroup(): %s", deviceRecord.ID, err.Error())
			err = nil
			off++ // 出错时定为离线
			continue
		}
		if status == 1 {
			on++
		} else if status == 2 {
			off++
		}
	}
	return
}

// GetDeviceStatus ...
// status: 1: 在线, 2: 离线, 3: 报警; 其他
func GetDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status int64, err error) {
	return _deviceManagementService.GetSafetyDeviceStatus(customerID, "trashcans", "trashcan", deviceSN)
}

// GetDeviceAlarmThresholds ...
func GetDeviceAlarmThresholds(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (tcAlarmThresholdsList *protos.TrashcanAlarmThresholdList, err error) {
	tcAlarmThresholdsList = &protos.TrashcanAlarmThresholdList{
		List: []*protos.TrashcanDeviceWithAlarmThreshold{},
	}
	req.DeviceTypeID = deviceTypeID
	alarmThresholdsList, err := _deviceManagementService.GetDeviceAlarmThresholds(req)
	if err != nil {
		return
	}
	for _, alarmThreshold := range alarmThresholdsList.List {
		trashcanAlarmThreshold := &protos.TrashcanDeviceWithAlarmThreshold{}
		err = trashcanModel2PB(alarmThreshold, trashcanAlarmThreshold)
		if err != nil {
			return
		}
		tcAlarmThresholdsList.List = append(tcAlarmThresholdsList.List, trashcanAlarmThreshold)
	}
	tcAlarmThresholdsList.Count = alarmThresholdsList.Count
	tcAlarmThresholdsList.TotalCount = alarmThresholdsList.Count
	tcAlarmThresholdsList.CurrentPage = req.Pagination.Page
	tcAlarmThresholdsList.PageSize = req.Pagination.PageSize
	return
}

// SetDeviceAlarmThreshold ...
func SetDeviceAlarmThreshold(db *gorm.DB, log log.Logger, req *protos.SetAlarmThresholdRequest) (err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.SetDeviceAlarmThreshold(req)
}

func trashcanModel2PB(record *protos.DeviceWithAlarmThreshold, pb *protos.TrashcanDeviceWithAlarmThreshold) (err error) {
	pb.ID = record.ID
	pb.Name = record.Name
	pb.Position = record.Position
	pb.SerialNumber = record.SerialNumber
	pb.Group = record.Group
	pb.GroupID = record.GroupID
	pb.CreatedAt = record.CreatedAt

	// 垃圾箱阈值是3元组
	if len(record.AlarmThresholds) == 3 {
		pb.StageInfo = record.AlarmThresholds[0]
		pb.StageWarn = record.AlarmThresholds[1]
		pb.StageAlert = record.AlarmThresholds[2]
	}

	return
}
