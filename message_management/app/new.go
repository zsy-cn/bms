package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/message_management"
	"github.com/zsy-cn/bms/message_management/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewMessageManagementService(
	l log.Logger,
	db *gorm.DB,
	cfg message_management.MessageManagementConfig,
) (message_management.MessageManagementService, error) {
	return internal.NewMessageManagementService(l, db, cfg)
}
