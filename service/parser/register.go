package parser

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	gkLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"

	"github.com/zsy-cn/bms/conf"
)

// Register ...
func Register(serviceAddr, servicePort, serviceName string) (registar sd.Registrar, err error) {
	// Logging domain.
	var gkLogger gkLog.Logger
	{
		gkLogger = gkLog.NewLogfmtLogger(os.Stderr)
		gkLogger = gkLog.With(gkLogger, "ts", gkLog.DefaultTimestampUTC)
		gkLogger = gkLog.With(gkLogger, "caller", gkLog.DefaultCaller)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	client, err := conf.ConnectConsul()
	if err != nil {
		logger.Errorf("connect consul failed: %s", err.Error())
		return
	}
	// 健康检查地址
	checkURL := "http://" + serviceAddr + ":" + servicePort + "/health"
	// 没有设置DeregisterCriticalServiceAfter字段, 服务不会因为check未成功而注销
	check := api.AgentServiceCheck{
		HTTP:     checkURL,
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(servicePort)
	// 确保服务ID唯一
	num := rand.Intn(100)
	asr := api.AgentServiceRegistration{
		ID:      serviceName + strconv.Itoa(num),
		Name:    serviceName,
		Address: serviceAddr,
		Port:    port,
		Tags:    []string{serviceName},
		Check:   &check,
	}
	registar = consulsd.NewRegistrar(client, &asr, gkLogger)
	return
}
