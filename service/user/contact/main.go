package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/user/contact/service"
	"github.com/zsy-cn/bms/service/user/contact/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("contact")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("grpc-addr", conf.ContactServicePort)
}

func main() {
	logger.SetLevel("debug")
	logger.Info("starting contact service")
	contactServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting contact service failed: %s", err.Error())
	}
	logger.Info("starting contact grpc server")
	transport.StartGrpcTransport(contactServ)
}
