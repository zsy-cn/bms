package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/parserhub/service"
)

// MakeParseAndSaveEndpoint ...
func MakeParseAndSaveEndpoint(srv *service.ParserHub) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.ParserHubUplinkMsg)
		return &protos.Empty{}, srv.ParseAndSave(req)
	}
}
