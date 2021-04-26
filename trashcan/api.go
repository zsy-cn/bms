package trashcan

import (
	"github.com/zsy-cn/bms/protos"
)

type TrashcanService interface {
	GetDevice(customerID uint64, deviceSN string) (device *protos.TrashcanDevice, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.TrashcanDeviceList, err error)
	GetDeviceGroups(req *protos.GetTrashCanDeviceGroupsRequest) (resp *protos.GetTrashCanDeviceGroupsResponse, err error)
	GetDeviceStatus(customerID uint64, deviceSN string) (status int64, err error)

	GetDeviceAlarmThresholds(req *protos.GetDevicesRequestForCustomer) (devices *protos.TrashcanAlarmThresholdList, err error)
	SetDeviceAlarmThreshold(req *protos.SetAlarmThresholdRequest) (err error)
}
