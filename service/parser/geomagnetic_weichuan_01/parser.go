package geomagnetic_weichuan_01

import (
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/service/parser"
)

func parser01(data []byte, payload *model.UplinkMsgPayload) {
	dayStart := data[1]
	dayEnd := data[2]
	dayCheckInterval := data[3] & 0x01
	daySensibility := (data[3] & 06) >> 1
	dayReportInterval := (data[3] & 0xe0) >> 5
	nightCheckInterval := data[4] & 0x01
	nightSensibility := (data[4] & 06) >> 1
	nightReportInterval := (data[4] & 0xe0) >> 5

	payload.Attributes = model.UplinkMsgAttribute{
		"DayStart": &model.UplinkMsgPayloadField{
			Unit:  "时",
			Value: float64(dayStart),
		},
		"DayEnd": &model.UplinkMsgPayloadField{
			Unit:  "时",
			Value: float64(dayEnd),
		},
		"DayCheckInterval": &model.UplinkMsgPayloadField{
			Unit:  "秒",
			Value: float64(dayCheckInterval),
		},
		"DaySensibility": &model.UplinkMsgPayloadField{
			Unit:  "",
			Value: float64(daySensibility),
		},
		"DayReportInterval": &model.UplinkMsgPayloadField{
			Unit:  "",
			Value: float64(dayReportInterval),
		},
		"NightCheckInterval": &model.UplinkMsgPayloadField{
			Unit:  "秒",
			Value: float64(nightCheckInterval),
		},
		"NightSensibility": &model.UplinkMsgPayloadField{
			Unit:  "秒",
			Value: float64(nightSensibility),
		},
		"NightReportInterval": &model.UplinkMsgPayloadField{
			Unit:  "秒",
			Value: float64(nightReportInterval),
		},
	}
}

func parser02(data []byte, payload *model.UplinkMsgPayload) {
	normalInfoParser(payload.Attributes, data[:3])

	rawFirmwareVersion := int(data[4])<<8 + int(data[5])
	rawSoftwareVersion := int(data[6])<<16 + int(data[7])<<8 + int(data[8])

	payload.Attributes = model.UplinkMsgAttribute{
		"Firmware": &model.UplinkMsgPayloadField{
			DisplayName: "固件版本",
			Unit:        "",
			Value:       float64(rawFirmwareVersion),
		},
		"Software": &model.UplinkMsgPayloadField{
			DisplayName: "软件版本",
			Unit:        "",
			Value:       float64(rawSoftwareVersion),
		},
	}
}

func parser03(data []byte, payload *model.UplinkMsgPayload) {
	normalInfoParser(payload.Attributes, data[1:4])
}

func parser04(data []byte, payload *model.UplinkMsgPayload) {
	isUsed := (data[1] & 0x80) >> 7

	statusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "停车状态",
		Unit:        "",
		Value:       0,
	}
	if isUsed == 1 {
		statusAttr.Value = 1
	}
	payload.Attributes["Used"] = statusAttr

	rawVoltage := [1]byte{data[1] & 0x7f}
	parser.VoltageDecoder(payload.Attributes, rawVoltage, "")
}

func parser07(data []byte, payload *model.UplinkMsgPayload) {
	xyzParser(payload.Attributes, data[1:7])
}

func parser08(data []byte, payload *model.UplinkMsgPayload) {
	xyzParser(payload.Attributes, data[1:7])
	newInfoParser(payload.Attributes, data[7:10])
}

func parser0a(data []byte, payload *model.UplinkMsgPayload) {
	newInfoParser(payload.Attributes, data[7:10])
}

func parser09(data []byte, payload *model.UplinkMsgPayload) {
	payload.Attributes["Error"] = &model.UplinkMsgPayloadField{
		DisplayName: "地磁设备出错",
		Unit:        "",
		Value:       1, // 只有一种错误: 未检测到传感器
	}
}

// normalInfoParser 处理正常包的前3个字节
func normalInfoParser(attributeMap model.UplinkMsgAttribute, data []byte) {
	isUsed := (data[0] & 0x80) >> 7

	statusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "停车状态",
		Unit:        "",
		Value:       0,
	}
	if isUsed == 1 {
		statusAttr.Value = 1
	}

	attributeMap["Used"] = statusAttr

	rawVoltage := [1]byte{data[0] & 0x7f}
	parser.VoltageDecoder(attributeMap, rawVoltage, "")

	rawTemperature := [2]byte{data[1], data[2]}
	parser.TemperatureDecoder(attributeMap, rawTemperature, "")
}

// newInfoParser 处理新版包的后3个字节
func newInfoParser(attributeMap model.UplinkMsgAttribute, data []byte) {
	rawTemperature := [2]byte{data[0], data[1]}
	parser.TemperatureDecoder(attributeMap, rawTemperature, "")

	isUsed := (data[2] & 0x80) >> 7

	statusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "停车状态",
		Unit:        "",
		Value:       0,
	}
	if isUsed == 1 {
		statusAttr.Value = 1
	}
	attributeMap["Used"] = statusAttr

	rawVoltage := [1]byte{data[2] & 0x7f}
	parser.VoltageDecoder(attributeMap, rawVoltage, "")
}

// xyzParser 解析包中XYZ三轴数据
func xyzParser(attributeMap model.UplinkMsgAttribute, data []byte) {
	xAxis := int(data[0])<<8 + int(data[1])
	xAttr := &model.UplinkMsgPayloadField{
		DisplayName: "X轴",
		Unit:        "",
		Value:       float64(xAxis),
	}
	attributeMap["XAxis"] = xAttr

	yAxis := int(data[2])<<8 + int(data[3])
	yAttr := &model.UplinkMsgPayloadField{
		DisplayName: "Y轴",
		Unit:        "",
		Value:       float64(yAxis),
	}
	attributeMap["YAxis"] = yAttr

	zAxis := int(data[4])<<8 + int(data[5])
	zAttr := &model.UplinkMsgPayloadField{
		DisplayName: "Z轴",
		Unit:        "",
		Value:       float64(zAxis),
	}
	attributeMap["ZAxis"] = zAttr
}
