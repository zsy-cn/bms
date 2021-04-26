package devicetype

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"
	"github.com/mitchellh/mapstructure"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// CreateParams ...
type CreateParams struct {
	// <in:body>只能有一个标签, faygo会将请求体格式化为map类型
	Body map[string]interface{} `param:"<in:body> <required>"`
}

// Serve ...
func (rp *CreateParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request manufacturer create controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	deviceTypePb := &protos.DeviceType{}
	body := rp.Body

	err = mapstructure.Decode(body, deviceTypePb)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}

	cli, exist := controller.ServiceMap.Get("core_cli")
	if !exist {
		logger.Info("Trying to create manufacture failed, core service grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.CoreServiceClient).AddDeviceType(_ctx, deviceTypePb)

	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	return
}
