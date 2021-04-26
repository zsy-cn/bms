package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/customer/endpoint"
	"github.com/zsy-cn/bms/service/user/customer/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// CustomerServiceServer ...
type CustomerServiceServer struct {
	GetHandler     transport_grpc.Handler
	GetListHandler transport_grpc.Handler
	AddHandler     transport_grpc.Handler
	UpdateHandler  transport_grpc.Handler
	DeleteHandler  transport_grpc.Handler
}

// Get ...
func (server *CustomerServiceServer) Get(ctx context.Context, req *protos.GetCustomerRequest) (*protos.Customer, error) {
	_, res, err := server.GetHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Customer), nil
}

// GetList ...
func (server *CustomerServiceServer) GetList(ctx context.Context, req *protos.GetCustomersRequest) (*protos.CustomerList, error) {
	_, res, err := server.GetListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.CustomerList), nil
}

// Add ...
func (server *CustomerServiceServer) Add(ctx context.Context, req *protos.Customer) (*protos.Empty, error) {
	_, res, err := server.AddHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Update ...
func (server *CustomerServiceServer) Update(ctx context.Context, req *protos.Customer) (*protos.Empty, error) {
	_, res, err := server.UpdateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Delete ...
func (server *CustomerServiceServer) Delete(ctx context.Context, req *protos.Customer) (*protos.Empty, error) {
	_, res, err := server.DeleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.Customer) (server *CustomerServiceServer) {
	getHandler := transport_grpc.NewServer(
		endpoint.MakeGetEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
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
	return &CustomerServiceServer{
		GetHandler:     getHandler,
		GetListHandler: getListHandler,
		AddHandler:     addHandler,
		UpdateHandler:  updateHandler,
		DeleteHandler:  deleteHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.Customer) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	customerServiceServer := NewGrpcServer(srv)
	protos.RegisterCustomerServiceServer(gprcServer, customerServiceServer)
	gprcServer.Serve(lis)
}
