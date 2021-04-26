package internal

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-24h" // 1天

type DefaultGeomagneticService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewGeomagneticService(
	l log.Logger,
	db *gorm.DB,
	cfg geomagnetic.GeomagneticConfig,
) (service geomagnetic.GeomagneticService, err error) {
	if db == nil {
		err = geomagnetic.ErrDbNotBeNil
		return
	}
	dmService := MustGetDeviceManagementService(db, l)
	deviceTypeID, err = MustGetDeviceTypeID(db, l)
	if err != nil {
		panic(err)
	}
	service = &DefaultGeomagneticService{
		l:         l,
		db:        db,
		dmService: dmService,
	}
	return service, nil
}

var _ geomagnetic.GeomagneticService = (*DefaultGeomagneticService)(nil)

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
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "geomagnetic"}).Error
	if err != nil {
		log.Errorf("find geomagnetic device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

func (ss *DefaultGeomagneticService) GetGeomagneticDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.GeomagneticDeviceList, err error) {
	return GetGeomagneticDevices(ss.db, ss.l, req)
}

func (ss *DefaultGeomagneticService) GetGeomagneticDevice(customerID uint64, deviceSN string) (geomagneticDevice *protos.GeomagneticDevice, err error) {
	return GetGeomagneticDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultGeomagneticService) GetParkingInfos(req *protos.GetParkingPlacesRequest) (list *protos.ParkingPlaceList, err error) {
	return GetParkingInfos(ss.db, ss.l, req)
}

func (ss *DefaultGeomagneticService) GetParkingInfo(customerID, groupID uint64) (parkingPlace *protos.ParkingPlace, err error) {
	return GetParkingInfo(ss.db, ss.l, customerID, groupID)
}

func (ss *DefaultGeomagneticService) GetParkingHistory(customerID uint64, groupID uint64, date string) (parkingHistory []map[string]interface{}, err error) {
	return GetParkingHistory(ss.db, ss.l, customerID, groupID, date)
}
