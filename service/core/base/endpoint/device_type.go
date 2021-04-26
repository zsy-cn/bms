package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/base/service"
)

// MakeGetDeviceTypeListEndpoint ...
func MakeGetDeviceTypeListEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetDeviceTypesRequest)
		return srv.GetDeviceTypeList(req)
	}
}

// MakeAddDeviceTypeEndpoint ...
func MakeAddDeviceTypeEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceType)
		return &protos.Empty{}, srv.AddDeviceType(req)
	}
}

// MakeUpdateDeviceTypeEndpoint ...
func MakeUpdateDeviceTypeEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceType)
		return &protos.Empty{}, srv.UpdateDeviceType(req)
	}
}

// MakeDeleteDeviceTypeEndpoint ...
func MakeDeleteDeviceTypeEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteRequest)
		return &protos.Empty{}, srv.DeleteDeviceType(req)
	}
}
