package internal

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
	"github.com/zsy-cn/bms/util/timesection"
)

// getLastGeomagneticDeviceInfos 查询指定条件下的设备最近一次的环境记录(不考虑离线状态, 以最近一次记录为准)
// @param: datetime 表示只查询指定时间点之前的记录, 如果为空, 则查询当前时间之前的环境设备记录
func getLastGeomagneticDeviceInfos(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, datetime string) (geoRecords []*geomagnetic.Geomagnetic, err error) {
	geoRecords = []*geomagnetic.Geomagnetic{}

	sqlStr1 := "select device_sn, max(created_at) as created_at from geomagnetics where 1 = 1 %s group by device_sn"
	sqlStr2 := "select main_tbl.* from geomagnetics as main_tbl inner join (%s) as b on main_tbl.device_sn = b.device_sn and main_tbl.created_at = b.created_at"
	whereStr := ""
	if customerID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and customer_id = %d ", customerID)
	}
	if groupID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and group_id = %d ", groupID)
	}
	if datetime != "" {
		whereStr = fmt.Sprintf(whereStr+" and created_at < timestamp '%s' ", datetime)
	}
	_sqlStr1 := fmt.Sprintf(sqlStr1, whereStr)
	sqlStr := fmt.Sprintf(sqlStr2, _sqlStr1)
	log.Debugf("execute sql in GetEnvironMonitorSectionAverageData(): %s", sqlStr)

	err = db.Raw(sqlStr).Scan(&geoRecords).Error
	if err != nil {
		log.Errorf("query geomagnetic message failed in getLastGeomagneticDeviceInfos(): %s", err.Error())
		return
	}
	return
}

