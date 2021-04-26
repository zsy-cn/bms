package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/base/service"
)

// MakeGetGroupListEndpoint ...
func MakeGetGroupListEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.GetGroupsRequest)
		return srv.GetGroupList(req)
	}
}

// MakeAddGroupEndpoint ...
func MakeAddGroupEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Group)
		return &protos.Empty{}, srv.AddGroup(req)
	}
}

// MakeUpdateGroupEndpoint ...
func MakeUpdateGroupEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Group)
		return &protos.Empty{}, srv.UpdateGroup(req)
	}
}

// MakeDeleteGroupEndpoint ...
func MakeDeleteGroupEndpoint(srv *service.Core) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.DeleteGroupRequest)
		return &protos.Empty{}, srv.DeleteGroup(req)
	}
}
