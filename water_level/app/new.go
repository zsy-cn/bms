package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/water_level"
	"github.com/zsy-cn/bms/water_level/internal"
)

func NewWaterLevelService(
	l log.Logger,
	db *gorm.DB,
	cfg water_level.WaterLevelConfig,
) (water_level.WaterLevelService, error) {
	return internal.NewWaterLevelService(l, db, cfg)
}
