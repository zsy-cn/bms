package geomagnetic

import (
	"github.com/zsy-cn/bms/protos"
)

type GeomagneticService interface {
	GetGeomagneticDevice(customerID uint64, deviceSN string) (device *protos.GeomagneticDevice, err error)
	GetGeomagneticDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.GeomagneticDeviceList, err error)

	GetParkingInfos(req *protos.GetParkingPlacesRequest) (list *protos.ParkingPlaceList, err error)
	GetParkingInfo(customerID, groupID uint64) (parkingPlace *protos.ParkingPlace, err error)
	GetParkingHistory(customerID uint64, groupID uint64, date string) (parkingHistory []map[string]interface{}, err error)
}