// getLastGeomagneticDeviceInfo 获取指定设备的最近一次的有效消息记录.
func getLastGeomagneticDeviceInfo(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (msg *geomagnetic.Geomagnetic, err error) {
	msg = &geomagnetic.Geomagnetic{}
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline := now.Add(d2)
	whereArgs := "device_sn = ? and created_at > ?"
	err = db.Where(whereArgs, deviceSN, deadline).Order("id desc").First(msg).Error
	return
}

// GetGeomagneticDevices ...
func GetGeomagneticDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (geomagneticDeviceList *protos.GeomagneticDeviceList, err error) {
	geomagneticDeviceList = &protos.GeomagneticDeviceList{
		List: []*protos.GeomagneticDevice{},
	}

	req.DeviceTypeID = deviceTypeID
	deviceList, err := _deviceManagementService.GetDevices(req)
	if err != nil {
		log.Errorf("get geomagnetic device list failed in GetGeomagneticDevices(): %s", err.Error())
		return
	}
	for _, deviceRecord := range deviceList.List {
		geomagneticDevice := &protos.GeomagneticDevice{
			ID:            deviceRecord.ID,
			Name:          deviceRecord.Name,
			SerialNumber:  deviceRecord.SerialNumber,
			Position:      deviceRecord.Position,
			GroupID:       deviceRecord.GroupID,
			Group:         deviceRecord.Group,
			DeviceTypeID:  deviceRecord.DeviceTypeID,
			DeviceType:    deviceRecord.DeviceType,
			DeviceModelID: deviceRecord.DeviceModelID,
			DeviceModel:   deviceRecord.DeviceModel,
			CustomerID:    deviceRecord.CustomerID,
			Latitude:      deviceRecord.Latitude,
			Longitude:     deviceRecord.Longitude,
		}
		msg, err := getLastGeomagneticDeviceInfo(db, log, deviceRecord.CustomerID, deviceRecord.SerialNumber)
		if err != nil {
			geomagneticDevice.StatusCode = 2
			geomagneticDevice.Status = "离线"
		} else {
			geomagneticDevice.StatusCode = 1
			geomagneticDevice.Status = "正常"
			if msg.Value == "0" {
				geomagneticDevice.Used = "false"
			} else {
				geomagneticDevice.Used = "true"
			}
		}
		geomagneticDeviceList.List = append(geomagneticDeviceList.List, geomagneticDevice)
	}
	geomagneticDeviceList.Count = deviceList.Count
	geomagneticDeviceList.TotalCount = deviceList.Count
	geomagneticDeviceList.CurrentPage = req.Pagination.Page
	geomagneticDeviceList.PageSize = req.Pagination.PageSize
	return
}

// GetGeomagneticDevice ...
func GetGeomagneticDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (geomagneticDevice *protos.GeomagneticDevice, err error) {
	req := &protos.GetDevicesRequestForCustomer{
		CustomerID:   customerID,
		SerialNumber: deviceSN,
		DeviceTypeID: deviceTypeID,
		Pagination: &protos.Pagination{
			PageSize: 1,
		},
	}
	deviceList, err := GetGeomagneticDevices(db, log, req)
	if err != nil {
		log.Errorf("get geomagnetic device failed in GetGeomagneticDevice(): %s", err.Error())
		return
	}
	if len(deviceList.List) > 0 {
		geomagneticDevice = deviceList.List[0]
	}
	return
}

// GetParkingInfos ...
// 获取目标客户下所有停车场分组, 并查询分组下的设备使用情况(即占位情况).
func GetParkingInfos(db *gorm.DB, log log.Logger, req *protos.GetParkingPlacesRequest) (parkingPlaces *protos.ParkingPlaceList, err error) {
	parkingPlaces = &protos.ParkingPlaceList{
		List:  []*protos.ParkingPlace{},
		Count: 0,
	}
	/////////////////////////////////////// 首先查询group分组数据
	query := db.Model(&model.Group{})
	// 条件查询语句
	whereArgs := map[string]interface{}{
		"customer_id":    req.CustomerID,
		"device_type_id": deviceTypeID,
	}
	query = query.Where(whereArgs)

	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("get customer: %d's group count failed in GetParkingInfos(): %s", req.CustomerID, err.Error())
		}
		return
	}
	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)
	groupRecords := []*model.Group{}
	err = query.Find(&groupRecords).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("get customer: %d's groups failed in GetParkingInfos(): %s", req.CustomerID, err.Error())
		}
		return
	}
	/////////////////////////////////////// 查询每个停车场分组的经纬度及设备总量
	for _, group := range groupRecords {
		parkingPlace, err := GetParkingInfo(db, log, group.CustomerID, group.ID)
		if err != nil {
			log.Errorf("get parking info for group %d failed in GetParkingInfos(): %s", group.ID, err.Error())
			return parkingPlaces, err
		}
		parkingPlace.ID = group.ID
		parkingPlace.Name = group.Name
		parkingPlaces.List = append(parkingPlaces.List, parkingPlace)
	}
	parkingPlaces.Count = count
	parkingPlaces.TotalCount = count
	parkingPlaces.CurrentPage = req.Pagination.Page
	parkingPlaces.PageSize = req.Pagination.PageSize
	return
}

// GetParkingInfo 获取单个停车场分组的详细状态, 包括名称, 设备总量, 占用数量, 停车场经纬度等.
func GetParkingInfo(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64) (parkingPlace *protos.ParkingPlace, err error) {
	parkingPlace = &protos.ParkingPlace{
		ID: groupID,
	}
	groupRecord := &model.Group{}
	whereStr := "id = ? and customer_id = ?"
	err = db.Where(whereStr, groupID, customerID).First(groupRecord).Error
	if err != nil {
		return
	}
	parkingPlace.Name = groupRecord.Name
	// 获取指定停车场分组下的设备数量及占用数量
	var count uint64
	whereStr1 := "group_id = ? and customer_id = ?"
	err = db.Model(&model.Device{}).Where(whereStr1, groupID, customerID).Count(&count).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("find device serial numbers failed in GetParkingInfo(): %s", err.Error())
		}
		return
	}
	geoRecords, err := getLastGeomagneticDeviceInfos(db, log, customerID, groupID, "")
	if err != nil {
		log.Errorf("get geomagnetic message failed in GetParkingInfo(): %s", err.Error())
		return
	}

	parkingPlace.Amount = count
	for _, msg := range geoRecords {
		if msg.Value == "1" {
			parkingPlace.Used++
		}
		if msg.Value == "0" {
			parkingPlace.Unused++
		}
	}

	// 如果该停车场下存在设备的话, 从中随机寻找一个拥有经纬度数据的设备信息, 将其作为停车场位置
	device := &model.Device{}
	whereStr2 := "customer_id = ? and group_id = ? and longitude != 0 and latitude != 0"
	err = db.Model(&model.Device{}).Where(whereStr2, customerID, groupID).First(device).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("get one device of group: %d in GetParkingInfo() failed: %s", groupID, err.Error())
		}
		return
	}
	parkingPlace.Latitude = device.Latitude
	parkingPlace.Longitude = device.Longitude

	return
}

