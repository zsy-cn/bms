package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"

	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
	"github.com/zsy-cn/bms/util/sql"
	"github.com/zsy-cn/bms/util/timesection"

	"github.com/jinzhu/gorm"
)

// GetDevices ...
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error) {
	deviceList = &protos.DeviceList{
		List:  []*protos.Device{},
		Count: 0,
	}
	query := db.Model(&model.Device{})
	// ...这里应是条件查询语句
	whereArgs := map[string]interface{}{}
	if req.CustomerID == 0 {
		err = errors.New("please specify the customer id, you can not get all devices")
		log.Error(err)
		return
	}
	whereArgs["customer_id"] = req.CustomerID

	if req.DeviceTypeID > 0 {
		whereArgs["device_type_id"] = req.DeviceTypeID
	}
	if req.GroupID > 0 {
		whereArgs["group_id"] = req.GroupID
	}
	if req.SerialNumber != "" {
		whereArgs["serial_number"] = req.SerialNumber
	}
	query = query.Where(whereArgs)

	var count uint64
	deviceRecords := []*model.Device{}
	// 首先得到count总量
	err = query.Count(&count).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Errorf("find device count failed: %s", err.Error())
		} else {
			// 如果没有找到相关记录, 直接返回空和0, 不再继续执行
			err = nil
		}
		return
	}

	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)
	err = query.Find(&deviceRecords).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			log.Errorf("find device failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}
	for _, deviceRecord := range deviceRecords {
		devicePb := &protos.Device{}
		err = model2Pb(db, log, deviceRecord, devicePb)
		if err != nil {
			return
		}
		deviceList.List = append(deviceList.List, devicePb)
	}
	deviceList.Count = count
	deviceList.TotalCount = count
	deviceList.CurrentPage = req.Pagination.Page
	deviceList.PageSize = req.Pagination.PageSize
	return
}

// GetDevice ...
func GetDevice(db *gorm.DB, log log.Logger, deviceSN string, customerID uint64, deviceTypeID uint64) (devicePb *protos.Device, err error) {
	deviceRecord := &model.Device{}
	whereArgs := map[string]interface{}{
		"customer_id":   customerID,
		"serial_number": deviceSN,
	}
	if deviceTypeID != 0 {
		whereArgs["device_type_id"] = deviceTypeID
	}
	err = db.Model(&model.Device{}).Where(whereArgs).First(&deviceRecord).Error
	if err != nil {
		log.Errorf("find device failed in GetDevice(): %s", err.Error())
		return
	}
	devicePb = &protos.Device{}
	err = model2Pb(db, log, deviceRecord, devicePb)
	return
}

type DeviceTypeCount struct {
	DeviceTypeID uint64
	Count        uint64
}

// GetDeviceGroupsByType 用于设备列表页显示分组汇总信息, 只简单返回指定客户名下拥有的设备类型和各类型的设备数量.
// 注意: 无分页.
func GetDeviceGroupsByType(db *gorm.DB, log log.Logger, customerID uint64) (deviceGroupsByType *protos.GetDeviceGroupsByTypeResponse, err error) {
	deviceGroupsByType = &protos.GetDeviceGroupsByTypeResponse{
		List: []*protos.DeviceGroupByType{},
	}
	sqlStr := `
	select device_type_id, count(id) from devices where customer_id = %d group by device_type_id order by device_type_id
	`
	sqlStr = fmt.Sprintf(sqlStr, customerID)
	deviceTypeCountRecords := []*DeviceTypeCount{}
	err = db.Raw(sqlStr).Scan(&deviceTypeCountRecords).Error
	if err != nil {
		log.Errorf("query device types record for customer: %d failed in GetDeviceGroupsByType(): %s", customerID, err.Error())
		err = nil
		return
	}
	for _, deviceTypeCountRecord := range deviceTypeCountRecords {
		deviceTypeRecord := &model.DeviceType{}
		err = db.First(deviceTypeRecord, deviceTypeCountRecord.DeviceTypeID).Error
		if err != nil {
			log.Errorf("query device types record by id: %d failed in GetDeviceGroupsByType(): %s", deviceTypeCountRecord.DeviceTypeID, err.Error())
			err = nil
			continue
		}
		deviceGroupByType := &protos.DeviceGroupByType{
			ID:    deviceTypeRecord.ID,
			Name:  deviceTypeRecord.Name,
			Key:   deviceTypeRecord.Key,
			Count: deviceTypeCountRecord.Count,
		}
		deviceGroupsByType.List = append(deviceGroupsByType.List, deviceGroupByType)
		deviceGroupsByType.Count++
	}
	return
}

