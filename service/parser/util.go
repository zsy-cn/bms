package parser

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// ConsulParserService ...
// 提供HealthCheck公共方法
type ConsulParserService struct{}

// HealthCheck ...
func (cs *ConsulParserService) HealthCheck(empty *protos.Empty) (resp *protos.ParserHealthCheckResponse, err error) {
	resp = &protos.ParserHealthCheckResponse{
		Msg: "success",
	}
	return
}

// FillBaseAppModel 根据devEUI查询设备ID, 填充UplinkMsg模型
func FillBaseAppModel(db *gorm.DB, uplinkMsg *protos.ParserHubUplinkMsg, uplinkMsgPayload *model.UplinkMsgPayload, baseAppModel *model.App) (err error) {
	sensorRecord := &model.Sensor{}
	whereArgs := map[string]interface{}{
		"dev_eui": uplinkMsg.DevEUI,
	}
	err = db.Where(whereArgs).First(sensorRecord).Error
	if err != nil {
		if err.Error() != "record not found" {
			logger.Errorf("query sensor with dev_eui %s failed: %s", uplinkMsg.DevEUI, err.Error())
			return
		}
		return
	}
	deviceRecord := &model.Device{}
	err = db.Where(&model.Device{SerialNumber: sensorRecord.DeviceSN}).First(deviceRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warnf("can not find sensor device by sn %s, you should repair it", sensorRecord.DeviceSN)
		} else {
			logger.Errorf("query sensor device by sn %s failed: %s", sensorRecord.DeviceSN, err.Error())
		}
		return
	}
	baseAppModel.DeviceSN = sensorRecord.DeviceSN
	baseAppModel.FPort = uint8(uplinkMsg.FPort)
	baseAppModel.RawData = uplinkMsg.Data
	baseAppModel.Data = uplinkMsgPayload
	baseAppModel.GroupID = deviceRecord.GroupID
	baseAppModel.CustomerID = deviceRecord.CustomerID
	return
}
