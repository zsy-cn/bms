package customer

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// QueryParams 此接口不接受ID字段
type QueryParams struct {
	controller.Pagination
	Enable uint8 `param:"<in:query> <name:enable>"`
}

// Serve ...
func (p *QueryParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request customer query controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	cli, exist := controller.ServiceMap.Get("customer_cli")
	if !exist {
		logger.Info("Trying to list customer failed, grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &protos.GetCustomersRequest{
		Pagination: &protos.Pagination{
			Page:     p.Page,
			PageSize: p.PageSize,
			SortBy:   p.SortBy,
			Order:    p.Order,
		},
	}
	customerList, err := cli.(protos.CustomerServiceClient).GetList(_ctx, req)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	result.Data = customerList
	return
}
