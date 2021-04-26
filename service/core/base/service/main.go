package service

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// Core 核心服务: 厂商, 类别, 型号, 上行信息解析
type Core struct {
	logger    *log.Logger
	db        *gorm.DB
	deviceCli protos.DeviceServiceClient
}

// New ...
func New(logger *log.Logger) (core *Core, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	address := viper.GetString("devicehub-addr")
	deviceConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect device service failed: %s", err.Error())
		return
	}
	deviceCli := protos.NewDeviceServiceClient(deviceConn)

	core = &Core{
		logger:    logger,
		db:        db,
		deviceCli: deviceCli,
	}
	return
}
