package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/devicehub/endpoint"
	"github.com/zsy-cn/bms/service/core/devicehub/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// DeviceServiceServer ...
type DeviceServiceServer struct {
	GetListHandler     transport_grpc.Handler
	AddHandler         transport_grpc.Handler
	UpdateHandler      transport_grpc.Handler
	DeleteHandler      transport_grpc.Handler
	DeleteGroupHandler transport_grpc.Handler
}

// GetList ...
func (server *DeviceServiceServer) GetList(ctx context.Context, req *protos.GetDevicesRequest) (*protos.DeviceList, error) {
	_, res, err := server.GetListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.DeviceList), nil
}

// Add ...
func (server *DeviceServiceServer) Add(ctx context.Context, req *protos.Device) (*protos.Empty, error) {
	_, res, err := server.AddHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Update ...
func (server *DeviceServiceServer) Update(ctx context.Context, req *protos.Device) (*protos.Empty, error) {
	_, res, err := server.UpdateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Delete ...
func (server *DeviceServiceServer) Delete(ctx context.Context, req *protos.DeleteDeviceRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteGroup ...
func (server *DeviceServiceServer) DeleteGroup(ctx context.Context, req *protos.DeleteGroupDeviceRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteGroupHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.Device) (server *DeviceServiceServer) {
	getListHandler := transport_grpc.NewServer(
		endpoint.MakeGetListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
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
	deleteGroupHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteGroupEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &DeviceServiceServer{
		GetListHandler:     getListHandler,
		AddHandler:         addHandler,
		UpdateHandler:      updateHandler,
		DeleteHandler:      deleteHandler,
		DeleteGroupHandler: deleteGroupHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.Device) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	deviceServiceServer := NewGrpcServer(srv)
	protos.RegisterDeviceServiceServer(gprcServer, deviceServiceServer)
	gprcServer.Serve(lis)
}
