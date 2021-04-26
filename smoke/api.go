package smoke

import (
	"github.com/zsy-cn/bms/protos"
)

type SmokeService interface {
	GetSmokeDevice(customerID uint64, deviceSN string) (device *protos.Device, err error)
	GetSmokeDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error)
	GetSmokeDeviceStatus(customerID uint64, deviceSN string) (status int8, err error)

	GetLastSmokeDeviceInfos(sns []string) (msg []Smoke, err error)
	GetLastSmokeDeviceInfo(sns string) (msg *Smoke, err error)
}
