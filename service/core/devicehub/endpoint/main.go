package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/devicehub/service"
)

// MakeGetListEndpoint ...
func MakeGetListEndpoint(srv *service.Device) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetDevicesRequest)
		return srv.GetList(req)
	}
}

// MakeAddEndpoint ...
func MakeAddEndpoint(srv *service.Device) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Device)
		return &protos.Empty{}, srv.Add(req)
	}
}

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(srv *service.Device) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Device)
		return &protos.Empty{}, srv.Update(req)
	}
}

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(srv *service.Device) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteDeviceRequest)
		return &protos.Empty{}, srv.Delete(req)
	}
}

// MakeDeleteGroupEndpoint ...
func MakeDeleteGroupEndpoint(srv *service.Device) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteGroupDeviceRequest)
		return &protos.Empty{}, srv.DeleteGroup(req)
	}
}
