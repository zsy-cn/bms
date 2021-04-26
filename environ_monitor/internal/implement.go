package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-12h" // 1小时

type DefaultEnvironMonitorService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewEnvironMonitorService(
	l log.Logger,
	db *gorm.DB,
	cfg environ_monitor.EnvironMonitorConfig,
) (service environ_monitor.EnvironMonitorService, err error) {
	if db == nil {
		err = environ_monitor.ErrDbNotBeNil
		return
	}

	dmService := MustGetDeviceManagementService(db, l)
	deviceTypeID, err = MustGetDeviceTypeID(db, l)
	if err != nil {
		panic(err)
	}
	service = &DefaultEnvironMonitorService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ environ_monitor.EnvironMonitorService = (*DefaultEnvironMonitorService)(nil)

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
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "environ_monitor"}).Error
	if err != nil {
		log.Errorf("find environ_monitor device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

// GetLastEnvironMoniterAverageInfo ...
func (ss *DefaultEnvironMonitorService) GetLastEnvironMoniterAverageInfo(customerID uint64, groupID uint64, datetime string) (envAverageData *protos.EnvironMonitorSectionAverageData, err error) {
	return GetLastEnvironMoniterAverageInfo(ss.db, ss.l, customerID, groupID, datetime)
}

func (ss *DefaultEnvironMonitorService) GetEnvironMonitorDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.EnvironMonitorDeviceList, err error) {
	return GetEnvironMonitorDevices(ss.db, ss.l, req)
}

func (ss *DefaultEnvironMonitorService) GetEnvironMonitorDevice(customerID uint64, deviceSN string) (device *protos.EnvironMonitorDevice, err error) {
	return GetEnvironMonitorDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultEnvironMonitorService) GetEnvironMonitorDeviceStatus(customerID uint64, deviceSN string) (status uint64, err error) {
	return GetEnvironMonitorDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultEnvironMonitorService) GetEnvironMonitorSectionAverageData(customerID uint64, groupID uint64, dateFrom string, dateTo string, delta uint64) (historyDataList []map[string]interface{}, err error) {
	return GetEnvironMonitorSectionAverageData(ss.db, ss.l, customerID, groupID, dateFrom, dateTo, delta)
}
