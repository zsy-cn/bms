package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sos"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultSosService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewSosService(
	l log.Logger,
	db *gorm.DB,
	cfg sos.SosConfig,
) (service sos.SosService, err error) {
	if db == nil {
		err = sos.ErrDbNotBeNil
		return
	}

	dmService := MustGetDeviceManagementService(db, l)
	service = &DefaultSosService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ sos.SosService = (*DefaultSosService)(nil)

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

func (ss *DefaultSosService) GetLastSOSDeviceInfos(sns []string) (msg []*sos.Sos, err error) {
	return GetLastSOSDeviceInfos(ss.db, ss.l, sns)
}

func (ss *DefaultSosService) GetLastSOSDeviceInfo(sn string) (msg *sos.Sos, err error) {
	return GetLastSOSDeviceInfo(ss.db, ss.l, sn)
}

func (ss *DefaultSosService) GetSOSDevices(req *protos.GetDevicesRequestForCustomer) (sosDevice *protos.DeviceList, err error) {
	return GetSOSDevices(ss.db, ss.l, req)
}

func (ss *DefaultSosService) GetSOSDevice(customerID uint64, deviceSN string) (sosDevice *protos.Device, err error) {
	return GetSOSDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultSosService) GetSOSDeviceStatus(customerID uint64, deviceSN string) (status int8, msgID uint64, err error) {
	return GetSOSDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}
