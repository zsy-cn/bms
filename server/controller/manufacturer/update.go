package manufacturer

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"
	"github.com/mitchellh/mapstructure"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// UpdateParams ...
type UpdateParams struct {
	ID uint64 `param:"<in:path> <required>"`
	// <in:body>只能有一个标签, faygo会将请求体格式化为map类型
	Body map[string]interface{} `param:"<in:body> <required>"`
}

// Serve ...
func (p *UpdateParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request manufacturer update controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	manufacturerPb := &protos.Manufacturer{}
	body := p.Body

	err = mapstructure.Decode(body, manufacturerPb)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	manufacturerPb.ID = p.ID
	cli, exist := controller.ServiceMap.Get("core_cli")
	if !exist {
		logger.Info("Trying to create manufacturer failed, core service grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.CoreServiceClient).UpdateManufacturer(_ctx, manufacturerPb)

	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	return
}
