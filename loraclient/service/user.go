package service

import (
	"context"
	"errors"

	"github.com/brocaar/lora-app-server/api"

	"github.com/zsy-cn/bms/protos"
)

func (lc *Loraclient) initUser(ctx context.Context, orgID int64, customer *protos.LoraclientCustormer) (userID int64, err error) {
	userOrg := &api.UserOrganization{
		OrganizationId: orgID,
		IsAdmin:        true,
	}
	user := &api.User{
		Username: customer.UserName,
		IsAdmin:  false,
		IsActive: true,
		Email:    "000000@qq.com",
	}
	req := &api.CreateUserRequest{
		User:          user,
		Password:      customer.Passwd,
		Organizations: []*api.UserOrganization{userOrg},
	}
	lc.logger.Debugf("creating user: %s", customer.UserName)
	resp, err := lc.userCli.Create(ctx, req)
	if err != nil {
		lc.logger.Errorf("create user %s failed: %s", customer.UserName, err.Error())
	}
	userID = resp.Id
	return
}

func (lc *Loraclient) updateUser(ctx context.Context, userID int64, customer *protos.LoraclientUpdateCustormerRequest) (err error) {
	user := &api.User{
		Id:       userID,
		Username: customer.UserName,
		IsAdmin:  false,
		IsActive: true,
		Email:    "000000@qq.com",
	}
	req := &api.UpdateUserRequest{
		User: user,
	}
	lc.logger.Debugf("updating user in updateUser(): %d", customer.UserID)
	_, err = lc.userCli.Update(ctx, req)
	if err != nil {
		lc.logger.Errorf("update user %d failed in updateUser(): %s", customer.UserID, err.Error())
		return
	}

	updatePasswdReq := &api.UpdateUserPasswordRequest{
		UserId:   userID,
		Password: customer.Passwd,
	}
	lc.logger.Debugf("updating user: %s's password", customer.UserName)
	_, err = lc.userCli.UpdatePassword(ctx, updatePasswdReq)
	if err != nil {
		lc.logger.Errorf("update user %s failed: %s", customer.UserName, err.Error())
		return
	}
	return
}

func (lc *Loraclient) getUserIDByName(ctx context.Context, name string) (userID int64, err error) {
	lc.logger.Debugf("get user id of this name: %s in getUserIDByName()", name)
	req := &api.ListUserRequest{
		Limit:  1,
		Search: name,
	}
	resp, err := lc.userCli.List(ctx, req)
	if err != nil {
		lc.logger.Errorf("list user %s failed: %s", name, err.Error())
		return
	}
	if len(resp.Result) == 0 {
		lc.logger.Infof("the user: %s doesn't exist", name)
		err = errors.New("record not found")
		return
	}
	userID = resp.Result[0].Id
	return
}

func (lc *Loraclient) deleteUser(ctx context.Context, id int64) (err error) {
	lc.logger.Debugf("try to delete user: %d in deleteUser()", id)
	req := &api.DeleteUserRequest{
		Id: id,
	}
	_, err = lc.userCli.Delete(ctx, req)
	if err != nil {
		lc.logger.Errorf("delete user: %d failed: %s", id, err.Error())
		return
	}
	return
}
