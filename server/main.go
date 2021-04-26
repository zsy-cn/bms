package main

import (
	"os"
	"strings"

	"github.com/henrylee2cn/faygo"
	"github.com/spf13/viper"

	_ "github.com/henrylee2cn/faygo/session/redis"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/server/controller"
	"github.com/zsy-cn/bms/server/controller/customer"
	"github.com/zsy-cn/bms/server/controller/devicehub"
	"github.com/zsy-cn/bms/server/controller/devicemodel"
	"github.com/zsy-cn/bms/server/controller/devicetype"
	"github.com/zsy-cn/bms/server/controller/group"
	"github.com/zsy-cn/bms/server/controller/manager"
	"github.com/zsy-cn/bms/server/controller/manufacturer"
	"github.com/zsy-cn/bms/server/controller/parserhub"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("webserver")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("customer-addr", conf.CustomerServiceAddr+conf.CustomerServicePort)
	viper.SetDefault("core-addr", conf.CoreServiceAddr+conf.CoreServicePort)
	viper.SetDefault("devicehub-addr", conf.DeviceHubServiceAddr+conf.DeviceHubServicePort)
	viper.SetDefault("manager-addr", conf.ManagerServiceAddr+conf.ManagerServicePort)
	viper.SetDefault("parserhub-addr", conf.ParserHubServiceAddr+conf.ParserHubServicePort)
}

func main() {
	// 建立grpc连接, 创建服务映射
	go controller.Start()

	mainApp := faygo.New("mainApp", "0.1")
	// mainApp.GET("/", &controller.EmptyHandler{})

	mainApp.POST("/login", &controller.EmptyHandler{})
	mainApp.POST("/logout", &controller.EmptyHandler{})
	mainApp.POST("/manager/login", &manager.LoginParams{})
	mainApp.POST("/manager/logout", &manager.LogoutParams{})

	// lora-app-server上行信息接口
	mainApp.POST("/uplink", faygo.HandlerFunc(parserhub.UplinkHandler))

	// 核心接口
	mainApp.Route(
		mainApp.NewGroup("core",
			mainApp.NewGroup("manufacturer",
				mainApp.NewNamedGET("获取厂商列表", "", &manufacturer.QueryParams{}),
				// mainApp.NewNamedGET("获取单个厂商", "/:id", &manufacturer.EmptyHandler{}),
				mainApp.NewNamedPOST("创建客户对象", "", &manufacturer.CreateParams{}),
				mainApp.NewNamedPUT("更新厂商", "/:id", &manufacturer.UpdateParams{}),
				mainApp.NewNamedDELETE("删除单个厂商", "/:id", &manufacturer.DeleteParams{}),
			),
			mainApp.NewGroup("devicetype",
				mainApp.NewNamedGET("获取设备类型列表", "", &devicetype.QueryParams{}),
				// mainApp.NewNamedGET("获取单个设备类型", "/:id", &devicetype.EmptyHandler{}),
				mainApp.NewNamedPOST("创建新类型", "", &devicetype.CreateParams{}),
				mainApp.NewNamedPUT("更新设备类型", "/:id", &devicetype.UpdateParams{}),
				mainApp.NewNamedDELETE("删除单个类型", "/:id", &devicetype.DeleteParams{}),
			),
			mainApp.NewGroup("devicemodel",
				mainApp.NewNamedGET("获取设备型号列表", "", &devicemodel.QueryParams{}),
				// mainApp.NewNamedGET("获取单个设备型号", "/:id", &devicemodel.EmptyHandler{}),
				mainApp.NewNamedPOST("创建新型号", "", &devicemodel.CreateParams{}),
				mainApp.NewNamedPUT("更新设备型号", "/:id", &devicemodel.UpdateParams{}),
				mainApp.NewNamedDELETE("删除单个型号对象", "/:id", &devicemodel.DeleteParams{}),
			),
			mainApp.NewNamedGroup("这一接口主要由客户操作, 只能操作当前登录用户的分组信息", "group",
				mainApp.NewNamedGET("获取分组列表", "", &group.QueryParams{}),
				// mainApp.NewNamedGET("获取单个分组", "/:id", &group.EmptyHandler{}),
				mainApp.NewNamedPOST("创建分组对象", "", &group.CreateParams{}),
				mainApp.NewNamedPUT("更新分组", "/:id", &group.UpdateParams{}),
				mainApp.NewNamedDELETE("删除单个分组", "/:id", &group.DeleteParams{}),
			),
		),
	)
	// 客户接口
	mainApp.Route(
		mainApp.NewNamedGroup("客户数据CURD操作", "customer",
			mainApp.NewNamedGET("获取客户信息列表", "", &customer.QueryParams{}),
			// mainApp.NewNamedGET("获取单个客户信息", "/:id", &customer.EmptyHandler{}),
			mainApp.NewNamedPOST("创建客户对象", "", &customer.CreateParams{}),
			mainApp.NewNamedPUT("更新客户信息", "/:id", &customer.UpdateParams{}),
			mainApp.NewNamedDELETE("删除单个客户", "/:id", &customer.DeleteParams{}),
		),
	)

	// 设备接口
	mainApp.Route(
		mainApp.NewNamedGroup("设备CURD操作, 传感器, IP音箱有其通用字段, 也有各自独立的字段", "device",
			mainApp.NewNamedGET("获取设备列表", "", &devicehub.QueryParams{}),
			// mainApp.NewNamedGET("获取单个设备", "/:id", &devicehub.EmptyHandler{}),
			// mainApp.NewNamedPOST("批量导入设备", "/import", &devicehub.EmptyHandler{}),
			mainApp.NewNamedPOST("新增单个设备", "", &devicehub.CreateParams{}),
			mainApp.NewNamedPUT("更新", "/:id", &devicehub.UpdateParams{}),
			mainApp.NewNamedDELETE("删除", "/:id", &devicehub.DeleteParams{}),
		),
	)

	faygo.Run()
}
