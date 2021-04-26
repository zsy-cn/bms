package internal

import (
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/wifi"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
)

// GetDevices ...
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (deviceList *protos.WifiDeviceList, err error) {
	deviceList = &protos.WifiDeviceList{
		List: []*protos.WifiDevice{},
	}
	var count uint64 = 1
	deviceList.List = append(deviceList.List, &protos.WifiDevice{
		ID:           10000,
		Name:         "嵩华路公交站01",
		Position:     "嵩华路公交站",
		SerialNumber: "ZHWF00002309",

		GroupID:      10000,
		Group:        "嵩华路",
		CustomerID:   1,
		DeviceTypeID: 9,
		DeviceType:   "Wifi路由器",
		Longitude:    102.4929760294991,
		Latitude:     24.937236234452867,

		Connections: 50,
		UpSpeed:     "300KB",
		DownSpeed:   "1.2M",
		DataTraffic: "23G",

		CreatedAt: "2019-03-23 17:50:34",
	})

	deviceList.Count = count
	deviceList.TotalCount = count
	deviceList.CurrentPage = req.Pagination.Page
	deviceList.PageSize = req.Pagination.PageSize
	return
}

// GetDeviceGroups 获取指定过滤条件下的设备总量, 在线数量及报警数量
func GetDeviceGroups(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (groupList *protos.WifiDeviceGroupList, err error) {
	groupList = &protos.WifiDeviceGroupList{
		List: []*protos.WifiDeviceGroup{},
	}
	var count uint64 = 1
	groupList.List = append(groupList.List, &protos.WifiDeviceGroup{
		ID:   10000,
		Name: "嵩华路",

		CustomerID:   1,
		DeviceTypeID: 9,

		CreatedAt: "2019-03-23 17:50:34",

		DeviceTotal: 1,
		DeviceOn:    1,
		DeviceOff:   0,
	})

	groupList.Count = count
	groupList.TotalCount = count
	groupList.CurrentPage = req.Pagination.Page
	groupList.PageSize = req.Pagination.PageSize

	return
}

// GetDeviceInfos 获取指定过滤条件下的设备总量, 在线数量及报警数量
func GetDeviceInfos(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error) {
	return
}

func record2PB(record *wifi.Wifi, pb *protos.WifiDevice) (err error) {
	return
}
