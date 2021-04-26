package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/device/camera/service"
	"github.com/zsy-cn/bms/service/device/camera/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("camera")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("grpc-addr", conf.DeviceSensorServicePort)
}

func main() {
	logger.SetLevel("debug")
	logger.Info("starting camera service")
	var cameraServ service.ICamera
	cameraServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting camera service failed: %s", err.Error())
	}
	logger.Info("starting camera grpc server")
	transport.StartGrpcTransport(cameraServ)
}
