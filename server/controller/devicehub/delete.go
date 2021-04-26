package devicehub

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// DeleteParams ...
type DeleteParams struct {
	ID uint64 `param:"<in:path> <required>"`
}

// Serve ...
func (p *DeleteParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request device delete controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	req := &protos.DeleteDeviceRequest{
		ID: p.ID,
	}

	cli, exist := controller.ServiceMap.Get("device_cli")
	if !exist {
		logger.Warn("Trying to delete group failed, grpc was not connected")
		result.Code = -1
		result.Msg = ErrDeviceHubServiceDisconnected.Error()
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.DeviceServiceClient).Delete(_ctx, req)
	if err != nil {
		logger.Errorf("delete device: %d by device service failed: %s", p.ID, err.Error())
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	return
}
