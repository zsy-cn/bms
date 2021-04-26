package service

import (
	"context"
	"time"

	"github.com/brocaar/lora-app-server/api"

	"github.com/zsy-cn/bms/protos"
)

func (lc *Loraclient) makeDevice(ctx context.Context, sensor *protos.LoraclientSensor) (device *api.Device, err error) {
	lc.logger.Debugf("making lora device object in makeDevice()")
	orgID, err := lc.getOrgIDByDisplayName(ctx, sensor.OrgDisplayName)
	if err != nil {
		lc.logger.Errorf("getOrgIDByDisplayName() failed in makeDevice: %s", err.Error())
		return
	}

	appID, err := lc.getAppIDByName(ctx, sensor.Type, orgID)
	if err != nil {
		lc.logger.Errorf("getAppIDByName() failed in makeDevice: %s", err.Error())
		return
	}

	dpID, err := lc.getDeviceProfileID(ctx, appID, orgID)
	if err != nil {
		lc.logger.Errorf("getDeviceProfileID() failed in makeDevice: %s", err.Error())
		return
	}

	device = &api.Device{
		Name:            sensor.Name,
		DevEui:          sensor.DevEUI,
		ApplicationId:   appID,
		DeviceProfileId: dpID,
		Description:     "",
	}
	return
}

// AddSensor 添加Lora设备
func (lc *Loraclient) AddSensor(sensor *protos.LoraclientSensor) (err error) {
	lc.logger.Debugf("add sensor: %s in AddSensor()", sensor.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	device, err := lc.makeDevice(ctx, sensor)
	if err != nil {
		// error日志(获取org及user等相关信息)在makeDevice()中打印
		return
	}
	req := &api.CreateDeviceRequest{
		Device: device,
	}
	_, err = lc.deviceCli.Create(ctx, req)
	if err != nil {
		lc.logger.Errorf("create sensor: %s failed: %s", sensor.Name, err.Error())
		return
	}

	keys := &api.DeviceKeys{
		DevEui: sensor.DevEUI,
		NwkKey: sensor.AppKey,
	}

	createKeysReq := &api.CreateDeviceKeysRequest{
		DeviceKeys: keys,
	}
	_, err = lc.deviceCli.CreateKeys(ctx, createKeysReq)
	if err != nil {
		return
	}

	return
}

// UpdateSensor 更新Lora设备
func (lc *Loraclient) UpdateSensor(sensor *protos.LoraclientSensor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	device, err := lc.makeDevice(ctx, sensor)
	if err != nil {
		// error日志(获取org及user等相关信息)在makeDevice()中打印
		return
	}
	req := &api.UpdateDeviceRequest{
		Device: device,
	}
	_, err = lc.deviceCli.Update(ctx, req)
	if err != nil {
		return
	}

	return
}

// DeleteSensor ...
// 删除设备只需要devEUI, 意味着全局(各组织间)devEUI是唯一的.
func (lc *Loraclient) DeleteSensor(sensor *protos.LoraclientSensor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &api.DeleteDeviceRequest{
		DevEui: sensor.DevEUI,
	}
	_, err = lc.deviceCli.Delete(ctx, req)
	if err != nil {
		return
	}

	return
}
