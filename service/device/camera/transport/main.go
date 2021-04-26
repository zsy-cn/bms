package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/device/camera/endpoint"
	"github.com/zsy-cn/bms/service/device/camera/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// DeviceSensorServiceServer ...
type DeviceSensorServiceServer struct {
	GetAccessTokenHandler transport_grpc.Handler
	GetMainScreenHandler  transport_grpc.Handler
	SetMainScreenHandler  transport_grpc.Handler
}

// GetAccessToken ...
func (server *DeviceSensorServiceServer) GetAccessToken(ctx context.Context, req *protos.Empty) (*protos.GetAccessTokenResponse, error) {
	_, res, err := server.GetAccessTokenHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.GetAccessTokenResponse), nil
}

// SetMainScreen ...
func (server *DeviceSensorServiceServer) SetMainScreen(ctx context.Context, req *protos.SetMainScreenRequest) (*protos.Empty, error) {
	_, res, err := server.SetMainScreenHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// GetMainScreen ...
func (server *DeviceSensorServiceServer) GetMainScreen(ctx context.Context, req *protos.GetMainScreenRequest) (*protos.GetMainScreenResponse, error) {
	_, res, err := server.GetMainScreenHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.GetMainScreenResponse), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv service.ICamera) (server *DeviceSensorServiceServer) {
	getAccessTokenHandler := transport_grpc.NewServer(
		endpoint.MakeGetAccessTokenEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &DeviceSensorServiceServer{
		GetAccessTokenHandler: getAccessTokenHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv service.ICamera) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	cameraServiceServer := NewGrpcServer(srv)
	protos.RegisterCameraServiceServer(gprcServer, cameraServiceServer)
	gprcServer.Serve(lis)
}
