package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/loraclient/service"
	"github.com/zsy-cn/bms/protos"
)

// MakeAddOrgEndpoint ...
func MakeAddOrgEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientCustormer)
		return srv.AddAndInitOrg(req)
	}
}

// MakeUpdateOrgEndpoint ...
func MakeUpdateOrgEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientUpdateCustormerRequest)
		return &protos.Empty{}, srv.UpdateOrgAndUser(req)
	}
}

// MakeDeleteOrgEndpoint ...
func MakeDeleteOrgEndpoint(srv service.ILoraclient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.LoraclientDeleteCustomerRequest)
		return &protos.Empty{}, srv.DeleteOrgAndUser(req)
	}
}
