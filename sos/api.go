package sos

import (
	"github.com/zsy-cn/bms/protos"
)

type SosService interface {
	GetSOSDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error)
	GetSOSDevice(customerID uint64, deviceSN string) (device *protos.Device, err error)
	GetLastSOSDeviceInfos(sns []string) (msg []*Sos, err error)
	GetLastSOSDeviceInfo(sn string) (msg *Sos, err error)
	GetSOSDeviceStatus(customerID uint64, deviceSN string) (status int8, msgID uint64, err error)
}
