package manhole_cover_gti_01

import "github.com/zsy-cn/bms/model"

func parser00(data []byte, payload *model.UplinkMsgPayload) {
	payload.Attributes = model.UplinkMsgAttribute{}

	var manholeCoverAlert float64
	if data[19] == 0x66 {
		manholeCoverAlert = 1
	}
	voltage := int(data[20])<<8 + int(data[21])
	payload.Attributes = model.UplinkMsgAttribute{
		"Status": &model.UplinkMsgPayloadField{
			DisplayName: "井盖报警",
			Unit:        "",
			Value:       manholeCoverAlert,
		},
		"Voltage": &model.UplinkMsgPayloadField{
			DisplayName: "井盖电压",
			Unit:        "V",
			Value:       float64(voltage) / 1000,
		},
	}
}
