package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/sos"
	"github.com/zsy-cn/bms/sos/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewSosService(
	l log.Logger,
	db *gorm.DB,
	cfg sos.SosConfig,
) (sos.SosService, error) {
	return internal.NewSosService(l, db, cfg)
}
