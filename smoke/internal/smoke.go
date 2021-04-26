package internal

import (
	"time"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/smoke"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
)

// GetLastSmokeDeviceInfos ...
func GetLastSmokeDeviceInfos(db *gorm.DB, log log.Logger, sns []string) (msg []smoke.Smoke, err error) {
	if len(sns) <= 0 {
		return nil, smoke.ErrIdNotBeNil
	}
	// 按照device_sn分组, 查询出最新的消息, 这条语句返回的记录只包含两个字段: device_sn和created_at
	sqlStr3 := "select device_sn, max(created_at) as created_at from smokes group by device_sn"
	// 将上述分组查询作为中间表, 与原表进行连接查询, 以查出表中的所有字段.
	sqlStr2 := "select a.* from smokes as a inner join (" + sqlStr3 + ") as b on a.device_sn = b.device_sn and a.created_at = b.created_at"
	// 主查询, 从所有不同设备的最新消息记录中取出目标设备的记录
	sqlStr1 := "select c.* from (" + sqlStr2 + ") as c where device_sn in (?)"
	err = db.Raw(sqlStr1, sns).Scan(&msg).Error
	if err != nil {
		panic(err)
	}
	return
}

// GetLastSmokeDeviceInfo ...
func GetLastSmokeDeviceInfo(db *gorm.DB, log log.Logger, sn string) (msg *smoke.Smoke, err error) {
	snArray := []string{sn}
	msgs, err := GetLastSmokeDeviceInfos(db, log, snArray)
	if err != nil {
		return nil, smoke.ErrIdNotBeNil
	}
	if len(msgs) > 0 {
		msg = &msgs[0]
	}
	return
}

// GetSmokeDevices ...
func GetSmokeDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "smoke"}).Error
	if err != nil {
		log.Errorf("find smoke device type failed: in GetSmokeDevices(): %s", err.Error())
		return
	}
	req.DeviceTypeID = deviceTypeRecord.ID
	return _deviceManagementService.GetDevices(req)
}

// GetSmokeDevice ...
func GetSmokeDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (smokeDevice *protos.Device, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "smoke"}).Error
	if err != nil {
		log.Errorf("find smoke device type failed: in GetSmokeDevice(): %s", err.Error())
		return
	}
	return _deviceManagementService.GetDevice(deviceSN, customerID, deviceTypeRecord.ID)
}

// GetSmokeDeviceStatus ...
// status: -1: 设备不存在, 1: 在线, 2: 离线, 3: 其他
func GetSmokeDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status int8, err error) {
	_, err = GetSmokeDevice(db, log, customerID, deviceSN)
	if err != nil {
		status = -1
		return
	}
	msg := &smoke.Smoke{}
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline := now.Add(d2)
	whereArgs := "device_sn = ? and created_at > ?"
	err = db.Where(whereArgs, deviceSN, deadline).First(msg).Error
	if err != nil {
		err = nil
		status = 2 // 离线
		return
	}
	status = 1
	return
}
