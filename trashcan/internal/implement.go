package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultTrashcanService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewTrashcanService(
	l log.Logger,
	db *gorm.DB,
	cfg trashcan.TrashcanConfig,
) (service trashcan.TrashcanService, err error) {
	if db == nil {
		err = trashcan.ErrDbNotBeNil
		return
	}

	dmService := MustGetDeviceManagementService(db, l)

	deviceTypeID, err = MustGetDeviceTypeID(db, l)
	if err != nil {
		panic(err)
	}

	service = &DefaultTrashcanService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ trashcan.TrashcanService = (*DefaultTrashcanService)(nil)

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
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "trashcan"}).Error
	if err != nil {
		log.Errorf("find manholeCover device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

func (ss *DefaultTrashcanService) GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.TrashcanDeviceList, err error) {
	return GetDevices(ss.db, ss.l, req)
}

func (ss *DefaultTrashcanService) GetDevice(customerID uint64, deviceSN string) (trashcanDevice *protos.TrashcanDevice, err error) {
	return GetDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultTrashcanService) GetDeviceGroups(req *protos.GetTrashCanDeviceGroupsRequest) (resp *protos.GetTrashCanDeviceGroupsResponse, err error) {
	return GetDeviceGroups(ss.db, ss.l, req)
}

func (ss *DefaultTrashcanService) GetDeviceStatus(customerID uint64, deviceSN string) (status int64, err error) {
	return GetDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultTrashcanService) GetDeviceAlarmThresholds(req *protos.GetDevicesRequestForCustomer) (msg *protos.TrashcanAlarmThresholdList, err error) {
	return GetDeviceAlarmThresholds(ss.db, ss.l, req)
}

func (ss *DefaultTrashcanService) SetDeviceAlarmThreshold(req *protos.SetAlarmThresholdRequest) (err error) {
	return SetDeviceAlarmThreshold(ss.db, ss.l, req)
}
