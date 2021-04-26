package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/parserhub/endpoint"
	"github.com/zsy-cn/bms/service/core/parserhub/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// ParserHubServiceServer ...
type ParserHubServiceServer struct {
	ParseHandler transport_grpc.Handler
}

// ParseAndSave ...
func (server *ParserHubServiceServer) ParseAndSave(ctx context.Context, req *protos.ParserHubUplinkMsg) (*protos.Empty, error) {
	_, res, err := server.ParseHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.ParserHub) (server *ParserHubServiceServer) {
	parseHandler := transport_grpc.NewServer(
		endpoint.MakeParseAndSaveEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &ParserHubServiceServer{
		ParseHandler: parseHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.ParserHub) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	parseHubServiceServer := NewGrpcServer(srv)
	protos.RegisterParserHubServiceServer(gprcServer, parseHubServiceServer)
	gprcServer.Serve(lis)
}
