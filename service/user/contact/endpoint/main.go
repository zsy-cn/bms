package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/contact/service"
)

// MakeGetEndpoint ...
func MakeGetEndpoint(srv *service.Contact) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetContactRequest)
		return srv.Get(req)
	}
}

// MakeGetListEndpoint ...
func MakeGetListEndpoint(srv *service.Contact) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetContactsRequest)
		return srv.GetList(req)
	}
}

// MakeAddEndpoint ...
func MakeAddEndpoint(srv *service.Contact) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Contact)
		return &protos.Empty{}, srv.Add(req)
	}
}

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(srv *service.Contact) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Contact)
		return &protos.Empty{}, srv.Update(req)
	}
}

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(srv *service.Contact) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Contact)
		return &protos.Empty{}, srv.Delete(req)
	}
}
