package device_management

import (
	"github.com/zsy-cn/bms/protos"
)

type DeviceManagementService interface {
	GetDevice(deviceSN string, customerID uint64, deviceTypeID uint64) (device *protos.Device, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (devices *protos.DeviceList, err error)
	GetDeviceGroupsByType(customerID uint64) (deviceGroups *protos.GetDeviceGroupsByTypeResponse, err error)

	ActiveDevice(deviceSN string, customerID uint64) (err error)
	DeactiveDevice(deviceSN string, customerID uint64) (err error)

	GetDeviceAlarmThresholds(req *protos.GetDevicesRequestForCustomer) (msg *protos.AlarmThresholdList, err error)
	SetDeviceAlarmThreshold(req *protos.SetAlarmThresholdRequest) (err error)

	GetDeviceInfos(req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (countInfo *protos.DeviceCountInfo, err error)
	GetSafetyDevices(req *protos.GetDevicesRequestForCustomer, tableName string, offlineTimeout string) (safetyDeviceList *protos.SafetyDeviceList, err error)
	GetSafetyDeviceStatus(customerID uint64, tableName string, deviceTypeKey string, deviceSN string) (status int64, err error)
}
