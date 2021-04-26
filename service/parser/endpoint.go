package parser

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zsy-cn/bms/protos"
)

// MakeDecodeEndpoint ...
func MakeDecodeEndpoint(srv Parser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.ParserHubUplinkMsg)
		return &protos.Empty{}, srv.Decode(req)
	}
}

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(srv Parser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*protos.Empty)
		return srv.HealthCheck(req)
	}
}
