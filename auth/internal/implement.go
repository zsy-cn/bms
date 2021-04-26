package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/auth"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/util/log"
)

var offlineTimeout = "-48h" // 2天

type DefaultAuthService struct {
	l  log.Logger
	db *gorm.DB
}

func NewAuthService(
	l log.Logger,
	db *gorm.DB,
	cfg auth.AuthConfig,
) (service auth.AuthService, err error) {
	if db == nil {
		err = auth.ErrDbNotBeNil
		return
	}
	service = &DefaultAuthService{
		l:  l,
		db: db,
	}

	return service, nil
}

var _ auth.AuthService = (*DefaultAuthService)(nil)

func (as *DefaultAuthService) Login(params *auth.LoginParams) (customerInfo *auth.CustomerInfo, err error) {
	customerRecord := &model.Customer{}
	err = as.db.Where(&model.Customer{Name: params.Username, Passwd1: params.Password}).First(customerRecord).Error
	if err != nil {
		as.l.Errorf("query customer: %s record failed in Login(): %s", params.Username, err.Error())
		return
	}

	customerInfo = &auth.CustomerInfo{
		ID:   customerRecord.ID,
		Name: customerRecord.Name,
	}
	return
}

// ReAuth 二次认证
func (as *DefaultAuthService) ReAuth(params *auth.ReAuthParams) (err error) {
	customerRecord := &model.Customer{}
	err = as.db.Where(&model.Customer{Base: model.Base{ID: params.ID}, Passwd2: params.Password}).First(customerRecord).Error
	if err != nil {
		as.l.Errorf("query customer: %d record failed in Login(): %s", params.ID, err.Error())
		return
	}

	return
}
