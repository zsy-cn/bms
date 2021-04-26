package geomagnetic_vchuan_01

import (
	"github.com/zsy-cn/bms/model"
)

func parser00(data []byte, payload *model.UplinkMsgPayload) {
	payload.Attributes = model.UplinkMsgAttribute{}

	isUsed := int(data[2])
	statusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "停车状态",
		Unit:        "",
		Value:       0,
	}
	if isUsed == 1 {
		statusAttr.Value = 1
	}
	payload.Attributes["Used"] = statusAttr

	sysStatus := int(data[3])<<8 + int(data[4])
	sysStatusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "系统运行状态",
		Unit:        "",
		Value:       float64(sysStatus),
	}
	payload.Attributes["SysStatus"] = sysStatusAttr

	msgOrder := int(data[5])<<8 + int(data[6])
	msgOrderAttr := &model.UplinkMsgPayloadField{
		DisplayName: "消息序号",
		Unit:        "",
		Value:       float64(msgOrder),
	}
	payload.Attributes["MsgOrder"] = msgOrderAttr

	voltage := int(data[7])<<8 + int(data[8])
	voltageAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电压",
		Unit:        "",
		Value:       float64(voltage) / 100,
	}
	payload.Attributes["Voltage"] = voltageAttr

	deviceID := int(data[9])<<24 + int(data[10])<<16 + int(data[11])<<8 + int(data[12])
	deviceIDAttr := &model.UplinkMsgPayloadField{
		DisplayName: "设备ID",
		Unit:        "",
		Value:       float64(deviceID),
	}
	payload.Attributes["DeviceID"] = deviceIDAttr

	rssi := int(data[13])<<8 + int(data[14])
	rssiAttr := &model.UplinkMsgPayloadField{
		DisplayName: "信号强度",
		Unit:        "",
		Value:       float64(rssi),
	}
	payload.Attributes["Rssi"] = rssiAttr

	snr := int(data[15])<<8 + int(data[16])
	snrAttr := &model.UplinkMsgPayloadField{
		DisplayName: "信噪比",
		Unit:        "",
		Value:       float64(snr),
	}
	payload.Attributes["SNR"] = snrAttr

	// xyzParser 解析包中XYZ三轴数据
	xAxis := int(data[17])<<8 + int(data[18])
	xAttr := &model.UplinkMsgPayloadField{
		DisplayName: "X轴",
		Unit:        "",
		Value:       float64(xAxis),
	}
	payload.Attributes["XAxis"] = xAttr

	yAxis := int(data[19])<<8 + int(data[20])
	yAttr := &model.UplinkMsgPayloadField{
		DisplayName: "Y轴",
		Unit:        "",
		Value:       float64(yAxis),
	}
	payload.Attributes["YAxis"] = yAttr

	zAxis := int(data[21])<<8 + int(data[22])
	zAttr := &model.UplinkMsgPayloadField{
		DisplayName: "Z轴",
		Unit:        "",
		Value:       float64(zAxis),
	}
	payload.Attributes["ZAxis"] = zAttr
}
