package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// DispatchAdd ...
// 按照不同设备类型, 调用不同的设备服务执行添加操作(sensor, sound_box, router等)
func (d *Device) DispatchAdd(deviceModel *model.Device, extraInfoPb *protos.ExtraDeviceInfo) (err error) {
	d.logger.Debugf("dispatch the add operation for device: %s", deviceModel.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	deviceTypeModel := &model.DeviceType{}
	err = d.db.First(deviceTypeModel, deviceModel.DeviceTypeID).Error
	if err != nil {
		d.logger.Errorf("get the device type: %d failed: %s", deviceModel.DeviceTypeID, err.Error())
		return
	}
	d.logger.Debugf("the device(to be added): %s's type is: %s-%s", deviceModel.Name, deviceTypeModel.Key, deviceTypeModel.Name)

	switch deviceTypeModel.Key {
	// sensor设备
	case "environ_monitor", "geomagnetic", "trashcan", "sos", "water_level", "manhole_cover":
		deviceSensorPb := &protos.DeviceSensor{}
		err = makeDeviceSensorPb(deviceModel, extraInfoPb, deviceSensorPb)
		if err != nil {
			d.logger.Errorf("make device sensor: %s protobuf object failed: %s", deviceSensorPb.Name, err.Error())
		}
		_, err = d.deviceSensorCli.Add(ctx, deviceSensorPb)
		break
	case "sound_box":
		break
	default:
		errStr := fmt.Sprintf("Unknown device type: %s", deviceTypeModel.Key)
		err = errors.New(errStr)
	}
	if err != nil {
		d.logger.Errorf("insert device: %s by %s service failed: %s", deviceModel.Name, deviceTypeModel.Name, err.Error())
		return
	}

	return
}
