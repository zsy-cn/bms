package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/core/parserhub/service"
	"github.com/zsy-cn/bms/service/core/parserhub/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("parserhub")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("grpc-addr", conf.ParserHubServicePort)
}

func main() {
	logger.SetLevel("debug")
	logger.Info("starting parserhub service")
	parserhubServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting parserhub service failed: %s", err.Error())
		return
	}
	logger.Info("starting parserhub grpc server")
	transport.StartGrpcTransport(parserhubServ)
}
