package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/environ_monitor/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewEnvironMonitorService(
	l log.Logger,
	db *gorm.DB,
	cfg environ_monitor.EnvironMonitorConfig,
) (environ_monitor.EnvironMonitorService, error) {
	return internal.NewEnvironMonitorService(l, db, cfg)
}
