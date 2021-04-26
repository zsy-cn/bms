package trashcan_lierda_01

import (
	"github.com/zsy-cn/bms/model"
)

func parser03(data []byte, payload *model.UplinkMsgPayload) {
	distanceAttr := &model.UplinkMsgPayloadField{
		DisplayName: "版本",
		Unit:        "",
		Value:       0, // 这个值不知道怎么算
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"Version": distanceAttr,
	}
}

func parser17(data []byte, payload *model.UplinkMsgPayload) {
	configAttr := &model.UplinkMsgPayloadField{
		DisplayName: "配置命令ID",
		Unit:        "",
		Value:       float64(data[3]),
	}

	param1 := int(data[4])<<8 + int(data[5])
	param1Attr := &model.UplinkMsgPayloadField{
		DisplayName: "参数1",
		Unit:        "",
		Value:       float64(param1),
	}
	param2 := int(data[6])<<8 + int(data[7])
	param2Attr := &model.UplinkMsgPayloadField{
		DisplayName: "参数2",
		Unit:        "",
		Value:       float64(param2),
	}
	param3 := int(data[8])<<8 + int(data[9])
	param3Attr := &model.UplinkMsgPayloadField{
		DisplayName: "参数3",
		Unit:        "",
		Value:       float64(param3),
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"ConfigID": configAttr,
		"Param1":   param1Attr,
		"Param2":   param2Attr,
		"Param3":   param3Attr,
	}
}

func parser31(data []byte, payload *model.UplinkMsgPayload) {
	isFullAttr := &model.UplinkMsgPayloadField{
		DisplayName: "垃圾桶状态",
		Unit:        "",
		Value:       0,
	}
	if int(data[3]) > 0 {
		isFullAttr.Value = 1
	}
	distanceAttr := &model.UplinkMsgPayloadField{
		DisplayName: "距离",
		Unit:        "cm",
		Value:       float64(data[4]),
	}

	temperatureAttr := &model.UplinkMsgPayloadField{
		DisplayName: "温度",
		Unit:        "C",
		Value:       float64(data[5]),
	}

	batteryAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电量",
		Unit:        "%",
		Value:       float64(data[6]),
	}

	rssiAttr := &model.UplinkMsgPayloadField{
		DisplayName: "信号强度",
		Unit:        "",
		Value:       float64(data[7]),
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"Status":      isFullAttr,
		"Distance":    distanceAttr,
		"Temperature": temperatureAttr,
		"Battery":     batteryAttr,
		"Rssi":        rssiAttr,
	}
}

func parser32(data []byte, payload *model.UplinkMsgPayload) {
	angleAttr := &model.UplinkMsgPayloadField{
		DisplayName: "上盖倾斜角度",
		Unit:        "",
		Value:       float64(data[3]),
	}

	distanceAttr := &model.UplinkMsgPayloadField{
		DisplayName: "距离",
		Unit:        "cm",
		Value:       float64(data[4]),
	}

	temperatureAttr := &model.UplinkMsgPayloadField{
		DisplayName: "温度",
		Unit:        "%",
		Value:       float64(data[5]), // 可能为0值
	}

	batteryAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电量",
		Unit:        "%",
		Value:       float64(data[6]),
	}

	rssiAttr := &model.UplinkMsgPayloadField{
		DisplayName: "信号强度",
		Unit:        "",
		Value:       float64(data[7]),
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"Angle":       angleAttr,
		"Distance":    distanceAttr,
		"Temperature": temperatureAttr,
		"Battery":     batteryAttr,
		"Rssi":        rssiAttr,
	}
}

func parser33(data []byte, payload *model.UplinkMsgPayload) {
	rawLongitude := (int(data[3]) << 24) + (int(data[4]) << 16) + (int(data[5]) << 8) + int(data[6])
	longitudeAttr := &model.UplinkMsgPayloadField{
		DisplayName: "垃圾桶经度",
		Unit:        "",
		Value:       float64(rawLongitude) / 10000,
	}

	rawLatitude := (int(data[7]) << 24) + (int(data[8]) << 16) + (int(data[9]) << 8) + int(data[10])
	latitudeAttr := &model.UplinkMsgPayloadField{
		DisplayName: "垃圾桶纬度",
		Unit:        "",
		Value:       float64(rawLatitude) / 10000,
	}

	batteryAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电量",
		Unit:        "%",
		Value:       float64(data[11]),
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"Longitude": longitudeAttr,
		"Latitude":  latitudeAttr,
		"Battery":   batteryAttr,
	}
}

func parser34(data []byte, payload *model.UplinkMsgPayload) {
	openAttr := &model.UplinkMsgPayloadField{
		DisplayName: "上盖状态", // (1打开,0关闭)
		Unit:        "",
		Value:       0,
	}
	if int(data[3]) > 0 {
		openAttr.Value = 1
	}

	angleAttr := &model.UplinkMsgPayloadField{
		DisplayName: "上盖倾斜角度",
		Unit:        "",
		Value:       float64(data[4]),
	}

	batteryAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电量",
		Unit:        "%",
		Value:       float64(data[5]),
	}

	rssiAttr := &model.UplinkMsgPayloadField{
		DisplayName: "信号强度",
		Unit:        "",
		Value:       float64(data[6]),
	}

	payload.Attributes = model.UplinkMsgAttribute{
		"Status":  openAttr,
		"Angle":   angleAttr,
		"Battery": batteryAttr,
		"Rssi":    rssiAttr,
	}
}
