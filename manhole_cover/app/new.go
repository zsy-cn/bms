package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/manhole_cover/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewManholeCoverService(
	l log.Logger,
	db *gorm.DB,
	cfg manhole_cover.ManholeCoverConfig,
) (manhole_cover.ManholeCoverService, error) {
	return internal.NewManholeCoverService(l, db, cfg)
}
