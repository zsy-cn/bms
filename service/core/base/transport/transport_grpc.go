package transport

import (
	"context"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/core/base/endpoint"
	"github.com/zsy-cn/bms/service/core/base/service"
)

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// CoreServiceServer ...
type CoreServiceServer struct {
	GetGroupListHandler transport_grpc.Handler
	AddGroupHandler     transport_grpc.Handler
	UpdateGroupHandler  transport_grpc.Handler
	DeleteGroupHandler  transport_grpc.Handler

	GetDeviceTypeListHandler transport_grpc.Handler
	AddDeviceTypeHandler     transport_grpc.Handler
	UpdateDeviceTypeHandler  transport_grpc.Handler
	DeleteDeviceTypeHandler  transport_grpc.Handler

	GetDeviceModelListHandler transport_grpc.Handler
	AddDeviceModelHandler     transport_grpc.Handler
	UpdateDeviceModelHandler  transport_grpc.Handler
	DeleteDeviceModelHandler  transport_grpc.Handler

	GetManufacturerListHandler transport_grpc.Handler
	AddManufacturerHandler     transport_grpc.Handler
	UpdateManufacturerHandler  transport_grpc.Handler
	DeleteManufacturerHandler  transport_grpc.Handler
}

// GetGroupList ...
func (server *CoreServiceServer) GetGroupList(ctx context.Context, req *protos.GetGroupsRequest) (*protos.GroupList, error) {
	_, res, err := server.GetGroupListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.GroupList), nil
}

// AddGroup ...
func (server *CoreServiceServer) AddGroup(ctx context.Context, req *protos.Group) (*protos.Empty, error) {
	_, res, err := server.AddGroupHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// UpdateGroup ...
func (server *CoreServiceServer) UpdateGroup(ctx context.Context, req *protos.Group) (*protos.Empty, error) {
	_, res, err := server.UpdateGroupHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteGroup ...
func (server *CoreServiceServer) DeleteGroup(ctx context.Context, req *protos.DeleteGroupRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteGroupHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// GetDeviceTypeList ...
func (server *CoreServiceServer) GetDeviceTypeList(ctx context.Context, req *protos.GetDeviceTypesRequest) (*protos.DeviceTypeList, error) {
	_, res, err := server.GetDeviceTypeListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.DeviceTypeList), nil
}

// AddDeviceType ...
func (server *CoreServiceServer) AddDeviceType(ctx context.Context, req *protos.DeviceType) (*protos.Empty, error) {
	_, res, err := server.AddDeviceTypeHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// UpdateDeviceType ...
func (server *CoreServiceServer) UpdateDeviceType(ctx context.Context, req *protos.DeviceType) (*protos.Empty, error) {
	_, res, err := server.UpdateDeviceTypeHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteDeviceType ...
func (server *CoreServiceServer) DeleteDeviceType(ctx context.Context, req *protos.DeleteRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteDeviceTypeHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// GetDeviceModelList ...
func (server *CoreServiceServer) GetDeviceModelList(ctx context.Context, req *protos.GetDeviceModelsRequest) (*protos.DeviceModelList, error) {
	_, res, err := server.GetDeviceModelListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.DeviceModelList), nil
}

// AddDeviceModel ...
func (server *CoreServiceServer) AddDeviceModel(ctx context.Context, req *protos.DeviceModel) (*protos.Empty, error) {
	_, res, err := server.AddDeviceModelHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// UpdateDeviceModel ...
func (server *CoreServiceServer) UpdateDeviceModel(ctx context.Context, req *protos.DeviceModel) (*protos.Empty, error) {
	_, res, err := server.UpdateDeviceModelHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteDeviceModel ...
func (server *CoreServiceServer) DeleteDeviceModel(ctx context.Context, req *protos.DeleteRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteDeviceModelHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// GetManufacturerList ...
func (server *CoreServiceServer) GetManufacturerList(ctx context.Context, req *protos.GetManufacturersRequest) (*protos.ManufacturerList, error) {
	_, res, err := server.GetManufacturerListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.ManufacturerList), nil
}

// AddManufacturer ...
func (server *CoreServiceServer) AddManufacturer(ctx context.Context, req *protos.Manufacturer) (*protos.Empty, error) {
	_, res, err := server.AddManufacturerHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// UpdateManufacturer ...
func (server *CoreServiceServer) UpdateManufacturer(ctx context.Context, req *protos.Manufacturer) (*protos.Empty, error) {
	_, res, err := server.UpdateManufacturerHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// DeleteManufacturer ...
func (server *CoreServiceServer) DeleteManufacturer(ctx context.Context, req *protos.DeleteRequest) (*protos.Empty, error) {
	_, res, err := server.DeleteManufacturerHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protos.Empty), nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *service.Core) (server *CoreServiceServer) {
	getGroupListHandler := transport_grpc.NewServer(
		endpoint.MakeGetGroupListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	addGroupHandler := transport_grpc.NewServer(
		endpoint.MakeAddGroupEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateGroupHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateGroupEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteGroupHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteGroupEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)

	getDeviceTypeListHandler := transport_grpc.NewServer(
		endpoint.MakeGetDeviceTypeListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	addDeviceTypeHandler := transport_grpc.NewServer(
		endpoint.MakeAddDeviceTypeEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateDeviceTypeHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateDeviceTypeEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteDeviceTypeHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteDeviceTypeEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)

	getDeviceModelListHandler := transport_grpc.NewServer(
		endpoint.MakeGetDeviceModelListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	addDeviceModelHandler := transport_grpc.NewServer(
		endpoint.MakeAddDeviceModelEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateDeviceModelHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateDeviceModelEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteDeviceModelHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteDeviceModelEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)

	getManufacturerListHandler := transport_grpc.NewServer(
		endpoint.MakeGetManufacturerListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	addManufacturerHandler := transport_grpc.NewServer(
		endpoint.MakeAddManufacturerEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	updateManufacturerHandler := transport_grpc.NewServer(
		endpoint.MakeUpdateManufacturerEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	deleteManufacturerHandler := transport_grpc.NewServer(
		endpoint.MakeDeleteManufacturerEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)

	return &CoreServiceServer{
		GetGroupListHandler: getGroupListHandler,
		AddGroupHandler:     addGroupHandler,
		UpdateGroupHandler:  updateGroupHandler,
		DeleteGroupHandler:  deleteGroupHandler,

		GetDeviceTypeListHandler: getDeviceTypeListHandler,
		AddDeviceTypeHandler:     addDeviceTypeHandler,
		UpdateDeviceTypeHandler:  updateDeviceTypeHandler,
		DeleteDeviceTypeHandler:  deleteDeviceTypeHandler,

		GetDeviceModelListHandler: getDeviceModelListHandler,
		AddDeviceModelHandler:     addDeviceModelHandler,
		UpdateDeviceModelHandler:  updateDeviceModelHandler,
		DeleteDeviceModelHandler:  deleteDeviceModelHandler,

		GetManufacturerListHandler: getManufacturerListHandler,
		AddManufacturerHandler:     addManufacturerHandler,
		UpdateManufacturerHandler:  updateManufacturerHandler,
		DeleteManufacturerHandler:  deleteManufacturerHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *service.Core) {
	lis, _ := net.Listen("tcp", viper.GetString("grpc-addr"))
	gprcServer := grpc.NewServer()
	coreServiceServer := NewGrpcServer(srv)
	protos.RegisterCoreServiceServer(gprcServer, coreServiceServer)
	gprcServer.Serve(lis)
}
