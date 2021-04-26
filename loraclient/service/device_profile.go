package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/brocaar/lora-app-server/api"
)

// getDeviceProfileID ...
func (lc *Loraclient) getDeviceProfileID(ctx context.Context, appID int64, orgID int64) (dpID string, err error) {
	req := &api.ListDeviceProfileRequest{
		Limit:          1,
		OrganizationId: orgID,
		ApplicationId:  appID,
	}
	resp, err := lc.deviceProfileCli.List(ctx, req)
	if err != nil {
		lc.logger.Errorf("Find device profile for app %d error: %s", appID, err.Error())
		return
	}
	if len(resp.Result) == 0 {
		errStr := fmt.Sprintf("Didn't find the app: %d's device profile", appID)
		lc.logger.Error(errStr)
		err = errors.New(errStr)
		return
	}
	dpID = resp.Result[0].Id
	return
}

// initDeviceProfile ...
func (lc *Loraclient) initDeviceProfile(ctx context.Context, orgID int64, name string) (err error) {
	// 创建device profile, 不同组织间不能共用
	deviceProfile := &api.DeviceProfile{
		Name:            name,
		OrganizationId:  orgID,
		MacVersion:      "1.0.2",
		SupportsJoin:    true, // 支持OTAA
		NetworkServerId: lc.defaultNetworkServerID,
	}
	req := &api.CreateDeviceProfileRequest{
		DeviceProfile: deviceProfile,
	}
	_, err = lc.deviceProfileCli.Create(ctx, req)
	lc.logger.Debugf("creating device profile: name: %s, org id: %d", name, orgID)
	if err != nil {
		lc.logger.Errorf("create device profile failed: %s", err.Error())
	}
	return
}
