package internal

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/message_management"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

type DefaultMessageManagementService struct {
	l  log.Logger
	db *gorm.DB
}

func NewMessageManagementService(
	l log.Logger,
	db *gorm.DB,
	cfg message_management.MessageManagementConfig,
) (service message_management.MessageManagementService, err error) {
	if db == nil {
		err = message_management.ErrDbNotBeNil
		return
	}

	service = &DefaultMessageManagementService{
		l:  l,
		db: db,
	}

	return service, nil
}

var _ message_management.MessageManagementService = (*DefaultMessageManagementService)(nil)

func (ss *DefaultMessageManagementService) GetNotifications(req *protos.GetDevicesRequestForCustomer) (notifications *protos.NotificationList, err error) {
	return GetNotifications(ss.db, ss.l, req)
}

func (ss *DefaultMessageManagementService) DiscardNotification(customerID uint64, deviceSN string) (err error) {
	return DiscardNotification(ss.db, ss.l, customerID, deviceSN)
}
