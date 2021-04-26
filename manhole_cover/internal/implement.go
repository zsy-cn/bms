package internal

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultManholeCoverService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewManholeCoverService(
	l log.Logger,
	db *gorm.DB,
	cfg manhole_cover.ManholeCoverConfig,
) (service manhole_cover.ManholeCoverService, err error) {
	if db == nil {
		panic("db can not be nil")
		return
	}
	dmService := MustGetDeviceManagementService(db, l)
	service = &DefaultManholeCoverService{
		l:         l,
		db:        db,
		dmService: dmService,
	}
	deviceTypeID, err = MustGetDeviceTypeID(db, l)
	if err != nil {
		panic(err)
	}
	return service, nil
}

var _ manhole_cover.ManholeCoverService = (*DefaultManholeCoverService)(nil)
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

	_deviceManagementService = deviceManagementService
	return _deviceManagementService
}

// MustGetDeviceTypeID 获取当前模块的设备类型ID(这个值会频繁使用到)
func MustGetDeviceTypeID(db *gorm.DB, log log.Logger) (id uint64, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "manhole_cover"}).Error
	if err != nil {
		log.Errorf("find manholeCover device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

func (ss *DefaultManholeCoverService) GetDeviceInfos(req *protos.GetDevicesRequestForCustomer) (countInfo *protos.DeviceCountInfo, err error) {
	return GetDeviceInfos(ss.db, ss.l, req)
}

func (ss *DefaultManholeCoverService) GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.SafetyDeviceList, err error) {
	return GetDevices(ss.db, ss.l, req)
}

func (ss *DefaultManholeCoverService) GetDevice(customerID uint64, deviceSN string) (manholeCoverDevice *protos.SafetyDevice, err error) {
	return GetDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultManholeCoverService) GetDeviceStatus(customerID uint64, deviceSN string) (status int64, err error) {
	return GetDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}
