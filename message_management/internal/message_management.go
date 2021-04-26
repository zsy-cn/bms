package internal

import (
	"errors"
	"time"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
)

// GetNotifications ...
func GetNotifications(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (notificationList *protos.NotificationList, err error) {
	notificationList = &protos.NotificationList{
		List:  []*protos.NotificationInfo{},
		Count: 0,
	}
	query := db.Model(&model.Notification{})
	// ...这里应是条件查询语句
	whereArgs := map[string]interface{}{}
	if req.CustomerID == 0 {
		err = errors.New("please specify the customer id, you can not get all notifications")
		log.Error(err)
		return
	}
	whereArgs["customer_id"] = req.CustomerID

	if req.DeviceTypeID > 0 {
		whereArgs["device_type_id"] = req.DeviceTypeID
	}
	if req.GroupID > 0 {
		whereArgs["group_id"] = req.GroupID
	}
	if req.SerialNumber != "" {
		whereArgs["device_sn"] = req.SerialNumber
	}
	query = query.Where(whereArgs)

	var count uint64
	// 首先得到count总量
	err = query.Count(&count).Error
	if err != nil {
		log.Errorf("find notifications count failed: %s", err.Error())
		return
	}
	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)
	notificationRecords := []*model.Notification{}
	err = query.Find(&notificationRecords).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			log.Errorf("find notifications failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}
	for _, notificationRecord := range notificationRecords {
		notificationInfoPb := &protos.NotificationInfo{}
		err = model2Pb(db, log, notificationRecord, notificationInfoPb)
		if err != nil {
			return
		}
		notificationList.List = append(notificationList.List, notificationInfoPb)
	}
	notificationList.Count = count
	notificationList.TotalCount = count
	notificationList.CurrentPage = req.Pagination.Page
	notificationList.PageSize = req.Pagination.PageSize
	return
}

// DiscardNotification ...
func DiscardNotification(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (err error) {
	whereArgs := map[string]interface{}{
		"customer_id": customerID,
		"device_sn":   deviceSN,
	}
	newColumns := map[string]interface{}{
		"solved":    true,
		"solved_at": time.Now(),
	}
	err = db.Model(&model.Notification{}).Where(whereArgs).Updates(newColumns).Error
	if err != nil {
		log.Errorf("update notification status failed in DiscardNotification(): %s", err.Error())
		return
	}
	return
}

func model2Pb(db *gorm.DB, log log.Logger, record *model.Notification, pb *protos.NotificationInfo) (err error) {
	pb.ID = record.ID
	pb.Content = record.Content

	deviceRecord := &model.Device{}
	err = db.Where(&model.Device{SerialNumber: record.DeviceSN}).First(deviceRecord).Error
	if err != nil {
		log.Errorf("find device failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}
	pb.DeviceSerialNumber = record.DeviceSN
	pb.DeviceName = deviceRecord.Name

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, record.DeviceTypeID).Error
	if err != nil {
		log.Errorf("find device type failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}
	pb.DeviceTypeID = deviceTypeRecord.ID
	pb.DeviceType = deviceTypeRecord.Name

	groupRecord := &model.Group{}
	err = db.First(groupRecord, record.GroupID).Error
	if err != nil {
		log.Errorf("find device group failed: in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}
	pb.GroupID = groupRecord.ID
	pb.Group = groupRecord.Name

	loc, _ := time.LoadLocation("Asia/Shanghai")
	pb.CreatedAt = record.CreatedAt.In(loc).Format("2006-01-02 15:04:05")
	pb.SolvedAt = record.SolvedAt.In(loc).Format("2006-01-02 15:04:05")

	if record.Solved {
		pb.Solved = "true"
	} else {
		pb.Solved = "false"
	}

	return
}
