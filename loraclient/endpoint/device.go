package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/loraclient/service"
	"github.com/zsy-cn/bms/protos"
)

// MakeAddSensorEndpoint ...
func MakeAddSensorEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientSensor)
		return &protos.Empty{}, srv.AddSensor(req)
	}
}

// MakeUpdateSensorEndpoint ...
func MakeUpdateSensorEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientSensor)
		return &protos.Empty{}, srv.UpdateSensor(req)
	}
}

// MakeDeleteSensorEndpoint ...
func MakeDeleteSensorEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientSensor)
		return &protos.Empty{}, srv.DeleteSensor(req)
	}
}