// GetSafetyDevices 获取安全设备信息列表, 各设备信息包含设备属性, 经纬度, 分组, 类型及名称.
// 尤其包含设备状态信息, 包括: 正常, 离线, 报警3种状态. 且按是否报警排序, 正在报警的设备排在前面.
// 函数由sql联合查询完成报警排序功能(与垃圾箱设备按使用率排序出于同样的考虑).
// @param: req 可设置通过设备类型, 分组, 型号等进行过滤, 一般只设置设备类型id
// @param: tableName 设备消息表名, 如果req中device type id为trashcan, 那么tableName应该为trashcans表.
// @param: offlineTime 离线阈值时间, 不同设备的离线判断时长不同.
func GetSafetyDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (safetyDeviceList *protos.SafetyDeviceList, err error) {
	safetyDeviceList = &protos.SafetyDeviceList{
		List: []*protos.SafetyDevice{},
	}

	safetyDeviceRecords := []*protos.SafetyDevice{}
	deadline := timesection.GetTimeDeadline(offlineTimeout)
	deadlineStr := deadline.Format("2006-01-02 15:04:05")

	// 查询所有设备最近的消息, 条件查询语句在后面拼接
	// 联合查询设备表, 消息表和报警表, 得到所有设备, 及最近一次有效的上行消息id和上行时间(如果有的话),
	// 还有报警表中最近一次的报警记录(如果有的话)
	// 根据消息id: m_id和报警id: n_id来判断该设备当前是正常, 离线还是报警.
	sqlStr := `
	select main_tbl.*, m_tbl.m_id, m_tbl.uplink_at, n_tbl.n_id, n_tbl.alert_at from devices as main_tbl 
	left join (
		select c.id as m_id, c.device_sn as m_device_sn, c.created_at as uplink_at from %s as c inner join (
			select device_sn, max(created_at) as created_at from %s where created_at > '%s' group by device_sn
		) as d on c.device_sn = d.device_sn and c.created_at = d.created_at
	) as m_tbl on main_tbl.serial_number = m_tbl.m_device_sn 
	left join (
		select a.id as n_id, a.device_sn as n_device_sn, a.updated_at as alert_at from notifications as a inner join (
			select device_sn, max(created_at) as created_at from notifications where solved = false and updated_at > '%s' group by device_sn
		) as b on a.device_sn = b.device_sn and a.created_at = b.created_at 
	) as n_tbl on main_tbl.serial_number = n_tbl.n_device_sn 
	`
	sqlStr = fmt.Sprintf(sqlStr, tableName, tableName, deadlineStr, deadlineStr)

	countSQLStr := "select count(id) from devices as main_tbl"

	whereStr := sql.MakeWhereStr(req)
	countSQLStr = countSQLStr + whereStr
	sqlStr = sqlStr + whereStr

	var count uint64
	err = db.Raw(countSQLStr).Count(&count).Error
	if err != nil {
		log.Errorf("query device message count failed in GetSafetyDevices(): %s", err.Error())
		return
	}
	sqlStr = pagination.BuildPaginationQueryInString(sqlStr, req.Pagination)
	log.Debugf("execute sql in GetSafetyDevices(): %s", sqlStr)
	err = db.Raw(sqlStr).Scan(&safetyDeviceRecords).Error
	if err != nil {
		log.Errorf("query device latest messages failed in GetSafetyDevices(): %s", err.Error())
		return
	}

	for _, safetyDeviceRecord := range safetyDeviceRecords {
		err = safetyModel2Pb(db, log, safetyDeviceRecord)
		if err != nil {
			log.Errorf("trans model to proto failed by safetyModel2Pb() in GetSafetyDevices(): %s", err.Error())
			return
		}
		safetyDeviceList.List = append(safetyDeviceList.List, safetyDeviceRecord)
	}
	safetyDeviceList.Count = count
	safetyDeviceList.TotalCount = count
	safetyDeviceList.CurrentPage = req.Pagination.Page
	safetyDeviceList.PageSize = req.Pagination.PageSize
	return
}

