package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectDB 连接数据库
func ConnectDB() (db *gorm.DB, err error) {
	connectStr := "host=postgres-serv port=5432 user=backend dbname=backend sslmode=disable password=backend"
	db, err = gorm.Open("postgres", connectStr)
	if err != nil {
		logger.Fatalf("Opens database failed: " + err.Error())
		return
	}
	err = db.Exec("set time zone 'Asia/Shanghai';").Error
	if err != nil {
		logger.Fatalf("set time zone failed: " + err.Error())
		return
	}
	logger.Info("connect database success")
	return
}

// DisconnectDB 断开数据库连接
func DisconnectDB(db *gorm.DB) {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}
