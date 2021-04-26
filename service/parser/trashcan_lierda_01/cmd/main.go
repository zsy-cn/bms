package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/viper"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/service/parser/trashcan_lierda_01"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("parser")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("bind-addr", conf.ParserTrashcanLierda01BindAddr)
	// 注意service地址要能被consul访问到, 且端口与服务监听端口保持一致
	viper.SetDefault("service-addr", conf.ParserTrashcanLierda01ServiceAddr)
	viper.SetDefault("service-port", conf.ParserTrashcanLierda01ServicePort)
	viper.SetDefault("service-name", "trashcan_lierda01")
}

func main() {
	logger.SetLevel("debug")

	var err error
	errChan := make(chan error)
	serviceName := viper.GetString("service-name")
	serviceAddr := viper.GetString("service-addr")
	servicePort := viper.GetString("service-port")

	var parserService parser.Parser
	parserService, err = trashcan_lierda_01.New(logger)
	if err != nil {
		logger.Errorf("starting service %s failed: %s", serviceName, err.Error())
		return
	}

	logger.Infof("start http transport for service %s", serviceName)
	// 启动微服务transport(其实是挂载路由)
	parser.StartHTTPTransport(parserService)

	logger.Debug("service addr: ", serviceAddr)
	registar, err := parser.Register(serviceAddr, servicePort, serviceName)
	if err != nil {
		// 错误信息在parser.Register()函数中打印过了
		return
	}
	go func() {
		// 注册服务到consul
		registar.Register()
		errChan <- http.ListenAndServe(viper.GetString("bind-addr"), nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	// http server运行出错或是ctrl + c取消后errChan中可取到值, 即可往下进行
	err = <-errChan
	registar.Deregister()
	logger.Error(err)
}
