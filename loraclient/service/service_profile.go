package service

import (
	"context"

	"github.com/brocaar/lora-app-server/api"
)

// initServiceProfile ...
func (lc *Loraclient) initServiceProfile(ctx context.Context, orgID int64, name string) (spID string, err error) {
	// 创建service profile, 创建app时需要
	serviceProfile := &api.ServiceProfile{
		Name:            name,
		OrganizationId:  orgID,
		NetworkServerId: lc.defaultNetworkServerID,
	}
	req := &api.CreateServiceProfileRequest{
		ServiceProfile: serviceProfile,
	}
	resp, err := lc.serviceProfileCli.Create(ctx, req)
	lc.logger.Debugf("creating service profile: name: %s, org id: %d", name, orgID)
	lc.logger.Debugf("NetworkServerId: %d", lc.defaultNetworkServerID)
	if err != nil {
		lc.logger.Errorf("create service profile failed: %s", err.Error())
	}
	spID = resp.Id
	return
}
