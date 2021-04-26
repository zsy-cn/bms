package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/brocaar/lora-app-server/api"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("loraclient")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("grpc-addr", conf.LoraclientServicePort)
	viper.SetDefault("appserver-cert", "config/http.pem")
	viper.SetDefault("appserver-addr", conf.LoraAppServerAddr+conf.LoraAppServerPort)
	viper.SetDefault("appserver-token", ``)
}

// ILoraclient Loraclient服务接口
type ILoraclient interface {
	AddAndInitOrg(customer *protos.LoraclientCustormer) (resp *protos.LoraclientAddCustormerResponse, err error)
	UpdateOrgAndUser(customerPb *protos.LoraclientUpdateCustormerRequest) (err error)
	DeleteOrgAndUser(req *protos.LoraclientDeleteCustomerRequest) (err error)

	AddSensor(sensor *protos.LoraclientSensor) (err error)
	UpdateSensor(sensor *protos.LoraclientSensor) (err error)
	DeleteSensor(sensor *protos.LoraclientSensor) (err error)
}

// Loraclient ...
type Loraclient struct {
	loraAppConn *grpc.ClientConn

	userCli                api.UserServiceClient
	orgCli                 api.OrganizationServiceClient
	appCli                 api.ApplicationServiceClient
	deviceProfileCli       api.DeviceProfileServiceClient
	deviceCli              api.DeviceServiceClient
	gatewayCli             api.GatewayServiceClient
	networkServerCli       api.NetworkServerServiceClient
	serviceProfileCli      api.ServiceProfileServiceClient
	defaultNetworkServerID int64

	logger *log.Logger
}

// New ...
func New(logger *log.Logger) (lc ILoraclient, err error) {
	lcServ := &Loraclient{
		logger: logger,
	}
	connectLoraAppServer(lcServ)
	getDefaultNetworkServer(lcServ)
	return lcServ, nil
}

func connectLoraAppServer(lc *Loraclient) {
	opts := getGRPCOpts()
	conn, err := grpc.Dial(viper.GetString("appserver-addr"), opts...)
	if err != nil {
		panic(err)
	}
	lc.loraAppConn = conn
	lc.userCli = api.NewUserServiceClient(conn)
	lc.orgCli = api.NewOrganizationServiceClient(conn)
	lc.deviceCli = api.NewDeviceServiceClient(conn)
	lc.deviceProfileCli = api.NewDeviceProfileServiceClient(conn)
	lc.appCli = api.NewApplicationServiceClient(conn)
	lc.gatewayCli = api.NewGatewayServiceClient(conn)
	lc.serviceProfileCli = api.NewServiceProfileServiceClient(conn)
	lc.networkServerCli = api.NewNetworkServerServiceClient(conn)
}

// getDefaultNetworkServer 获取默认的network server, 如果不存在, 则创建.
func getDefaultNetworkServer(lc *Loraclient) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	listReq := &api.ListNetworkServerRequest{
		Limit: 1,
	}
	listResp, err := lc.networkServerCli.List(ctx, listReq)
	if err != nil {
		lc.logger.Errorf("get network server list failed: %s", err.Error())
	}

	lc.logger.Debugf("network server list: %+v", listResp.Result)

	if listResp.TotalCount == 1 {
		nsID := listResp.Result[0].Id
		lc.logger.Infof("network server already exist: %d", nsID)
		lc.defaultNetworkServerID = nsID
		return
	}

	lc.logger.Infof("network server doesn't exist, try to create it")

	// 如果network server为空, 则创建默认的network server
	networkServer := &api.NetworkServer{
		Name:   "DefaultNetworkServer",
		Server: "loraserver-serv:8000",
	}
	createReq := &api.CreateNetworkServerRequest{
		NetworkServer: networkServer,
	}
	createResp, err := lc.networkServerCli.Create(ctx, createReq)
	if err != nil {
		return err
	}
	if createResp.Id == 0 {
		err = errors.New("创建network server失败")
		return err
	}
	lc.defaultNetworkServerID = createResp.Id
	return
}

// AddAndInitOrg 新增组织, 同时新增其用户, 初始化device profile与app
func (lc *Loraclient) AddAndInitOrg(customer *protos.LoraclientCustormer) (resp *protos.LoraclientAddCustormerResponse, err error) {
	// 由于通过loraclient创建新客户需要同时创建org, user, deviceprofile, serviceprofile, 还初始化各种app, 所以耗时较长
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	orgID, err := lc.initOrg(ctx, customer)
	if err != nil {
		return
	}
	userID, err := lc.initUser(ctx, orgID, customer)
	if err != nil {
		return
	}
	/*
		err = lc.initOrgUser(ctx, orgID, customer.Name)
		if err != nil {
			return
		}
	*/
	spID, err := lc.initServiceProfile(ctx, orgID, customer.OrgName)
	if err != nil {
		return
	}
	err = lc.initDeviceProfile(ctx, orgID, customer.OrgName)
	if err != nil {
		return
	}
	err = lc.initApps(ctx, orgID, spID)
	if err != nil {
		return
	}
	resp = &protos.LoraclientAddCustormerResponse{
		OrgID:  orgID,
		UserID: userID,
	}
	lc.logger.Debug("AddAndInitOrg() completed and returning")
	return
}

// UpdateOrgAndUser ...
func (lc *Loraclient) UpdateOrgAndUser(req *protos.LoraclientUpdateCustormerRequest) (err error) {
	lc.logger.Debug("update org and user in UpdateOrgAndUser() function")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = lc.updateOrg(ctx, req.OrgID, req)
	if err != nil {
		// 错误日志在deleteOrgByID()函数内打印过了
		return
	}

	err = lc.updateUser(ctx, req.UserID, req)
	if err != nil {
		// 错误日志在deleteOrgByID()函数内打印过了
		return
	}
	return
}

// DeleteOrgAndUser 清除org及其user
// 当lora-app-server中不存在目标org和user时, 不返回错误, 只在日志中打印一下此信息即可.
func (lc *Loraclient) DeleteOrgAndUser(req *protos.LoraclientDeleteCustomerRequest) (err error) {
	lc.logger.Debug("delete org and user in DeleteOrgAndUser() function")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = lc.deleteOrgByID(ctx, req.OrgID)
	if err != nil {
		// 错误日志在deleteOrgByID()函数内打印过了
		return
	}

	err = lc.deleteUser(ctx, req.UserID)
	if err != nil {
		// 错误日志在deleteOrgByID()函数内打印过了
		return
	}
	return
}
