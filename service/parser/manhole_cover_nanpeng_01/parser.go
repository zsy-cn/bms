package manhole_cover_nanpeng_01

import (
	"github.com/zsy-cn/bms/model"
)

func parser00(data []byte, payload *model.UplinkMsgPayload) {
	payload.Attributes = model.UplinkMsgAttribute{}

	voltage := float64(data[3])
	voltageAttr := &model.UplinkMsgPayloadField{
		DisplayName: "电压",
		Unit:        "",
		Value:       float64(voltage) / 10,
	}
	payload.Attributes["Voltage"] = voltageAttr

	status := data[4] & 0x01
	statusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "井盖状态",
		Unit:        "",
		Value:       0,
	}
	if status > 0 {
		statusAttr.Value = 1
	}
	payload.Attributes["Status"] = statusAttr

	waterLevelStatus := data[4] & 0x02
	waterLevelStatusAttr := &model.UplinkMsgPayloadField{
		DisplayName: "水位状态",
		Unit:        "",
		Value:       0,
	}
	if waterLevelStatus > 0 {
		waterLevelStatusAttr.Value = 1
	}

	payload.Attributes["WaterLevelStatus"] = waterLevelStatusAttr
}
