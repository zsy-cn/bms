package manhole_cover

import (
	"github.com/zsy-cn/bms/protos"
)

type ManholeCoverService interface {
	GetDevice(customerID uint64, deviceSN string) (device *protos.SafetyDevice, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.SafetyDeviceList, err error)
	GetDeviceStatus(customerID uint64, deviceSN string) (status int64, err error)

	GetDeviceInfos(req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error)
}
