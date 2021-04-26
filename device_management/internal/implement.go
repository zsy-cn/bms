package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultDeviceManagementService struct {
	l  log.Logger
	db *gorm.DB
}

func NewDeviceManagementService(
	l log.Logger,
	db *gorm.DB,
	cfg device_management.DeviceManagementConfig,
) (service device_management.DeviceManagementService, err error) {
	if db == nil {
		err = device_management.ErrDbNotBeNil
		return
	}

	service = &DefaultDeviceManagementService{
		l:  l,
		db: db,
	}

	return service, nil
}

var _ device_management.DeviceManagementService = (*DefaultDeviceManagementService)(nil)

func (ss *DefaultDeviceManagementService) GetDevices(req *protos.GetDevicesRequestForCustomer) (devices *protos.DeviceList, err error) {
	return GetDevices(ss.db, ss.l, req)
}

func (ss *DefaultDeviceManagementService) GetDevice(deviceSN string, customerID uint64, deviceTypeID uint64) (device *protos.Device, err error) {
	return GetDevice(ss.db, ss.l, deviceSN, customerID, deviceTypeID)
}

func (ss *DefaultDeviceManagementService) GetDeviceGroupsByType(customerID uint64) (deviceGroups *protos.GetDeviceGroupsByTypeResponse, err error) {
	return GetDeviceGroupsByType(ss.db, ss.l, customerID)
}

func (ss *DefaultDeviceManagementService) GetSafetyDevices(req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (countInfo *protos.SafetyDeviceList, err error) {
	return GetSafetyDevices(ss.db, ss.l, req, tableName, offlineTimeout)
}

func (ss *DefaultDeviceManagementService) GetDeviceInfos(req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (countInfo *protos.DeviceCountInfo, err error) {
	return GetDeviceInfos(ss.db, ss.l, req, tableName, offlineTimeout)
}

func (ss *DefaultDeviceManagementService) GetSafetyDeviceStatus(customerID uint64, tableName string, deviceTypeKey string, deviceSN string) (status int64, err error) {
	return GetSafetyDeviceStatus(ss.db, ss.l, customerID, tableName, deviceTypeKey, deviceSN)
}

func (ss *DefaultDeviceManagementService) GetDeviceAlarmThresholds(req *protos.GetDevicesRequestForCustomer) (msg *protos.AlarmThresholdList, err error) {
	return GetDeviceAlarmThresholds(ss.db, ss.l, req)
}

func (ss *DefaultDeviceManagementService) SetDeviceAlarmThreshold(req *protos.SetAlarmThresholdRequest) (err error) {
	return SetDeviceAlarmThreshold(ss.db, ss.l, req)
}

func (ss *DefaultDeviceManagementService) ActiveDevice(deviceSN string, customerID uint64) (err error) {
	return ActiveDevice(ss.db, ss.l, deviceSN, customerID)
}

func (ss *DefaultDeviceManagementService) DeactiveDevice(deviceSN string, customerID uint64) (err error) {
	return DeactiveDevice(ss.db, ss.l, deviceSN, customerID)
}
