package customer

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
	logger.Debug("request customer delete controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	customerPb := &protos.Customer{
		ID: p.ID,
	}

	cli, exist := controller.ServiceMap.Get("customer_cli")
	if !exist {
		logger.Info("Trying to delete customer failed, grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.CustomerServiceClient).Delete(_ctx, customerPb)

	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	return
}