// GetDeviceInfos 获取某一类型的设备汇总信息, 包括设备总量, 在线数量及报警数量.
// 目前只有安全设备有调用此接口的需要.
// @param: req 可设置通过设备类型, 分组, 型号等进行过滤, 一般只设置设备类型id
// @param: tableName 设备消息表名, 如果req中device type id为trashcan, 那么tableName应该为trashcans表.
// @param: offlineTime 离线阈值时间, 不同设备的离线判断时长不同.
func GetDeviceInfos(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (countInfo *protos.DeviceCountInfo, err error) {
	countInfo = &protos.DeviceCountInfo{}
	var allCount, onlineCount, alarmCount uint64

	deadline := timesection.GetTimeDeadline(offlineTimeout)
	deadlineStr := deadline.Format("2006-01-02 15:04:05")

	allCountSQLStr := "select count(id) from devices as main_tbl"
	onlineCountSQLStr := `
		select count(main_tbl.id) from devices as main_tbl inner join (
			select device_sn, max(created_at) as created_at from %s where created_at > '%s' group by device_sn
		) as a on main_tbl.serial_number = a.device_sn
	`
	onlineCountSQLStr = fmt.Sprintf(onlineCountSQLStr, tableName, deadlineStr)

	alarmCountSQLStr := `
		select count(main_tbl.id) from devices as main_tbl inner join (
			select device_sn, max(created_at) as created_at from notifications where solved = false and updated_at > '%s' group by device_sn
		) as a on main_tbl.serial_number = a.device_sn
	`
	alarmCountSQLStr = fmt.Sprintf(alarmCountSQLStr, deadlineStr)

	whereStr := sql.MakeWhereStr(req)
	allCountSQLStr += whereStr
	onlineCountSQLStr += whereStr
	alarmCountSQLStr += whereStr

	err = db.Raw(allCountSQLStr).Count(&allCount).Error
	if err != nil {
		log.Errorf("query all device count failed in getDeviceInfos(): %s", err.Error())
		return
	}
	err = db.Raw(onlineCountSQLStr).Count(&onlineCount).Error
	if err != nil {
		log.Errorf("query device message count failed in getDeviceInfos(): %s", err.Error())
		return
	}
	err = db.Raw(alarmCountSQLStr).Count(&alarmCount).Error
	if err != nil {
		log.Errorf("query device alarm count failed in getDeviceInfos(): %s", err.Error())
		return
	}
	countInfo.Amount = allCount
	countInfo.Alive = onlineCount
	countInfo.Alarm = alarmCount
	return
}

// getDeviceTypeID 获取指定key的设备类型ID(这个值会频繁使用到)
func getDeviceTypeID(db *gorm.DB, log log.Logger, key string) (id uint64, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: key}).Error
	if err != nil {
		log.Errorf("find device type failed: in getDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

// GetSafetyDeviceStatus 获取指定安全设备状态
// @param: customerID 客户ID
// @param: deviceSN 设备序列号
// @return: status: -1: 设备不存在, 1: 正常, 2: 离线, 3: 报警, 4: 其他
func GetSafetyDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, tableName string, deviceTypeKey string, deviceSN string) (status int64, err error) {
	log.Debugf("GetSafetyDeviceStatus(). deviceSN: %s, customerID: %d", deviceSN, customerID)
	deviceTypeID, err := getDeviceTypeID(db, log, deviceTypeKey)
	if err != nil {
		log.Errorf("find device type failed: in GetSafetyDeviceStatus() by getDeviceTypeID(): %s", err.Error())
		return
	}
	// 确认设备是否存在
	deviceRecord := &model.Device{}
	whereArgs := map[string]interface{}{
		"device_type_id": deviceTypeID,
		"customer_id":    customerID,
		"serial_number":  deviceSN,
	}
	err = db.Model(&model.Device{}).Where(whereArgs).First(&deviceRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
			status = -1
		} else {
			log.Errorf("find device: %s failed in GetManholeCoverDevice(): %s", deviceSN, err.Error())
		}
		return
	}

	deadline := timesection.GetTimeDeadline(offlineTimeout)
	var whereStr string

	// 查看最近消息
	var msgCount uint64
	whereStr = "device_sn = ? and created_at > ?"
	err = db.Table(tableName).Where(whereStr, deviceSN, deadline).Count(&msgCount).Error
	if msgCount == 0 {
		status = 2
		err = nil
		return
	}

	// 如果设备在线, 查看是否有相关警报
	notificationRecord := &model.Notification{}
	whereStr += "and solved = false"
	err = db.Where(whereStr, deviceSN, deadline).Order("id desc").First(notificationRecord).Error
	if err == nil {
		status = 3
	} else {
		if err.Error() == "record not found" {
			status = 1
			err = nil
		} else {
			log.Errorf("find notification record for manhole cover device: %s failed in GetSafetyDeviceStatus(): %s", deviceSN, err.Error())
		}
		return
	}
	return
}

