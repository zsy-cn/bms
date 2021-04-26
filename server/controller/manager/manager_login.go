package manager

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"
	"github.com/mitchellh/mapstructure"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// LoginParams ...
type LoginParams struct {
	// <in:body>只能有一个标签, faygo会将请求体格式化为map类型
	Body map[string]interface{} `param:"<in:body> <required>"`
}

// Serve ...
func (p *LoginParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request manager login controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	loginReqPb := &protos.ManagerLoginRequest{}
	body := p.Body

	err = mapstructure.Decode(body, loginReqPb)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		logger.Errorf("decode login params failed: %s", err.Error())
		return
	}

	cli, exist := controller.ServiceMap.Get("manager_cli")
	if !exist {
		logger.Info("Trying to login manager failed, grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := cli.(protos.ManagerServiceClient).Login(_ctx, loginReqPb)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		logger.Errorf("decode login params failed: %s", err.Error())
		return
	}

	// 登陆成功, 设置session
	faygo.Debugf("session object: %+v", resp)
	ctx.SetSession("Manager", resp)
	return
}
