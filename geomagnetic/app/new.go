package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/geomagnetic/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewGeomagneticService(
	l log.Logger,
	db *gorm.DB,
	cfg geomagnetic.GeomagneticConfig,
) (geomagnetic.GeomagneticService, error) {
	return internal.NewGeomagneticService(l, db, cfg)
}
