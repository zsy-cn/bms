package internal

import (
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sound_box"

	"github.com/zsy-cn/bms/model"

	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"

	"github.com/jinzhu/gorm"
)

// GetDevices ...
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (soundboxDeviceList *protos.SoundBoxDeviceList, err error) {
	req.DeviceTypeID = deviceTypeID
	deviceList, err := _deviceManagementService.GetDevices(req)
	if err != nil {
		log.Errorf("get device list by device management service failed in GetDevices(): %s", err.Error())
		return
	}
	soundboxDeviceList = &protos.SoundBoxDeviceList{
		List: []*protos.SoundBoxDevice{},
	}
	for _, device := range deviceList.List {
		soundboxDevice := &protos.SoundBoxDevice{}
		err = deviceToSoundBoxDevice(device, soundboxDevice)
		if err != nil {
			log.Errorf("trans device to soundbox failed in GetDevices(): %s", err.Error())
			return
		}
		soundboxDeviceList.List = append(soundboxDeviceList.List, soundboxDevice)
	}
	soundboxDeviceList.Count = deviceList.Count
	soundboxDeviceList.TotalCount = deviceList.Count
	soundboxDeviceList.CurrentPage = req.Pagination.Page
	soundboxDeviceList.PageSize = req.Pagination.PageSize
	return
}

// GetDevice ...
func GetDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (soundboxDevice *protos.SoundBoxDevice, err error) {
	req := &protos.GetDevicesRequestForCustomer{
		CustomerID:   customerID,
		SerialNumber: deviceSN,
		DeviceTypeID: deviceTypeID,
		Pagination: &protos.Pagination{
			PageSize: 1,
		},
	}
	deviceList, err := GetDevices(db, log, req)
	if err != nil {
		log.Errorf("get geomagnetic device failed in GetDevice(): %s", err.Error())
		return
	}
	if len(deviceList.List) > 0 {
		soundboxDevice = deviceList.List[0]
	}
	return
}

// GetSoundBoxDeviceStatus ...
// status: -1: 设备不存在, 1: 在线, 2: 离线, 3: 播放中
func GetSoundBoxDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status int8, err error) {
	_, err = GetDevice(db, log, customerID, deviceSN)
	if err != nil {
		status = -1
		return
	}
	soundboxRecord := &sound_box.SoundBox{}
	whereArgs := map[string]interface{}{
		"device_sn": deviceSN,
	}
	err = db.Where(whereArgs).First(soundboxRecord).Error
	if err != nil {
		log.Errorf("can not find record for soundbox %s, you should repair it", deviceSN)
		status = -1
		return
	}

	return getSoundBoxStatus(log, soundboxRecord.SeedCode)
}

// GetSoundBoxDeviceGroups ...
func GetSoundBoxDeviceGroups(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (resp *protos.SoundBoxDeviceGroupList, err error) {
	resp = &protos.SoundBoxDeviceGroupList{
		List: []*protos.SoundBoxDeviceGroup{},
	}
	query := db.Model(&model.Group{})

	var count uint64
	whereArgs := map[string]interface{}{
		"customer_id":    req.CustomerID,
		"device_type_id": deviceTypeID,
	}

	query = query.Where(whereArgs)
	err = query.Count(&count).Error
	if err != nil {
		log.Errorf("get customer: %d's soundbox device count failed in GetSoundBoxDeviceGroups(): %s", req.CustomerID, err.Error())
		err = nil
		return
	}

	query = pagination.BuildPaginationQuery(query, req.Pagination)
	groupRecords := []*model.Group{}
	err = query.Find(&groupRecords).Error
	if err != nil {
		log.Errorf("find group of soundbox failed in GetSoundBoxDeviceGroups(): %s", err.Error())
		err = nil
		return
	}

	for _, group := range groupRecords {
		var _count uint64
		err = db.Model(&model.Device{}).Where("group_id = ? and customer_id = ?", group.ID, group.CustomerID).Count(&_count).Error
		if err != nil {
			if err.Error() == "record not found" {
				err = nil
			} else {
				log.Errorf("find soundbox count failed in GetSoundBoxDeviceGroups(): %s", err.Error())
				return
			}
		}
		_group := &protos.SoundBoxDeviceGroup{
			ID:          group.ID,
			Name:        group.Name,
			CustomerID:  group.CustomerID,
			DeviceTotal: _count,
			// 分组的播放状态需要找一个地方存储, 以后再说吧.
			Status: "就绪",
		}
		resp.List = append(resp.List, _group)
	}
	resp.Count = count
	resp.TotalCount = count
	resp.CurrentPage = req.Pagination.Page
	resp.PageSize = req.Pagination.PageSize
	return
}

func deviceToSoundBoxDevice(device *protos.Device, soundboxDevice *protos.SoundBoxDevice) (err error) {
	soundboxDevice.ID = device.ID
	soundboxDevice.Name = device.Name
	soundboxDevice.SerialNumber = device.SerialNumber
	soundboxDevice.DeviceModel = device.DeviceModel
	soundboxDevice.DeviceModelID = device.DeviceModelID
	soundboxDevice.DeviceType = device.DeviceType
	soundboxDevice.DeviceTypeID = device.DeviceTypeID
	soundboxDevice.Group = device.Group
	soundboxDevice.GroupID = device.GroupID
	soundboxDevice.CustomerID = device.CustomerID
	soundboxDevice.Position = device.Position
	soundboxDevice.Latitude = device.Latitude
	soundboxDevice.Longitude = device.Longitude
	soundboxDevice.Actived = device.Actived
	soundboxDevice.CreatedAt = device.CreatedAt

	// 这里可以获取设备的在线状态和音量信息
	return
}
