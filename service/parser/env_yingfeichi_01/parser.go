package env_yingfeichi_01

import (
	"github.com/zsy-cn/bms/model"
)

var taskList4WeichuanENV01 = []func([]int, *model.UplinkMsgPayload){
	part01, part02, part03, part04, part05, part06, part07, part08, part09, part10,
	part11, part12, part13, part14, part15, part16, part17, part18, part19, part20,
	part21, part22, part23, part24, part25, part26, part27, part28, part29,
}

func part01(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[0]<<8 + intData[1]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Temperature"] = &model.UplinkMsgPayloadField{
		DisplayName: "温度",
		Unit:        "C",
		Value:       float64(rawVal) / 10,
	}
}
func part02(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[2]<<8 + intData[3]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Humidity"] = &model.UplinkMsgPayloadField{
		DisplayName: "湿度",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part03(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[4]<<8 + intData[5]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["NH3"] = &model.UplinkMsgPayloadField{
		DisplayName: "氨气含量", // NH3
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part04(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[6]<<8 + intData[7]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["O2"] = &model.UplinkMsgPayloadField{
		DisplayName: "氧气含量", // O2
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part05(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[8]<<8 + intData[9]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["H2S"] = &model.UplinkMsgPayloadField{
		DisplayName: "硫化氢含量", // H2S
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part06(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[10]<<8 + intData[11]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["CH4"] = &model.UplinkMsgPayloadField{
		DisplayName: "甲烷含量", // CH4
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part07(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[12]<<8 + intData[13]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SoilTemperature"] = &model.UplinkMsgPayloadField{
		DisplayName: "土壤温度",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part08(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[14]<<8 + intData[15]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SoilHumidity"] = &model.UplinkMsgPayloadField{
		DisplayName: "土壤湿度",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part09(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[16]<<8 + intData[17]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["AirPressure"] = &model.UplinkMsgPayloadField{
		DisplayName: "气压",
		Unit:        "千帕",
		Value:       float64(rawVal) / 10,
	}
}
func part10(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[18]<<8 + intData[19]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["WindSpeed"] = &model.UplinkMsgPayloadField{
		DisplayName: "风速",
		Unit:        "m/s",
		Value:       float64(rawVal) / 10,
	}
}
func part11(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[20]<<8 + intData[21]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["WindDirection"] = &model.UplinkMsgPayloadField{
		DisplayName: "风向",
		Unit:        "",
		Value:       float64(rawVal) / 10,
	}
}
func part12(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[22]<<8 + intData[23]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["CO"] = &model.UplinkMsgPayloadField{
		DisplayName: "一氧化碳含量",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part13(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[24]<<8 + intData[25]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SolarRadiation"] = &model.UplinkMsgPayloadField{
		DisplayName: "太阳辐射",
		Unit:        "",
		Value:       float64(rawVal) / 10,
	}
}
func part14(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[26]<<8 + intData[27]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SO2"] = &model.UplinkMsgPayloadField{
		DisplayName: "二氧化硫含量",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part15(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[28]<<8 + intData[29]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["O3"] = &model.UplinkMsgPayloadField{
		DisplayName: "臭氧含量",
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part16(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[30]<<8 + intData[31]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["CH2O"] = &model.UplinkMsgPayloadField{
		DisplayName: "甲醛含量", // CH2O
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part17(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[32]<<8 + intData[33]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["H2"] = &model.UplinkMsgPayloadField{
		DisplayName: "氢气含量", // H2
		Unit:        "%",
		Value:       float64(rawVal) / 10,
	}
}
func part18(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[34]<<8 + intData[35]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Noise"] = &model.UplinkMsgPayloadField{
		DisplayName: "噪音",
		Unit:        "db",
		Value:       float64(rawVal) / 10,
	}
}
func part19(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[36]<<8 + intData[37]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SoilSalinity"] = &model.UplinkMsgPayloadField{
		DisplayName: "土壤盐度",
		Unit:        "%",
		Value:       float64(rawVal) / 100,
	}
}
func part20(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[38]<<8 + intData[39]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SoilPH"] = &model.UplinkMsgPayloadField{
		DisplayName: "土壤PH",
		Unit:        "",
		Value:       float64(rawVal) / 100,
	}
}
func part21(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[40]<<8 + intData[41]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Illumination"] = &model.UplinkMsgPayloadField{
		DisplayName: "光照",
		Unit:        "Lux", // 照度
		Value:       float64(rawVal) * 10,
	}
}
func part22(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[42]<<8 + intData[43]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["SoilConductivity"] = &model.UplinkMsgPayloadField{
		DisplayName: "土壤电导率",
		Unit:        "",
		Value:       float64(rawVal) / 1000,
	}
}
func part23(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[44]<<8 + intData[45]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	attr23 := &model.UplinkMsgPayloadField{
		DisplayName: "雨雪",
		Unit:        "",
		Value:       0,
	}
	if rawVal == 1 {
		attr23.Value = 1
	}
	payload.Attributes["RainOrSnow"] = attr23
}
func part24(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[46]<<8 + intData[47]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["CO2"] = &model.UplinkMsgPayloadField{
		DisplayName: "二氧化碳含量",
		Unit:        "",
		Value:       float64(rawVal),
	}
}
func part25(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[48]<<8 + intData[49]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["PM025"] = &model.UplinkMsgPayloadField{
		DisplayName: "PM2.5",
		Unit:        "微克每立方米",
		Value:       float64(rawVal),
	}
}
func part26(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[50]<<8 + intData[51]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["PM10"] = &model.UplinkMsgPayloadField{
		DisplayName: "PM10",
		Unit:        "%",
		Value:       float64(rawVal),
	}
}
func part27(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[52]<<8 + intData[53]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Dust1_0"] = &model.UplinkMsgPayloadField{
		DisplayName: "粉尘1.0",
		Unit:        "%",
		Value:       float64(rawVal),
	}
}
func part28(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[54]<<8 + intData[55]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Dust2_5"] = &model.UplinkMsgPayloadField{
		DisplayName: "粉尘2.5",
		Unit:        "%",
		Value:       float64(rawVal),
	}
}
func part29(intData []int, payload *model.UplinkMsgPayload) {
	rawVal := intData[56]<<8 + intData[57]
	if rawVal == 32767 || rawVal == 65535 {
		return
	}
	payload.Attributes["Dust10"] = &model.UplinkMsgPayloadField{
		DisplayName: "粉尘10",
		Unit:        "%",
		Value:       float64(rawVal),
	}
}
