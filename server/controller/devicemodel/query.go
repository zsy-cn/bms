package devicemodel

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// QueryParams 此接口不接受ID字段
type QueryParams struct {
	Name string `param:"<in:query> <name:name>"`
}

// Serve ...
func (p *QueryParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request device model query controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	cli, exist := controller.ServiceMap.Get("core_cli")
	if !exist {
		logger.Info("Trying to list device model failed, grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 这里要加session中的user作为条件
	req := &protos.GetDeviceModelsRequest{}
	deviceModelList, err := cli.(protos.CoreServiceClient).GetDeviceModelList(_ctx, req)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		logger.Errorf("get device model list by core service failed: %s", err.Error())
		return
	}
	result.Data = deviceModelList
	return
}
