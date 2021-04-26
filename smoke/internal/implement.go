package internal

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/smoke"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultSmokeService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewSmokeService(
	l log.Logger,
	db *gorm.DB,
	cfg smoke.SmokeConfig,
) (service smoke.SmokeService, err error) {
	if db == nil {
		err = smoke.ErrDbNotBeNil
		return
	}
	dmService := MustGetDeviceManagementService(db, l)
	service = &DefaultSmokeService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ smoke.SmokeService = (*DefaultSmokeService)(nil)

var _deviceManagementService device_management.DeviceManagementService

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

func (ss *DefaultSmokeService) GetLastSmokeDeviceInfos(sns []string) (msg []smoke.Smoke, err error) {
	return GetLastSmokeDeviceInfos(ss.db, ss.l, sns)
}

func (ss *DefaultSmokeService) GetLastSmokeDeviceInfo(sns string) (msg *smoke.Smoke, err error) {
	return GetLastSmokeDeviceInfo(ss.db, ss.l, sns)
}

func (ss *DefaultSmokeService) GetSmokeDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error) {
	return GetSmokeDevices(ss.db, ss.l, req)
}

func (ss *DefaultSmokeService) GetSmokeDevice(customerID uint64, deviceSN string) (smokeDevice *protos.Device, err error) {
	return GetSmokeDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultSmokeService) GetSmokeDeviceStatus(customerID uint64, deviceSN string) (status int8, err error) {
	return GetSmokeDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}
