package smoke

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&Smoke{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

// 烟感当前状态表
type Smoke struct {
	model.Base
	model.App
	Status int `json:"status"`
}
