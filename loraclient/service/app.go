package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"

	"github.com/brocaar/lora-app-server/api"
)

// getAppIDByName ...
func (lc *Loraclient) getAppIDByName(ctx context.Context, name string, orgID int64) (appID int64, err error) {
	req := &api.ListApplicationRequest{
		Limit:          1,
		Search:         name,
		OrganizationId: orgID,
	}
	resp, err := lc.appCli.List(ctx, req)
	if err != nil {
		lc.logger.Errorf("Find app %s error: %s", name, err.Error())
		return
	}
	if len(resp.Result) == 0 {
		errStr := fmt.Sprintf("Didn't find the app: %s", name)
		lc.logger.Error(errStr)
		err = errors.New(errStr)
		return
	}
	appID = resp.Result[0].Id
	return
}

// initApps 初始化App列表
// @spID: service profile id
func (lc *Loraclient) initApps(ctx context.Context, orgID int64, spID string) (err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		return
	}
	sensorTypes := []string{}
	// 为每种传感器设备都创建一个app, 方便后期管理
	db.Model(&model.DeviceType{}).Where(&model.DeviceType{IsSensor: true}).Pluck("key", &sensorTypes)
	for _, sensorType := range sensorTypes {
		app := &api.Application{
			Name:             sensorType,
			Description:      "",
			OrganizationId:   orgID,
			ServiceProfileId: spID,
		}
		req := &api.CreateApplicationRequest{
			Application: app,
		}
		_, err := lc.appCli.Create(ctx, req)
		lc.logger.Errorf("creating app %s", sensorType)
		if err != nil {
			lc.logger.Errorf("create app %s error: %s", sensorType, err.Error())
		}
	}
	return
}
