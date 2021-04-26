package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/base/service"
)

// MakeGetManufacturerListEndpoint ...
func MakeGetManufacturerListEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetManufacturersRequest)
		return srv.GetManufacturerList(req)
	}
}

// MakeAddManufacturerEndpoint ...
func MakeAddManufacturerEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Manufacturer)
		return &protos.Empty{}, srv.AddManufacturer(req)
	}
}

// MakeUpdateManufacturerEndpoint ...
func MakeUpdateManufacturerEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Manufacturer)
		return &protos.Empty{}, srv.UpdateManufacturer(req)
	}
}

// MakeDeleteManufacturerEndpoint ...
func MakeDeleteManufacturerEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteRequest)
		return &protos.Empty{}, srv.DeleteManufacturer(req)
	}
}
