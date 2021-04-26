package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/device/camera/service"
)

// MakeGetAccessTokenEndpoint ...
func MakeGetAccessTokenEndpoint(srv service.ICamera) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		// 这里不用管输入
		return srv.GetAccessToken()
	}
}

// MakeGetMainScreenEndpoint ...
func MakeGetMainScreenEndpoint(srv service.ICamera) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetMainScreenRequest)
		return srv.GetMainScreen(req)
	}
}

// MakeSetMainScreenEndpoint ...
func MakeSetMainScreenEndpoint(srv service.ICamera) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.SetMainScreenRequest)
		return &protos.Empty{}, srv.SetMainScreen(req)
	}
}
