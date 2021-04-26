package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/smoke"
	"github.com/zsy-cn/bms/smoke/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewSmokeService(
	l log.Logger,
	db *gorm.DB,
	cfg smoke.SmokeConfig,
) (smoke.SmokeService, error) {
	return internal.NewSmokeService(l, db, cfg)
}
