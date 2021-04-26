package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/device/sensor/service"
)

// MakeAddEndpoint ...
func MakeAddEndpoint(srv *service.DeviceSensor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceSensor)
		return &protos.Empty{}, srv.Add(req)
	}
}

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(srv *service.DeviceSensor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceSensor)
		return &protos.Empty{}, srv.Update(req)
	}
}

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(srv *service.DeviceSensor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeviceSensor)
		return &protos.Empty{}, srv.Delete(req)
	}
}
