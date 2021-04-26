package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/wifi"
	"github.com/zsy-cn/bms/wifi/internal"
)

func NewWifiService(
	l log.Logger,
	db *gorm.DB,
	cfg wifi.WifiConfig,
) (wifi.WifiService, error) {
	return internal.NewWifiService(l, db, cfg)
}