// GetDeviceAlarmThresholds 获取指定类型设备的报警阈值列表
// @param: req 必须指定CustomerID和DeviceTypeID, 其他如分组ID, 型号ID可选
func GetDeviceAlarmThresholds(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (msg *protos.AlarmThresholdList, err error) {
	msg = &protos.AlarmThresholdList{
		List:  []*protos.DeviceWithAlarmThreshold{},
		Count: 0,
	}
	deviceList, err := GetDevices(db, log, req)
	if err != nil {
		log.Errorf("get devices failed in GetDeviceAlarmThresholds(): %s", err.Error())
		return
	}
	for _, device := range deviceList.List {
		ruleRecord := &model.Rule{}
		err = db.Where(&model.Rule{DeviceSN: device.SerialNumber}).First(ruleRecord).Error
		if err != nil {
			if err.Error() == "record not found" {
				ruleRecord.Section = &model.Float64Array{}
				err = nil
			} else {
				log.Errorf("get rule record for device: %s failed in GetDeviceAlarmThresholds(): %s", device.SerialNumber, err.Error())
				return
			}
		}
		groupRecord := &model.Group{}
		err = db.First(groupRecord, device.GroupID).Error
		if err != nil {
			log.Errorf("find device group failed: in GetDeviceAlarmThresholds(): %s", err.Error())
			return
		}
		loc, _ := time.LoadLocation("Asia/Shanghai")
		createdAt := ruleRecord.CreatedAt.In(loc).Format("2006-01-02 15:04:05")
		_msg := &protos.DeviceWithAlarmThreshold{
			ID:              device.ID,
			Name:            device.Name,
			Position:        device.Position,
			SerialNumber:    device.SerialNumber,
			GroupID:         device.GroupID,
			Group:           groupRecord.Name,
			DeviceModelID:   device.DeviceModelID,
			CreatedAt:       createdAt,
			AlarmThresholds: []float64(*ruleRecord.Section),
		}
		msg.List = append(msg.List, _msg)
	}
	msg.Count = deviceList.Count
	msg.TotalCount = deviceList.Count
	msg.CurrentPage = req.Pagination.Page
	msg.PageSize = req.Pagination.PageSize
	return
}

// SetDeviceAlarmThreshold ...
func SetDeviceAlarmThreshold(db *gorm.DB, log log.Logger, req *protos.SetAlarmThresholdRequest) (err error) {
	_, err = GetDevice(db, log, req.DeviceSN, req.CustomerID, req.DeviceTypeID)
	if err != nil {
		log.Errorf("get device: %s failed in SetDeviceAlarmThreshold(): %s", req.DeviceSN, err.Error())
		return
	}
	section := model.Float64Array(req.AlarmThresholds)

	whereArgs := &model.Rule{
		DeviceSN:     req.DeviceSN,
		CustomerID:   req.CustomerID,
		DeviceTypeID: req.DeviceTypeID,
	}
	ruleRecord := &model.Rule{}
	err = db.Where(whereArgs).First(ruleRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("get rule for device: %s failed in SetDeviceAlarmThreshold(): %s", req.DeviceSN, err.Error())
			return
		}
	}
	if ruleRecord.ID == 0 {
		ruleModel := &model.Rule{
			CustomerID:   req.CustomerID,
			DeviceSN:     req.DeviceSN,
			DeviceTypeID: req.DeviceTypeID,
			Section:      &section,
		}
		err = db.Create(ruleModel).Error
		if err != nil {
			log.Errorf("create rule for device: %s failed in SetDeviceAlarmThreshold(): %s", req.DeviceSN, err.Error())
			return
		}
	} else {
		err = db.Model(ruleRecord).UpdateColumn("section", &section).Error
		if err != nil {
			log.Errorf("update rule for device: %s failed in SetDeviceAlarmThreshold(): %s", req.DeviceSN, err.Error())
			return
		}
	}
	return
}

