package water_level

import (
	"github.com/zsy-cn/bms/protos"
)

type WaterLevelService interface {
	GetDevice(customerID uint64, deviceSN string) (device *protos.SafetyDevice, err error)
	GetDeviceAlarmThresholds(req *protos.GetDevicesRequestForCustomer) (devices *protos.WaterLevelAlarmThresholdList, err error)
	SetDeviceAlarmThreshold(req *protos.SetAlarmThresholdRequest) (err error)

	GetDeviceStatus(customerID uint64, deviceSN string) (status int64, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.SafetyDeviceList, err error)
	GetDeviceInfos(req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error)
}
