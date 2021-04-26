package devicehub

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

	CustomerID    uint64 `param:"<in:query> <name:customerId>"`
	DeviceTypeID  uint64 `param:"<in:query> <name:deviceTypeId>"`
	GroupID       uint64 `param:"<in:query> <name:groupId>"`
	DeviceModelID uint64 `param:"<in:query> <name:deviceModelId>"`

	Group string `param:"<in:query> <name:group>"`
}

// Serve ...
func (p *QueryParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request device query controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	cli, exist := controller.ServiceMap.Get("device_cli")
	if !exist {
		logger.Warn("Trying to list device failed, grpc was not connected")
		result.Code = -1
		result.Msg = ErrDeviceHubServiceDisconnected.Error()
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &protos.GetDevicesRequest{
		Pagination: &protos.Pagination{
			Page:     p.Page,
			PageSize: p.PageSize,
			SortBy:   p.SortBy,
			Order:    p.Order,
		},
		CustomerID:    p.CustomerID,
		GroupID:       p.GroupID,
		DeviceTypeID:  p.DeviceTypeID,
		DeviceModelID: p.DeviceModelID,

		Group: p.Group,
	}
	deviceList, err := cli.(protos.DeviceServiceClient).GetList(_ctx, req)
	if err != nil {
		logger.Errorf("query devices by device service failed: %s", err.Error())
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	result.Data = deviceList
	return
}
