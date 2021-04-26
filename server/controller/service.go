// Package controller 自动发现可连接的服务, 并添加到map对象中
package controller

import (
	"sync"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
)

// ServiceMapObject ...
// 未使用读写锁, 读操作直接读.
// TheMap 服务映射字典, 键为服务名, 值为grpc连接
type ServiceMapObject struct {
	Lock   sync.Mutex
	TheMap map[string]interface{}
}

// Get ...
func (sm *ServiceMapObject) Get(key string) (val interface{}, exist bool) {
	val, exist = sm.TheMap[key]
	return
}

// Set ...
func (sm *ServiceMapObject) Set(key string, val interface{}) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()

	sm.TheMap[key] = val
}

// ServiceMap ...
var ServiceMap = &ServiceMapObject{
	TheMap: make(map[string]interface{}),
}

func connectCustomerServ() {
	address := viper.GetString("customer-addr")
	customerConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect customer service failed: %s", err.Error())
		return
	}
	customerCli := protos.NewCustomerServiceClient(customerConn)
	ServiceMap.Set("customer_conn", customerConn)
	ServiceMap.Set("customer_cli", customerCli)

	logger.Info("connect customer service successfully")
}

func connectCoreServ() {
	address := viper.GetString("core-addr")
	coreConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect core service failed: %s", err.Error())
		return
	}
	coreCli := protos.NewCoreServiceClient(coreConn)
	ServiceMap.Set("core_conn", coreConn)
	ServiceMap.Set("core_cli", coreCli)
	logger.Info("connect core service successfully")
}

func connectDeviceServ() {
	address := viper.GetString("devicehub-addr")
	deviceConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect devicehub service failed: %s", err.Error())
		return
	}
	deviceCli := protos.NewDeviceServiceClient(deviceConn)
	ServiceMap.Set("device_conn", deviceConn)
	ServiceMap.Set("device_cli", deviceCli)

	logger.Info("connect devicehub service successfully")
}

func connectManagerServ() {
	address := viper.GetString("manager-addr")
	managerConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect manager service failed: %s", err.Error())
		return
	}
	managerCli := protos.NewManagerServiceClient(managerConn)
	ServiceMap.Set("manager_conn", managerConn)
	ServiceMap.Set("manager_cli", managerCli)

	logger.Info("connect manager service successfully")
}

func connectParserHubServ() {
	address := viper.GetString("parserhub-addr")
	parserhubConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect parserhub service failed: %s", err.Error())
		return
	}
	parserhubCli := protos.NewParserHubServiceClient(parserhubConn)
	ServiceMap.Set("parserhub_conn", parserhubConn)
	ServiceMap.Set("parserhub_cli", parserhubCli)

	logger.Info("connect parserhub service successfully")
}

// Start ...
func Start() {
	go connectCoreServ()
	go connectCustomerServ()
	go connectDeviceServ()
	go connectManagerServ()
	go connectParserHubServ()
}
