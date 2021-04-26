package sos

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&Sos{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

// Sos 报警点
type Sos struct {
	model.Base
	model.App
	Status int8 `json:"status"` // [0: 掉线, 1: 正常, 2: 报警]
}
