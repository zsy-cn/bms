package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/manager/service"
)

// MakeLoginEndpoint ...
func MakeLoginEndpoint(srv *service.Manager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.ManagerLoginRequest)
		return srv.Login(req)
	}
}
