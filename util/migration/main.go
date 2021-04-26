package migration

import (
	"os"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func Migrate() {
	var err error
	db, err := conf.ConnectDB()
	if err != nil {
		panic(err)
	}
	logger.Info("Database init start")
	InitTables(db)
	SetupBaseDatas(db)
	logger.Info("Database init complete")
}
