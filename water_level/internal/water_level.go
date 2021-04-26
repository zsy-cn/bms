package internal

import (
	"github.com/zsy-cn/bms/protos"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
)

// GetDevices ...
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (deviceList *protos.SafetyDeviceList, err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.GetSafetyDevices(req, "water_levels", offlineTimeout)
}

// GetDevice ...
func GetDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (waterLevelDevice *protos.SafetyDevice, err error) {
	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			PageSize: 1,
		},
		CustomerID:   customerID,
		SerialNumber: deviceSN,
		DeviceTypeID: deviceTypeID,
	}
	deviceList, err := _deviceManagementService.GetSafetyDevices(req, "water_levels", offlineTimeout)
	if err != nil {
		return
	}
	if deviceList.Count > 0 {
		waterLevelDevice = deviceList.List[0]
	}
	return
}

// GetDeviceInfos 获取指定过滤条件下的设备总量, 在线数量及报警数量
func GetDeviceInfos(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.GetDeviceInfos(req, "water_levels", offlineTimeout)
}

// GetDeviceStatus ...
// status: -1: 设备不存在, 1: 在线, 2: 离线, 3: 其他
func GetDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status int64, err error) {
	return _deviceManagementService.GetSafetyDeviceStatus(customerID, "water_levels", "water_level", deviceSN)
}

// GetDeviceAlarmThresholds ...
func GetDeviceAlarmThresholds(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (wlAlarmThresholdsList *protos.WaterLevelAlarmThresholdList, err error) {
	wlAlarmThresholdsList = &protos.WaterLevelAlarmThresholdList{
		List: []*protos.WaterLevelDeviceWithAlarmThreshold{},
	}
	req.DeviceTypeID = deviceTypeID
	alarmThresholdsList, err := _deviceManagementService.GetDeviceAlarmThresholds(req)
	if err != nil {
		return
	}
	for _, alarmThreshold := range alarmThresholdsList.List {
		waterlevelAlarmThreshold := &protos.WaterLevelDeviceWithAlarmThreshold{}
		err = waterlevelModel2PB(alarmThreshold, waterlevelAlarmThreshold)
		if err != nil {
			return
		}
		wlAlarmThresholdsList.List = append(wlAlarmThresholdsList.List, waterlevelAlarmThreshold)
	}
	wlAlarmThresholdsList.Count = alarmThresholdsList.Count
	wlAlarmThresholdsList.TotalCount = alarmThresholdsList.Count
	wlAlarmThresholdsList.CurrentPage = req.Pagination.Page
	wlAlarmThresholdsList.PageSize = req.Pagination.PageSize
	return
}

// SetDeviceAlarmThreshold ...
func SetDeviceAlarmThreshold(db *gorm.DB, log log.Logger, req *protos.SetAlarmThresholdRequest) (err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.SetDeviceAlarmThreshold(req)
}

func waterlevelModel2PB(record *protos.DeviceWithAlarmThreshold, pb *protos.WaterLevelDeviceWithAlarmThreshold) (err error) {
	pb.ID = record.ID
	pb.Name = record.Name
	pb.Position = record.Position
	pb.SerialNumber = record.SerialNumber
	pb.Group = record.Group
	pb.GroupID = record.GroupID
	pb.CreatedAt = record.CreatedAt

	// 水位阈值是4元组
	if len(record.AlarmThresholds) == 4 {
		pb.LowStageWarn = record.AlarmThresholds[0]
		pb.LowStageInfo = record.AlarmThresholds[1]
		pb.HighStageInfo = record.AlarmThresholds[2]
		pb.HighStageWarn = record.AlarmThresholds[3]
	}

	return
}
