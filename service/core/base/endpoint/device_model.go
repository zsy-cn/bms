package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/base/service"
)

// MakeGetDeviceModelListEndpoint ...
func MakeGetDeviceModelListEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetDeviceModelsRequest)
		return srv.GetDeviceModelList(req)
	}
}

// MakeAddDeviceModelEndpoint ...
func MakeAddDeviceModelEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceModel)
		return &protos.Empty{}, srv.AddDeviceModel(req)
	}
}

// MakeUpdateDeviceModelEndpoint ...
func MakeUpdateDeviceModelEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceModel)
		return &protos.Empty{}, srv.UpdateDeviceModel(req)
	}
}

// MakeDeleteDeviceModelEndpoint ...
func MakeDeleteDeviceModelEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteRequest)
		return &protos.Empty{}, srv.DeleteDeviceModel(req)
	}
}
