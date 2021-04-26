package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/manager/endpoint"
	"github.com/zsy-cn/bms/service/user/manager/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// ManagerServiceServer ...
type ManagerServiceServer struct {
	LoginHandler transport_grpc.Handler
}

// Login ...
func (server *ManagerServiceServer) Login(ctx context.Context, req *protos.ManagerLoginRequest) (*protos.ManagerLoginResponse, error) {
	_, res, err := server.LoginHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.ManagerLoginResponse), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.Manager) (server *ManagerServiceServer) {
	loginHandler := transport_grpc.NewServer(
		endpoint.MakeLoginEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &ManagerServiceServer{
		LoginHandler: loginHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.Manager) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	managerServiceServer := NewGrpcServer(srv)
	protos.RegisterManagerServiceServer(gprcServer, managerServiceServer)
	gprcServer.Serve(lis)
}
