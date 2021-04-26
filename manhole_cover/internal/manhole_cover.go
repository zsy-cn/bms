package internal

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// GetDevices 获取设备列表, 需要带有报警状态
// 由于要按照报警状态排序, 所以不能按常规的分页查询方法再获取每一页设备各自的状态,
// 所以使用sql联合查询来完成.
func GetDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (manholeCoverDeviceList *protos.SafetyDeviceList, err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.GetSafetyDevices(req, "manhole_covers", offlineTimeout)
}

// GetDevice ...
func GetDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (manholeCoverDevice *protos.SafetyDevice, err error) {
	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			PageSize: 1,
		},
		CustomerID:   customerID,
		SerialNumber: deviceSN,
		DeviceTypeID: deviceTypeID,
	}
	deviceList, err := _deviceManagementService.GetSafetyDevices(req, "manhole_covers", offlineTimeout)
	if err != nil {
		return
	}
	if deviceList.Count > 0 {
		manholeCoverDevice = deviceList.List[0]
	}
	return
}

// GetDeviceInfos 获取指定过滤条件下的设备总量, 在线数量及报警数量
func GetDeviceInfos(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error) {
	req.DeviceTypeID = deviceTypeID
	return _deviceManagementService.GetDeviceInfos(req, "manhole_covers", offlineTimeout)
}

// GetDeviceStatus 获取指定设备状态
// @param: customerID 客户ID
// @param: deviceSN 设备序列号
// @return:
// 		status: -1: 设备不存在, 1: 正常, 2: 离线, 3: 报警, 4: 其他
func GetDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status int64, err error) {
	return _deviceManagementService.GetSafetyDeviceStatus(customerID, "manhole_covers", "manhole_cover", deviceSN)
}
