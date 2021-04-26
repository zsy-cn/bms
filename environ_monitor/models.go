package environ_monitor

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&EnvironMonitor{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

// 环境监测点当前状态表
type EnvironMonitor struct {
	model.Base
	model.App
	Temperature float64 `json:"temperature"`
	PM025       float64 `json:"pm2.5"`
	Noise       float64 `json:"noise"`
	AirSpeed    float64 `json:"air_speed"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
}
