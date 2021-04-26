package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/user/contact/endpoint"
	"github.com/zsy-cn/bms/service/user/contact/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// ContactServiceServer ...
type ContactServiceServer struct {
	GetHandler     transport_grpc.Handler
	GetListHandler transport_grpc.Handler
	AddHandler     transport_grpc.Handler
	UpdateHandler  transport_grpc.Handler
	DeleteHandler  transport_grpc.Handler
}

// Get ...
func (server *ContactServiceServer) Get(ctx context.Context, req *protos.GetContactRequest) (*protos.Contact, error) {
	_, res, err := server.GetHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Contact), nil
}

// GetList ...
func (server *ContactServiceServer) GetList(ctx context.Context, req *protos.GetContactsRequest) (*protos.ContactList, error) {
	_, res, err := server.GetListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.ContactList), nil
}

// Add ...
func (server *ContactServiceServer) Add(ctx context.Context, req *protos.Contact) (*protos.Empty, error) {
	_, res, err := server.AddHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Update ...
func (server *ContactServiceServer) Update(ctx context.Context, req *protos.Contact) (*protos.Empty, error) {
	_, res, err := server.UpdateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// Delete ...
func (server *ContactServiceServer) Delete(ctx context.Context, req *protos.Contact) (*protos.Empty, error) {
	_, res, err := server.DeleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.Contact) (server *ContactServiceServer) {
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
	return &ContactServiceServer{
		GetHandler:     getHandler,
		GetListHandler: getListHandler,
		AddHandler:     addHandler,
		UpdateHandler:  updateHandler,
		DeleteHandler:  deleteHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.Contact) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	contactServiceServer := NewGrpcServer(srv)
	protos.RegisterContactServiceServer(gprcServer, contactServiceServer)
	gprcServer.Serve(lis)
}
