package service

import (
	"context"
	"errors"

	"github.com/zsy-cn/bms/protos"

	"github.com/brocaar/lora-app-server/api"
)

// getOrgIDByDisplayName ...
// 注意: name应该是DisplayName
func (lc *Loraclient) getOrgIDByDisplayName(ctx context.Context, name string) (orgID int64, err error) {
	lc.logger.Debugf("get org id by name: %s DeleteOrgAndUser() function", name)
	req := &api.ListOrganizationRequest{
		Limit:  1,
		Search: name,
	}
	resp, err := lc.orgCli.List(ctx, req)
	if err != nil {
		lc.logger.Errorf("list org %s failed: %s", name, err.Error())
		return
	}
	if len(resp.Result) == 0 {
		lc.logger.Infof("the org: %s doesn't exist", name)
		err = errors.New("record not found")
		return
	}
	orgID = resp.Result[0].Id
	return
}

func (lc *Loraclient) initOrg(ctx context.Context, customer *protos.LoraclientCustormer) (orgID int64, err error) {
	// 创建组织
	org := &api.Organization{
		Name:            customer.OrgName,
		DisplayName:     customer.OrgDisplayName,
		CanHaveGateways: true, // 默认可拥有名下网关
	}
	req := &api.CreateOrganizationRequest{Organization: org}
	lc.logger.Debugf("creating org: %s - %s", customer.OrgName, customer.OrgDisplayName)
	createOrgResp, err := lc.orgCli.Create(ctx, req)
	if err != nil {
		lc.logger.Errorf("create org %s error: %s", customer.OrgName, err.Error())
		return
	}
	orgID = createOrgResp.Id
	return
}

// updateOrg ...
func (lc *Loraclient) updateOrg(ctx context.Context, id int64, customer *protos.LoraclientUpdateCustormerRequest) (err error) {
	lc.logger.Debugf("try to update org: %d in updateOrgByID()", id)
	org := &api.Organization{
		Id:              id,
		Name:            customer.OrgName,
		DisplayName:     customer.OrgDisplayName,
		CanHaveGateways: true, // 默认可拥有名下网关
	}
	req := &api.UpdateOrganizationRequest{
		Organization: org,
	}
	_, err = lc.orgCli.Update(ctx, req)
	if err != nil {
		lc.logger.Errorf("update org: %d failed: %s", id, err.Error())
		return
	}
	return
}

// deleteOrgByID ...
func (lc *Loraclient) deleteOrgByID(ctx context.Context, id int64) (err error) {
	lc.logger.Debugf("try to delete org: %d in deleteOrgByID()", id)
	req := &api.DeleteOrganizationRequest{
		Id: id,
	}
	_, err = lc.orgCli.Delete(ctx, req)
	if err != nil {
		lc.logger.Errorf("delete org: %d failed: %s", id, err.Error())
	}
	return
}

/*
// initOrgUser ...(弃用)
// 先创建org, 再创建目标org的用户
func (lc *Loraclient) initOrgUser(ctx context.Context, orgID int64, name string) (err error) {
	// 创建用户
	orgUser := &api.OrganizationUser{
		OrganizationId: orgID,
		Username:       name,
		IsAdmin:        true,
	}
	req := &api.AddOrganizationUserRequest{
		OrganizationUser: orgUser,
	}
	lc.logger.Debugf("creating user: %s for org: %d", name, orgID)
	_, err = lc.orgCli.AddUser(ctx, req)
	if err != nil {
		lc.logger.Errorf("create user %s failed: %s", name, err.Error())
	}
	return
}
*/
