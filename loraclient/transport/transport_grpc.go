package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/loraclient/endpoint"
	"github.com/zsy-cn/bms/loraclient/service"
	"github.com/zsy-cn/bms/protos"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// LoraclientServiceServer ...
type LoraclientServiceServer struct {
	AddOrgHandler       transport_grpc.Handler
	UpdateOrgHandler    transport_grpc.Handler
	DeleteOrgHandler    transport_grpc.Handler
	AddSensorHandler    transport_grpc.Handler
	UpdateSensorHandler transport_grpc.Handler
	DeleteSensorHandler transport_grpc.Handler
}

// AddCustmoer ...
func (server *LoraclientServiceServer) AddCustmoer(ctx context.Context, req *protos.LoraclientCustormer) (*protos.LoraclientAddCustormerResponse, error) {
	_, res, err := server.AddOrgHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.LoraclientAddCustormerResponse), nil
}

// UpdateCustomer ...
func (server *LoraclientServiceServer) UpdateCustomer(ctx context.Context, req *protos.LoraclientUpdateCustormerRequest) (*protos.Empty, error) {
	_, res, err := server.UpdateOrgHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteCustmoer ...
func (server *LoraclientServiceServer) DeleteCustmoer(ctx context.Context, req *protos.LoraclientDeleteCustomerRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteOrgHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// AddSensor ...
func (server *LoraclientServiceServer) AddSensor(ctx context.Context, req *protos.LoraclientSensor) (*protos.Empty, error) {
	_, res, err := server.AddSensorHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// UpdateSensor ...
func (server *LoraclientServiceServer) UpdateSensor(ctx context.Context, req *protos.LoraclientSensor) (*protos.Empty, error) {
	_, res, err := server.UpdateSensorHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteSensor ...
func (server *LoraclientServiceServer) DeleteSensor(ctx context.Context, req *protos.LoraclientSensor) (*protos.Empty, error) {
	_, res, err := server.DeleteSensorHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv service.ILoraclient) (server *LoraclientServiceServer) {
	addOrgHandler := transport_grpc.NewServer(
		endpoint.MakeAddOrgEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateOrgHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateOrgEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteOrgHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteOrgEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	addSensorHandler := transport_grpc.NewServer(
		endpoint.MakeAddSensorEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateSensorHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateSensorEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteSensorHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteSensorEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &LoraclientServiceServer{
		AddOrgHandler:       addOrgHandler,
		UpdateOrgHandler:    updateOrgHandler,
		DeleteOrgHandler:    deleteOrgHandler,
		AddSensorHandler:    addSensorHandler,
		UpdateSensorHandler: updateSensorHandler,
		DeleteSensorHandler: deleteSensorHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv service.ILoraclient) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	loraclientServiceServer := NewGrpcServer(srv)
	protos.RegisterLoraclientServiceServer(gprcServer, loraclientServiceServer)
	gprcServer.Serve(lis)
}
