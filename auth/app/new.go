package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/auth"
	"github.com/zsy-cn/bms/auth/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewAuthService(
	l log.Logger,
	db *gorm.DB,
	cfg auth.AuthConfig,
) (auth.AuthService, error) {
	return internal.NewAuthService(l, db, cfg)
}
