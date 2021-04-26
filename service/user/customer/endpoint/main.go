package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/customer/service"
)

// MakeGetEndpoint ...
func MakeGetEndpoint(srv *service.Customer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetCustomerRequest)
		return srv.Get(req)
	}
}

// MakeGetListEndpoint ...
func MakeGetListEndpoint(srv *service.Customer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetCustomersRequest)
		return srv.GetList(req)
	}
}

// MakeAddEndpoint ...
func MakeAddEndpoint(srv *service.Customer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Customer)
		return &protos.Empty{}, srv.Add(req)
	}
}

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(srv *service.Customer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Customer)
		return &protos.Empty{}, srv.Update(req)
	}
}

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(srv *service.Customer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Customer)
		return &protos.Empty{}, srv.Delete(req)
	}
}
