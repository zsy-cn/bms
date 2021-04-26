package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/user/customer/service"
	"github.com/zsy-cn/bms/service/user/customer/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("customer")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("loraclient-addr", conf.LoraclientServiceAddr+conf.LoraclientServicePort)
	viper.SetDefault("contact-addr", conf.ContactServiceAddr+conf.ContactServicePort)
	viper.SetDefault("grpc-addr", conf.CustomerServicePort)
}

func main() {
	logger.SetLevel("debug")
	logger.Info("starting customer service")
	customerServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting customer service failed: %s", err.Error())
	}
	logger.Info("starting customer grpc server")
	transport.StartGrpcTransport(customerServ)
}
