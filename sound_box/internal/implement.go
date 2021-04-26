package internal

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/device_management"
	dapp "github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultSoundBoxService struct {
	l         log.Logger
	db        *gorm.DB
	dmService device_management.DeviceManagementService
}

func NewSoundBoxService(
	l log.Logger,
	db *gorm.DB,
	cfg sound_box.SoundBoxConfig,
) (service sound_box.SoundBoxService, err error) {
	if db == nil {
		err = sound_box.ErrDbNotBeNil
		return
	}
	dmService := MustGetDeviceManagementService(db, l)
	deviceTypeID, err = MustGetDeviceTypeID(db, l)
	if err != nil {
		panic(err)
	}

	service = &DefaultSoundBoxService{
		l:         l,
		db:        db,
		dmService: dmService,
	}

	return service, nil
}

var _ sound_box.SoundBoxService = (*DefaultSoundBoxService)(nil)
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
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "sound_box"}).Error
	if err != nil {
		log.Errorf("find manholeCover device type failed: in MustGetDeviceTypeID(): %s", err.Error())
		return
	}
	id = deviceTypeRecord.ID
	return
}

func (ss *DefaultSoundBoxService) GetDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.SoundBoxDeviceList, err error) {
	return GetDevices(ss.db, ss.l, req)
}

func (ss *DefaultSoundBoxService) GetDevice(customerID uint64, deviceSN string) (soundboxDevice *protos.SoundBoxDevice, err error) {
	return GetDevice(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultSoundBoxService) GetSoundBoxDeviceStatus(customerID uint64, deviceSN string) (status int8, err error) {
	return GetSoundBoxDeviceStatus(ss.db, ss.l, customerID, deviceSN)
}

func (ss *DefaultSoundBoxService) GetSoundBoxDeviceGroups(req *protos.GetDevicesRequestForCustomer) (resp *protos.SoundBoxDeviceGroupList, err error) {
	return GetSoundBoxDeviceGroups(ss.db, ss.l, req)
}

func (ss *DefaultSoundBoxService) SaveMediaFile(customerID uint64, name string, path string, duration float64, size uint64) (err error) {
	return SaveMediaFile(ss.db, ss.l, customerID, name, path, duration, size)
}

func (ss *DefaultSoundBoxService) GetSoundBoxMedias(req *protos.GetSoundBoxMediasRequest) (resp *protos.GetSoundBoxMediasResponse, err error) {
	return GetSoundBoxMedias(ss.db, ss.l, req)
}

func (ss *DefaultSoundBoxService) UpdateSoundBoxMedia(req *protos.UpdateSoundBoxMediaRequest) (err error) {
	return UpdateSoundBoxMedia(ss.db, ss.l, req)
}

func (ss *DefaultSoundBoxService) DeleteSoundBoxMedia(id, customerID uint64) (err error) {
	return DeleteSoundBoxMedia(ss.db, ss.l, id, customerID)
}
