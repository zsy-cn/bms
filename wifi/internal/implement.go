package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/wifi"
)

var offlineTimeout = "-48h" // 2天

type DefaultWifiService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewWifiService(
	l log.Logger,
	db *gorm.DB,
	cfg wifi.WifiConfig,
) (service wifi.WifiService, err error) {
	if db == nil {
		err = wifi.ErrDbNotBeNil
		return
	}
	dmService := MustGetDeviceManagementService(db, l)
	service = &DefaultWifiService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ wifi.WifiService = (*DefaultWifiService)(nil)

var _deviceManagementService device_management.DeviceManagementService

// 当前模块的设备类型ID, 在实例化服务时会从数据库中查询取得
var deviceTypeID uint64

func MustGetDeviceManagementService(db *gorm.DB, log log.Logger) device_management.DeviceManagementService {
	if _deviceManagementService != nil {
		return _deviceManagementService
	}

	deviceManagementService, err := dapp.NewDeviceManagementService(
		log,
		db,
		device_management.DeviceManagementConfig{},
	)
	if err != nil {
		panic(err)
	}
	deviceTypeID, err = MustGetDeviceTypeID(db, log)
	if err != nil {
		panic(err)
	}
	_deviceManagementService = deviceManagementService
	return _deviceManagementService
}

// MustGetDeviceTypeID 获取当前模块的设备类型ID(这个值会频繁使用到)
func MustGetDeviceTypeID(db *gorm.DB, log log.Logger) (id uint64, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "router"}).Error
	if err != nil {
		log.Errorf("find manholeCover device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

func (ss *DefaultWifiService) GetDeviceGroups(req *protos.GetDevicesRequestForCustomer) (groupList *protos.WifiDeviceGroupList, err error) {
	return GetDeviceGroups(ss.db, ss.l, req)
}

func (ss *DefaultWifiService) GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.WifiDeviceList, err error) {
	return GetDevices(ss.db, ss.l, req)
}
