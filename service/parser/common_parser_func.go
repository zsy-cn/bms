package parser

import (
	"bytes"
	"encoding/binary"

	"github.com/zsy-cn/bms/model"
)

// VoltageDecoder 解析单字节电压值
// rawVoltage 为单字节变量
// displayName默认为"电压"
func VoltageDecoder(attrbutes model.UplinkMsgAttribute, rawVoltage [1]byte, displayName string) {
	if displayName == "" {
		displayName = "电压"
	}
	_voltageAttr := &model.UplinkMsgPayloadField{
		DisplayName: displayName,
		Value:       0,
		Unit:        "V",
	}
	buf := bytes.NewBuffer(rawVoltage[:])
	// intVoltage的类型为int8, 有符号数, binary.Read()会自动进行正负转换的
	var intVoltage int8
	binary.Read(buf, binary.BigEndian, &intVoltage)
	_voltageAttr.Value = float64(intVoltage) / 10
	attrbutes["Voltage"] = _voltageAttr
}

// TemperatureDecoder 解析双字节温度值
// rawTemperature 为双字节变量
// displayName默认为"温度"
func TemperatureDecoder(attrbutes model.UplinkMsgAttribute, rawTemperature [2]byte, displayName string) {
	if displayName == "" {
		displayName = "温度"
	}
	_tempAttr := &model.UplinkMsgPayloadField{
		DisplayName: displayName,
		Value:       0,
		Unit:        "C",
	}

	buf := bytes.NewBuffer(rawTemperature[:])
	// intVoltage的类型为int8, 有符号数, binary.Read()会自动进行正负转换的
	var intTemperature int16
	binary.Read(buf, binary.BigEndian, &intTemperature)
	_tempAttr.Value = float64(intTemperature) / 100
	attrbutes["Temperature"] = _tempAttr
}
