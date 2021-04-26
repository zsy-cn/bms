package wifi

import (
	"github.com/zsy-cn/bms/protos"
)

type WifiService interface {
	GetDeviceGroups(req *protos.GetDevicesRequestForCustomer) (groupList *protos.WifiDeviceGroupList, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.WifiDeviceList, err error)
}
