package parser

import (
	"fmt"
	"math"
	"time"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/trashcan"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/model"
)

func reportNotification(db *gorm.DB, notificationModel *model.Notification) (err error) {
	logger.Debugf("report notification model: %+v", notificationModel)

	now := time.Now()
	m15, _ := time.ParseDuration("-15m") // 暂定为15分钟
	deadline := now.Add(m15)
	whereStr := "device_sn = ? and key = ? and updated_at > ? and solved = false"

	// 此设备最近是否有报警记录, 如果有, 直拉在该报警记录次数加1, 没有则新建.
	// 因为一旦报警, 可能触发多条上行信息, 对用户很不友好.
	notificationRecord := &model.Notification{}
	err = db.Where(whereStr, notificationModel.DeviceSN, notificationModel.Key, deadline).Order("id desc").First(notificationRecord).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() == "record not found" {
			err = nil
		} else {
			logger.Errorf("find notifications for device: %s failed in ReportManholeCoverNotification: %s", notificationModel.DeviceSN, err.Error())
			return
		}
	}
	// 新建通知记录
	if notificationRecord.ID == 0 {
		logger.Debugf("try to create new notification model: %+v", notificationModel)
		err = db.Create(notificationModel).Error
		if err != nil {
			logger.Errorf("create notification record for msg: %d failed in ReportManholeCoverNotification(): %s", notificationModel.MsgID, err.Error())
			return
		}
	} else {
		logger.Debugf("try to update notification record: %+v", notificationRecord)

		newColumns := map[string]interface{}{
			"msg_id":     notificationModel.MsgID,
			"updated_at": now,
			"times":      notificationRecord.Times + 1,
		}
		err = db.Model(notificationModel).Updates(newColumns).Error
		if err != nil {
			logger.Errorf("update notification for id: %d failed in ReportManholeCoverNotification(): %s", notificationRecord.ID, err.Error())
			return
		}
	}
	return
}

// ReportManholeCoverNotification 尝试创建异常通知
func ReportManholeCoverNotification(db *gorm.DB, manholeCoverModel *manhole_cover.ManholeCover) (err error) {
	logger.Debugf("report manhole cover notification for model: %+v", manholeCoverModel)

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "manhole_cover"}).Error
	if err != nil {
		logger.Errorf("find manhole_cover device type failed: in ReportManholeCoverNotification(): %s", err.Error())
		return
	}
	// 井盖设备无需指定规则, 只要打开就报警
	notificationModel := &model.Notification{
		DeviceSN:     manholeCoverModel.DeviceSN,
		MsgID:        manholeCoverModel.ID,
		GroupID:      manholeCoverModel.GroupID,
		CustomerID:   manholeCoverModel.CustomerID,
		DeviceTypeID: deviceTypeRecord.ID,
		Status:       3,
		Times:        1,
	}
	// 井盖设备有两种报警情况, 正常的话无操作
	if manholeCoverModel.Status != 0 {
		key := "manhole_cover_open_alert"
		content := "井盖报警"
		notificationModel.Key = key
		notificationModel.Content = content
		reportNotification(db, notificationModel)
	}
	if manholeCoverModel.WaterLevelStatus != 0 {
		key := "manhole_cover_water_alert"
		content := "井盖水位报警"
		notificationModel.Key = key
		notificationModel.Content = content
		reportNotification(db, notificationModel)
	}
	return
}

// ReportTrashcanNotification 尝试创建异常通知
func ReportTrashcanNotification(db *gorm.DB, trashcanModel *trashcan.Trashcan) (err error) {
	logger.Debugf("report trashcan notification for model: %+v", trashcanModel)

	deviceTypeRecord := &model.DeviceType{}
	err = db.First(deviceTypeRecord, &model.DeviceType{Key: "trashcan"}).Error
	if err != nil {
		logger.Errorf("find trashcan device type failed: in ReportTrashcanNotification(): %s", err.Error())
		return
	}

	notificationModel := &model.Notification{
		DeviceSN:     trashcanModel.DeviceSN,
		MsgID:        trashcanModel.ID,
		GroupID:      trashcanModel.GroupID,
		CustomerID:   trashcanModel.CustomerID,
		DeviceTypeID: deviceTypeRecord.ID,
		Key:          "trashcan_usage",
		Status:       3,
		Times:        1,
	}
	ruleRecord := &model.Rule{}
	whereArgs := map[string]interface{}{
		"customer_id": trashcanModel.CustomerID,
		"device_sn":   trashcanModel.DeviceSN,
	}
	err = db.Where(whereArgs).First(ruleRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			logger.Errorf("find trashcan rule failed for %s: in ReportTrashcanNotification(): %s", trashcanModel.DeviceSN, err.Error())
			return
		}
	}
	// 如果没有对应的规则, 则结束
	if ruleRecord.ID == 0 {
		logger.Infof("didn't find rule for trashcan device: %s, return", trashcanModel.DeviceSN)
		return
	}
	logger.Infof("find rule for trashcan device: %+v", ruleRecord)

	notificationModel.RuleID = ruleRecord.ID
	percent := float64(int(math.Abs((conf.TrashcanHeight - trashcanModel.Percent) / conf.TrashcanHeight * 100)))
	baseContent := "使用率高于 %d%%"
	if percent >= []float64(*ruleRecord.Section)[2] {
		notificationModel.Content = fmt.Sprintf(baseContent, int([]float64(*ruleRecord.Section)[2]))
		notificationModel.Status = 3
	} else if percent >= []float64(*ruleRecord.Section)[1] {
		notificationModel.Content = fmt.Sprintf(baseContent, int([]float64(*ruleRecord.Section)[1]))
		notificationModel.Status = 2
	} else if percent >= []float64(*ruleRecord.Section)[0] {
		notificationModel.Content = fmt.Sprintf(baseContent, int([]float64(*ruleRecord.Section)[0]))
		notificationModel.Status = 1
	} else {
		return
	}

	return reportNotification(db, notificationModel)
}
