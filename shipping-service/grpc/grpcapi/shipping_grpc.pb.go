// Code generated by protoc-gen-go-grpcapi. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpcapi v1.5.1
// - protoc             v5.28.2
// source: shipping.proto

package grpcapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcapi package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShippingService_CalculatePrice_FullMethodName       = "/truckgo.ShippingService/CalculatePrice"
	ShippingService_CalculateRoute_FullMethodName       = "/truckgo.ShippingService/CalculateRoute"
	ShippingService_FindTheNearestDriver_FullMethodName = "/truckgo.ShippingService/FindTheNearestDriver"
	ShippingService_TestFunc_FullMethodName             = "/truckgo.ShippingService/TestFunc"
)

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CalculatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error)
	CalculateRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	FindTheNearestDriver(ctx context.Context, in *DriverRequest, opts ...grpc.CallOption) (*DriverResponse, error)
	TestFunc(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CalculatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, ShippingService_CalculatePrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) CalculateRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, ShippingService_CalculateRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) FindTheNearestDriver(ctx context.Context, in *DriverRequest, opts ...grpc.CallOption) (*DriverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DriverResponse)
	err := c.cc.Invoke(ctx, ShippingService_FindTheNearestDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) TestFunc(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, ShippingService_TestFunc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
// All implementations must embed UnimplementedShippingServiceServer
// for forward compatibility.
type ShippingServiceServer interface {
	CalculatePrice(context.Context, *PriceRequest) (*PriceResponse, error)
	CalculateRoute(context.Context, *RouteRequest) (*RouteResponse, error)
	FindTheNearestDriver(context.Context, *DriverRequest) (*DriverResponse, error)
	TestFunc(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedShippingServiceServer()
}

// UnimplementedShippingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShippingServiceServer struct{}

func (UnimplementedShippingServiceServer) CalculatePrice(context.Context, *PriceRequest) (*PriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculatePrice not implemented")
}
func (UnimplementedShippingServiceServer) CalculateRoute(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateRoute not implemented")
}
func (UnimplementedShippingServiceServer) FindTheNearestDriver(context.Context, *DriverRequest) (*DriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTheNearestDriver not implemented")
}
func (UnimplementedShippingServiceServer) TestFunc(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestFunc not implemented")
}
func (UnimplementedShippingServiceServer) mustEmbedUnimplementedShippingServiceServer() {}
func (UnimplementedShippingServiceServer) testEmbeddedByValue()                         {}

// UnsafeShippingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServiceServer will
// result in compilation errors.
type UnsafeShippingServiceServer interface {
	mustEmbedUnimplementedShippingServiceServer()
}

func RegisterShippingServiceServer(s grpc.ServiceRegistrar, srv ShippingServiceServer) {
	// If the following call pancis, it indicates UnimplementedShippingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShippingService_ServiceDesc, srv)
}

func _ShippingService_CalculatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CalculatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_CalculatePrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CalculatePrice(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_CalculateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CalculateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_CalculateRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CalculateRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_FindTheNearestDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).FindTheNearestDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_FindTheNearestDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).FindTheNearestDriver(ctx, req.(*DriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_TestFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).TestFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_TestFunc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).TestFunc(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingService_ServiceDesc is the grpc.ServiceDesc for ShippingService service.
// It's only intended for direct use with grpcapi.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "truckgo.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculatePrice",
			Handler:    _ShippingService_CalculatePrice_Handler,
		},
		{
			MethodName: "CalculateRoute",
			Handler:    _ShippingService_CalculateRoute_Handler,
		},
		{
			MethodName: "FindTheNearestDriver",
			Handler:    _ShippingService_FindTheNearestDriver_Handler,
		},
		{
			MethodName: "TestFunc",
			Handler:    _ShippingService_TestFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipping.proto",
}
