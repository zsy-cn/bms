package water_level

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&WaterLevel{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

// WaterLevel 水位当前状态表
type WaterLevel struct {
	model.Base
	model.App
	Value  int // 水位值
	Status int // [0: 掉线, 1: 正常, 2: 报警], 入库时根据规则判断
}
