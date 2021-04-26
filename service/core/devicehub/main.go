package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/core/devicehub/service"
	"github.com/zsy-cn/bms/service/core/devicehub/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("devicehub")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("device-sensor-addr", conf.DeviceSensorServiceAddr+conf.DeviceSensorServicePort)
	viper.SetDefault("grpc-addr", conf.DeviceHubServicePort)
}

func main() {
	logger.SetLevel("debug")
	logger.Info("starting devicehub service")
	deviceServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting devicehub service failed: %s", err.Error())
	}
	logger.Info("starting devicehub grpc server")
	transport.StartGrpcTransport(deviceServ)
}
