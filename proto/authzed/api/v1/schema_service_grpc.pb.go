// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: authzed/api/v1/schema_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SchemaService_ReadSchema_FullMethodName            = "/authzed.api.v1.SchemaService/ReadSchema"
	SchemaService_WriteSchema_FullMethodName           = "/authzed.api.v1.SchemaService/WriteSchema"
	SchemaService_ReflectSchema_FullMethodName         = "/authzed.api.v1.SchemaService/ReflectSchema"
	SchemaService_ComputablePermissions_FullMethodName = "/authzed.api.v1.SchemaService/ComputablePermissions"
	SchemaService_DependentRelations_FullMethodName    = "/authzed.api.v1.SchemaService/DependentRelations"
	SchemaService_DiffSchema_FullMethodName            = "/authzed.api.v1.SchemaService/DiffSchema"
)

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SchemaService implements operations on a Permissions System's Schema.
type SchemaServiceClient interface {
	// Read returns the current Object Definitions for a Permissions System.
	//
	// Errors include:
	// - INVALID_ARGUMENT: a provided value has failed to semantically validate
	// - NOT_FOUND: no schema has been defined
	ReadSchema(ctx context.Context, in *ReadSchemaRequest, opts ...grpc.CallOption) (*ReadSchemaResponse, error)
	// Write overwrites the current Object Definitions for a Permissions System.
	WriteSchema(ctx context.Context, in *WriteSchemaRequest, opts ...grpc.CallOption) (*WriteSchemaResponse, error)
	// ReflectSchema reflects the current schema stored in SpiceDB, returning a structural
	// form of the schema for use by client tooling.
	ReflectSchema(ctx context.Context, in *ReflectSchemaRequest, opts ...grpc.CallOption) (*ReflectSchemaResponse, error)
	// ComputablePermissions returns the set of permissions that compute based off a relation
	// in the current schema. For example, if the schema has a relation `viewer` and a permission
	// `view` defined as `permission view = viewer + editor`, then the
	// computable permissions for the relation `viewer` will include `view`.
	ComputablePermissions(ctx context.Context, in *ComputablePermissionsRequest, opts ...grpc.CallOption) (*ComputablePermissionsResponse, error)
	// DependentRelations returns the set of relations and permissions that used
	// to compute a permission, recursively, in the current schema. It is the
	// inverse of the ComputablePermissions API.
	DependentRelations(ctx context.Context, in *DependentRelationsRequest, opts ...grpc.CallOption) (*DependentRelationsResponse, error)
	// DiffSchema returns the difference between the specified schema and the current
	// schema stored in SpiceDB.
	DiffSchema(ctx context.Context, in *DiffSchemaRequest, opts ...grpc.CallOption) (*DiffSchemaResponse, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) ReadSchema(ctx context.Context, in *ReadSchemaRequest, opts ...grpc.CallOption) (*ReadSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_ReadSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) WriteSchema(ctx context.Context, in *WriteSchemaRequest, opts ...grpc.CallOption) (*WriteSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_WriteSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ReflectSchema(ctx context.Context, in *ReflectSchemaRequest, opts ...grpc.CallOption) (*ReflectSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReflectSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_ReflectSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ComputablePermissions(ctx context.Context, in *ComputablePermissionsRequest, opts ...grpc.CallOption) (*ComputablePermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ComputablePermissionsResponse)
	err := c.cc.Invoke(ctx, SchemaService_ComputablePermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) DependentRelations(ctx context.Context, in *DependentRelationsRequest, opts ...grpc.CallOption) (*DependentRelationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DependentRelationsResponse)
	err := c.cc.Invoke(ctx, SchemaService_DependentRelations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) DiffSchema(ctx context.Context, in *DiffSchemaRequest, opts ...grpc.CallOption) (*DiffSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiffSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_DiffSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility.
//
// SchemaService implements operations on a Permissions System's Schema.
type SchemaServiceServer interface {
	// Read returns the current Object Definitions for a Permissions System.
	//
	// Errors include:
	// - INVALID_ARGUMENT: a provided value has failed to semantically validate
	// - NOT_FOUND: no schema has been defined
	ReadSchema(context.Context, *ReadSchemaRequest) (*ReadSchemaResponse, error)
	// Write overwrites the current Object Definitions for a Permissions System.
	WriteSchema(context.Context, *WriteSchemaRequest) (*WriteSchemaResponse, error)
	// ReflectSchema reflects the current schema stored in SpiceDB, returning a structural
	// form of the schema for use by client tooling.
	ReflectSchema(context.Context, *ReflectSchemaRequest) (*ReflectSchemaResponse, error)
	// ComputablePermissions returns the set of permissions that compute based off a relation
	// in the current schema. For example, if the schema has a relation `viewer` and a permission
	// `view` defined as `permission view = viewer + editor`, then the
	// computable permissions for the relation `viewer` will include `view`.
	ComputablePermissions(context.Context, *ComputablePermissionsRequest) (*ComputablePermissionsResponse, error)
	// DependentRelations returns the set of relations and permissions that used
	// to compute a permission, recursively, in the current schema. It is the
	// inverse of the ComputablePermissions API.
	DependentRelations(context.Context, *DependentRelationsRequest) (*DependentRelationsResponse, error)
	// DiffSchema returns the difference between the specified schema and the current
	// schema stored in SpiceDB.
	DiffSchema(context.Context, *DiffSchemaRequest) (*DiffSchemaResponse, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSchemaServiceServer struct{}

func (UnimplementedSchemaServiceServer) ReadSchema(context.Context, *ReadSchemaRequest) (*ReadSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSchema not implemented")
}
func (UnimplementedSchemaServiceServer) WriteSchema(context.Context, *WriteSchemaRequest) (*WriteSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteSchema not implemented")
}
func (UnimplementedSchemaServiceServer) ReflectSchema(context.Context, *ReflectSchemaRequest) (*ReflectSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReflectSchema not implemented")
}
func (UnimplementedSchemaServiceServer) ComputablePermissions(context.Context, *ComputablePermissionsRequest) (*ComputablePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputablePermissions not implemented")
}
func (UnimplementedSchemaServiceServer) DependentRelations(context.Context, *DependentRelationsRequest) (*DependentRelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DependentRelations not implemented")
}
func (UnimplementedSchemaServiceServer) DiffSchema(context.Context, *DiffSchemaRequest) (*DiffSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffSchema not implemented")
}
func (UnimplementedSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {}
func (UnimplementedSchemaServiceServer) testEmbeddedByValue()                       {}

// UnsafeSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServiceServer will
// result in compilation errors.
type UnsafeSchemaServiceServer interface {
	mustEmbedUnimplementedSchemaServiceServer()
}

func RegisterSchemaServiceServer(s grpc.ServiceRegistrar, srv SchemaServiceServer) {
	// If the following call pancis, it indicates UnimplementedSchemaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SchemaService_ServiceDesc, srv)
}

func _SchemaService_ReadSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ReadSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ReadSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ReadSchema(ctx, req.(*ReadSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_WriteSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).WriteSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_WriteSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).WriteSchema(ctx, req.(*WriteSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ReflectSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReflectSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ReflectSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ReflectSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ReflectSchema(ctx, req.(*ReflectSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ComputablePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputablePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ComputablePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ComputablePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ComputablePermissions(ctx, req.(*ComputablePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_DependentRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependentRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).DependentRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_DependentRelations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).DependentRelations(ctx, req.(*DependentRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_DiffSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).DiffSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_DiffSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).DiffSchema(ctx, req.(*DiffSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaService_ServiceDesc is the grpc.ServiceDesc for SchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.v1.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadSchema",
			Handler:    _SchemaService_ReadSchema_Handler,
		},
		{
			MethodName: "WriteSchema",
			Handler:    _SchemaService_WriteSchema_Handler,
		},
		{
			MethodName: "ReflectSchema",
			Handler:    _SchemaService_ReflectSchema_Handler,
		},
		{
			MethodName: "ComputablePermissions",
			Handler:    _SchemaService_ComputablePermissions_Handler,
		},
		{
			MethodName: "DependentRelations",
			Handler:    _SchemaService_DependentRelations_Handler,
		},
		{
			MethodName: "DiffSchema",
			Handler:    _SchemaService_DiffSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authzed/api/v1/schema_service.proto",
}
