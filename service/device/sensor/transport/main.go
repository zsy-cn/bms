package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/device/sensor/endpoint"
	"github.com/zsy-cn/bms/service/device/sensor/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// DeviceSensorServiceServer ...
type DeviceSensorServiceServer struct {
	AddHandler    transport_grpc.Handler
	UpdateHandler transport_grpc.Handler
	DeleteHandler transport_grpc.Handler
}

// Add ...
func (server *DeviceSensorServiceServer) Add(ctx context.Context, req *protos.DeviceSensor) (*protos.Empty, error) {
	_, res, err := server.AddHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Update ...
func (server *DeviceSensorServiceServer) Update(ctx context.Context, req *protos.DeviceSensor) (*protos.Empty, error) {
	_, res, err := server.UpdateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Delete ...
func (server *DeviceSensorServiceServer) Delete(ctx context.Context, req *protos.DeviceSensor) (*protos.Empty, error) {
	_, res, err := server.DeleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.DeviceSensor) (server *DeviceSensorServiceServer) {
	addHandler := transport_grpc.NewServer(
		endpoint.MakeAddEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &DeviceSensorServiceServer{
		AddHandler:    addHandler,
		UpdateHandler: updateHandler,
		DeleteHandler: deleteHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.DeviceSensor) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	customerServiceServer := NewGrpcServer(srv)
	protos.RegisterDeviceSensorServiceServer(gprcServer, customerServiceServer)
	gprcServer.Serve(lis)
}