// ActiveDevice ...
func ActiveDevice(db *gorm.DB, log log.Logger, deviceSN string, customerID uint64) (err error) {
	deviceRecord := &model.Device{}
	whereArgs := map[string]interface{}{
		"customer_id":   customerID,
		"serial_number": deviceSN,
	}

	err = db.Where(whereArgs).First(&deviceRecord).Error
	if err != nil {
		log.Errorf("find device failed in ActiveDevice(): %s", err.Error())
		return
	}

	err = db.Model(deviceRecord).UpdateColumn("actived", true).Error
	if err != nil {
		log.Errorf("active device failed in ActiveDevice(): %s", err.Error())
		return
	}
	return
}

// DeactiveDevice ...代码同ActiveDevice()
func DeactiveDevice(db *gorm.DB, log log.Logger, deviceSN string, customerID uint64) (err error) {
	deviceRecord := &model.Device{}
	whereArgs := map[string]interface{}{
		"customer_id":   customerID,
		"serial_number": deviceSN,
	}

	err = db.Where(whereArgs).First(&deviceRecord).Error
	if err != nil {
		log.Errorf("find device failed in DeactiveDevice(): %s", err.Error())
		return
	}

	err = db.Model(deviceRecord).UpdateColumn("actived", false).Error
	if err != nil {
		log.Errorf("deactive device failed in DeactiveDevice(): %s", err.Error())
		return
	}
	return
}

func model2Pb(db *gorm.DB, log log.Logger, record *model.Device, pb *protos.Device) (err error) {
	pb.ID = record.ID
	pb.Name = record.Name
	pb.SerialNumber = record.SerialNumber
	pb.DeviceModelID = record.DeviceModelID
	pb.DeviceTypeID = record.DeviceTypeID
	pb.GroupID = record.GroupID
	pb.Description = record.Description
	pb.CustomerID = record.CustomerID
	pb.Position = record.Position
	pb.Latitude = record.Latitude
	pb.Longitude = record.Longitude
	pb.StatusCode = 1

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, record.DeviceTypeID).Error
	if err != nil {
		log.Errorf("find device type failed for device: %s in model2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	pb.DeviceType = deviceTypeRecord.Name

	deviceModelRecord := &model.DeviceModel{}
	err = db.First(deviceModelRecord, record.DeviceModelID).Error
	if err != nil {
		log.Errorf("find device model failed for device: %s in model2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	pb.DeviceModel = deviceModelRecord.Name

	groupRecord := &model.Group{}
	err = db.First(groupRecord, record.GroupID).Error
	if err != nil {
		log.Errorf("find device group failed for device: %s in model2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	pb.Group = groupRecord.Name

	loc, _ := time.LoadLocation("Asia/Shanghai")
	pb.CreatedAt = record.CreatedAt.In(loc).Format("2006-01-02")
	if record.Actived {
		pb.Actived = "true"
	} else {
		pb.Actived = "false"
	}

	return
}

func safetyModel2Pb(db *gorm.DB, log log.Logger, record *protos.SafetyDevice) (err error) {
	if record.NID != 0 {
		record.StatusCode = 3
		record.Status = "报警"
	} else if record.MID == 0 {
		record.StatusCode = 2
		record.Status = "离线"
	} else {
		record.StatusCode = 1
		record.Status = "正常"
	}

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, record.DeviceTypeID).Error
	if err != nil {
		log.Errorf("find device type failed for device: %s in safetyModel2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	record.DeviceType = deviceTypeRecord.Name

	deviceModelRecord := &model.DeviceModel{}
	err = db.First(deviceModelRecord, record.DeviceModelID).Error
	if err != nil {
		log.Errorf("find device model failed for device: %s in safetyModel2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	record.DeviceModel = deviceModelRecord.Name

	groupRecord := &model.Group{}
	err = db.First(groupRecord, record.GroupID).Error
	if err != nil {
		log.Errorf("find device group failed for device: %s in safetyModel2Pb(): %s", record.SerialNumber, err.Error())
		return
	}
	record.Group = groupRecord.Name

	if record.CreatedAt != "" {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		// 这里的record.CreatedAt是程序中数据库连接后使用上海时区格式化过的时间, 其实可以不用处理直接使用的.
		createdAt, err := time.Parse("2006-01-02T15:04:05-07:00", record.CreatedAt)
		if err != nil {
			log.Errorf("parse time for record: %s in safetyModel2Pb(): %s", record.SerialNumber, err.Error())
			return nil
		}
		record.CreatedAt = createdAt.In(loc).Format("2006-01-02")
	}

	return
}
