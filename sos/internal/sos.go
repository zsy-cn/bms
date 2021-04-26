package internal

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sos"
	"github.com/zsy-cn/bms/util/log"
)

/********************************************************************************************************/
// GetLastSOSDeviceInfos 获取指定设备列表的最近一条消息
// 注意: 离线设备将无法获取到信息, 主要用于查询在线设备数量统计信息
func GetLastSOSDeviceInfos(db *gorm.DB, log log.Logger, sns []string) (msg []*sos.Sos, err error) {
	if len(sns) <= 0 {
		return nil, sos.ErrIdNotBeNil
	}
	// 按照device_sn分组, 查询出最新的消息, 这条语句返回的记录只包含两个字段: device_sn和created_at
	// 将上述分组查询作为中间表, 与原表进行连接查询, 以查出表中的所有字段.
	// 主查询, 从所有不同设备的最新消息记录中取出目标设备的记录
	sqlStr := `
		select c.* from (select a.* from sos as a inner join (select device_sn, max(created_at) as created_at from sos group by device_sn) as b on a.device_sn = b.device_sn and a.created_at = b.created_at) as c where device_sn in (?)
	`
	err = db.Raw(sqlStr, sns).Scan(&msg).Error
	if err != nil {
		panic(err)
	}
	return
}

// GetLastSOSDeviceInfo ...
func GetLastSOSDeviceInfo(db *gorm.DB, log log.Logger, sns string) (msg *sos.Sos, err error) {
	snsArray := []string{sns}
	msgs, err := GetLastSOSDeviceInfos(db, log, snsArray)
	if err != nil {
		return nil, sos.ErrIdNotBeNil
	}
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return
}

// GetSOSDevices ...
func GetSOSDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (deviceList *protos.DeviceList, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "sos"}).Error
	if err != nil {
		log.Errorf("find sos device type failed: in GetSOSDevices(): %s", err.Error())
		return
	}
	req.DeviceTypeID = deviceTypeRecord.ID
	return _deviceManagementService.GetDevices(req)
}

// GetSOSDevice ...
func GetSOSDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (sosDevice *protos.Device, err error) {
	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "environ_monitor"}).Error
	if err != nil {
		log.Errorf("find environ_monitor device type failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}
	return _deviceManagementService.GetDevice(deviceSN, customerID, deviceTypeRecord.ID)
}

// GetSOSDeviceStatus ...
// status: -1: 设备不存在, 1: 在线, 2: 离线; 3: 报警, 4: 其他
func GetSOSDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (code int8, msgID uint64, err error) {
	_, err = GetSOSDevice(db, log, customerID, deviceSN)
	if err != nil {
		code = -1
		return
	}
	msg := &sos.Sos{}
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline := now.Add(d2)
	whereArgs := "device_sn = ? and created_at > ?"
	err = db.Where(whereArgs, deviceSN, deadline).Order("id desc").First(msg).Error
	if err != nil {
		err = nil
		code = 2 // 离线
		return
	}
	// 接下来处理设备在线的情况, 要查看该设备是否有告警信息, 如果有告警, 返回告警记录id.
	if msg.Status != 1 {
		notification := &model.Notification{}
		err = db.Where("device_sn = ? and msg_id = ?", deviceSN, msg.ID).First(notification).Error
		if err != nil {
			log.Errorf("find notification for device: %s message: %d failed: %s", deviceSN, msg.ID, err.Error())
			return
		}
		return 3, notification.ID, nil
	}
	return 1, 0, nil
}