type usedInfo struct {
	Count uint64
}

// GetParkingHistory 获取停车场指定日期的占位历史(单日历史)
// @param date: 目标日期, 格式为 2019-01-01. 默认为空, 返回当天的占位历史.
// @return parkingHistory: 长度确定的切片.
func GetParkingHistory(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, date string) (parkingHistory []map[string]interface{}, err error) {
	dateFrom := date
	dateTo := date
	return getParkingHistory(db, log, customerID, groupID, dateFrom, dateTo)
}

// getParkingHistory 获取停车场指定日期的占位历史(可返回多天的占位历史数据)
// @param dateFrom: 起始日期, 格式为 2019-01-01, 默认为空.
// @param dateTo: 截止日期, 格式相同, 默认为空.
// @return parkingHistory: 长度确定的切片.
func getParkingHistory(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, dateFrom string, dateTo string) (parkingHistory []map[string]interface{}, err error) {
	log.Debug("get parking history in getParkingHistory()")

	parkingHistory = []map[string]interface{}{}

	var count uint64
	whereArgs := map[string]interface{}{
		"customer_id": customerID,
	}
	if groupID != 0 {
		whereArgs["group_id"] = groupID
	}
	err = db.Model(&model.Device{}).Where(whereArgs).Count(&count).Error
	if err != nil {
		log.Errorf("get device count for customer: %d failed in getParkingHistory(): %s", customerID, err.Error())
		return
	}

	startDate, endDate, now, err := timesection.GetTimeSection(dateFrom, dateTo)
	if err != nil {
		log.Errorf("get time section failed in getParkingHistory(): %s", err.Error())
		return
	}
	log.Debugf("startDate: %+v, endDate: %+v", startDate, endDate)

	sqlStr := `
		select count(value) from geomagnetics as a inner join (
			select device_sn, max(created_at) as created_at from geomagnetics where created_at < ? group by device_sn
		) as b on a.device_sn = b.device_sn and a.created_at = b.created_at and a.value = '1'
	`
	d1h, _ := time.ParseDuration("1h")
	startDate = startDate.Add(d1h)
	for {
		if startDate.After(endDate) || startDate.Equal(endDate) {
			break
		}
		key := fmt.Sprintf("%d时", startDate.Hour())
		// 如果构造的时间点已经超过了当前此刻, 则返回0, 不再查询.
		if startDate.After(now) {
			parkingHistory = append(parkingHistory, map[string]interface{}{
				"hour":    key,
				"percent": float64(0),
			})
		} else {
			_usedInfo := &usedInfo{}
			err = db.Raw(sqlStr, startDate).Scan(_usedInfo).Error
			if err != nil {
				log.Errorf("get parking history failed in getParkingHistory(): %s", err.Error())
				return
			}
			parkingHistory = append(parkingHistory, map[string]interface{}{
				"hour":    key,
				"percent": float64(_usedInfo.Count) / float64(count),
			})
		}
		startDate = startDate.Add(d1h * time.Duration(2))
	}

	return
}
