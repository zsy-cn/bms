package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/trashcan/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewTrashcanService(
	l log.Logger,
	db *gorm.DB,
	cfg trashcan.TrashcanConfig,
) (trashcan.TrashcanService, error) {
	return internal.NewTrashcanService(l, db, cfg)
}
