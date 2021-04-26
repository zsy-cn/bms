package service

import (
	"encoding/json"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

func makeDeviceSensorPb(deviceModel *model.Device, extraInfoPb *protos.ExtraDeviceInfo, deviceSensorPb *protos.DeviceSensor) (err error) {
	err = json.Unmarshal([]byte(extraInfoPb.Info), deviceSensorPb)

	deviceSensorPb.Name = deviceModel.Name
	deviceSensorPb.DeviceSN = deviceModel.SerialNumber
	deviceSensorPb.CustomerID = deviceModel.CustomerID
	deviceSensorPb.DeviceTypeID = deviceModel.DeviceTypeID
	return
}
