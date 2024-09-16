// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: authzed/api/materialize/v0/watchpermissionsets.proto

package v0

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
	WatchPermissionSetsService_WatchPermissionSets_FullMethodName  = "/authzed.api.materialize.v0.WatchPermissionSetsService/WatchPermissionSets"
	WatchPermissionSetsService_LookupPermissionSets_FullMethodName = "/authzed.api.materialize.v0.WatchPermissionSetsService/LookupPermissionSets"
)

// WatchPermissionSetsServiceClient is the client API for WatchPermissionSetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchPermissionSetsServiceClient interface {
	// WatchPermissionSets returns a stream of changes to the sets which can be used to compute the watched permissions.
	//
	// WatchPermissionSets lets consumers achieve the same thing as WatchPermissions, but trades off a simpler usage model with
	// significantly lower computational requirements. Unlike WatchPermissions, this method returns changes to the sets of permissions,
	// rather than the individual permissions. Permission sets are a normalized form of the computed permissions, which
	// means that the consumer must perform an extra computation over this representation to obtain the final computed
	// permissions, typically by intersecting the provided sets.
	//
	// For example, this would look like a JOIN between the
	// materialize permission sets table in a target relation database, the table with the resources to authorize access
	// to, and the table with the subject (e.g. a user).
	//
	// In exchange, the number of changes issued by WatchPermissionSets will be several orders of magnitude less than those
	// emitted by WatchPermissions, which has several implications:
	// - significantly less resources to compute the sets
	// - significantly less messages to stream over the network
	// - significantly less events to ingest on the consumer side
	// - less ingestion lag from the origin SpiceDB mutation
	//
	// The type of scenarios WatchPermissionSets is particularly well suited is when a single change
	// in the origin SpiceDB can yield millions of changes. For example, in the GitHub authorization model, assigning a role
	// to a top-level team of an organization with hundreds of thousands of employees can lead to an explosion of
	// permission change events that would require a lot of computational resources to process, both on Materialize and
	// the consumer side.
	//
	// WatchPermissionSets is thus recommended for any larger scale use case where the fan-out in permission changes that
	// emerges from a specific schema and data shape is too large to handle effectively.
	//
	// The API does not offer a sharding mechanism and thus there should only be one consumer per target system.
	// Implementing an active-active HA consumer setup over the same target system will require coordinating which
	// revisions have been consumed in order to prevent transitioning to an inconsistent state.
	WatchPermissionSets(ctx context.Context, in *WatchPermissionSetsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchPermissionSetsResponse], error)
	// LookupPermissionSets returns the current state of the permission sets which can be used to derive the computed permissions.
	// It's typically used to backfill the state of the permission sets in the consumer side.
	//
	// It's a cursored API and the consumer is responsible to keep track of the cursor and use it on each subsequent call.
	// Each stream will return <N> permission sets defined by the specified request limit. The server will keep streaming until
	// the sets per stream is hit, or the current state of the sets is reached,
	// whatever happens first, and then close the stream. The server will indicate there are no more changes to stream
	// through the `completed_members` in the cursor.
	//
	// There may be many elements to stream, and so the consumer should be prepared to resume the stream from the last
	// cursor received. Once completed, the consumer may start streaming permission set changes using WatchPermissionSets
	// and the revision token from the last LookupPermissionSets response.
	LookupPermissionSets(ctx context.Context, in *LookupPermissionSetsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LookupPermissionSetsResponse], error)
}

type watchPermissionSetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchPermissionSetsServiceClient(cc grpc.ClientConnInterface) WatchPermissionSetsServiceClient {
	return &watchPermissionSetsServiceClient{cc}
}

func (c *watchPermissionSetsServiceClient) WatchPermissionSets(ctx context.Context, in *WatchPermissionSetsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchPermissionSetsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WatchPermissionSetsService_ServiceDesc.Streams[0], WatchPermissionSetsService_WatchPermissionSets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchPermissionSetsRequest, WatchPermissionSetsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WatchPermissionSetsService_WatchPermissionSetsClient = grpc.ServerStreamingClient[WatchPermissionSetsResponse]

func (c *watchPermissionSetsServiceClient) LookupPermissionSets(ctx context.Context, in *LookupPermissionSetsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LookupPermissionSetsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WatchPermissionSetsService_ServiceDesc.Streams[1], WatchPermissionSetsService_LookupPermissionSets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LookupPermissionSetsRequest, LookupPermissionSetsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WatchPermissionSetsService_LookupPermissionSetsClient = grpc.ServerStreamingClient[LookupPermissionSetsResponse]

// WatchPermissionSetsServiceServer is the server API for WatchPermissionSetsService service.
// All implementations must embed UnimplementedWatchPermissionSetsServiceServer
// for forward compatibility.
type WatchPermissionSetsServiceServer interface {
	// WatchPermissionSets returns a stream of changes to the sets which can be used to compute the watched permissions.
	//
	// WatchPermissionSets lets consumers achieve the same thing as WatchPermissions, but trades off a simpler usage model with
	// significantly lower computational requirements. Unlike WatchPermissions, this method returns changes to the sets of permissions,
	// rather than the individual permissions. Permission sets are a normalized form of the computed permissions, which
	// means that the consumer must perform an extra computation over this representation to obtain the final computed
	// permissions, typically by intersecting the provided sets.
	//
	// For example, this would look like a JOIN between the
	// materialize permission sets table in a target relation database, the table with the resources to authorize access
	// to, and the table with the subject (e.g. a user).
	//
	// In exchange, the number of changes issued by WatchPermissionSets will be several orders of magnitude less than those
	// emitted by WatchPermissions, which has several implications:
	// - significantly less resources to compute the sets
	// - significantly less messages to stream over the network
	// - significantly less events to ingest on the consumer side
	// - less ingestion lag from the origin SpiceDB mutation
	//
	// The type of scenarios WatchPermissionSets is particularly well suited is when a single change
	// in the origin SpiceDB can yield millions of changes. For example, in the GitHub authorization model, assigning a role
	// to a top-level team of an organization with hundreds of thousands of employees can lead to an explosion of
	// permission change events that would require a lot of computational resources to process, both on Materialize and
	// the consumer side.
	//
	// WatchPermissionSets is thus recommended for any larger scale use case where the fan-out in permission changes that
	// emerges from a specific schema and data shape is too large to handle effectively.
	//
	// The API does not offer a sharding mechanism and thus there should only be one consumer per target system.
	// Implementing an active-active HA consumer setup over the same target system will require coordinating which
	// revisions have been consumed in order to prevent transitioning to an inconsistent state.
	WatchPermissionSets(*WatchPermissionSetsRequest, grpc.ServerStreamingServer[WatchPermissionSetsResponse]) error
	// LookupPermissionSets returns the current state of the permission sets which can be used to derive the computed permissions.
	// It's typically used to backfill the state of the permission sets in the consumer side.
	//
	// It's a cursored API and the consumer is responsible to keep track of the cursor and use it on each subsequent call.
	// Each stream will return <N> permission sets defined by the specified request limit. The server will keep streaming until
	// the sets per stream is hit, or the current state of the sets is reached,
	// whatever happens first, and then close the stream. The server will indicate there are no more changes to stream
	// through the `completed_members` in the cursor.
	//
	// There may be many elements to stream, and so the consumer should be prepared to resume the stream from the last
	// cursor received. Once completed, the consumer may start streaming permission set changes using WatchPermissionSets
	// and the revision token from the last LookupPermissionSets response.
	LookupPermissionSets(*LookupPermissionSetsRequest, grpc.ServerStreamingServer[LookupPermissionSetsResponse]) error
	mustEmbedUnimplementedWatchPermissionSetsServiceServer()
}

// UnimplementedWatchPermissionSetsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWatchPermissionSetsServiceServer struct{}

func (UnimplementedWatchPermissionSetsServiceServer) WatchPermissionSets(*WatchPermissionSetsRequest, grpc.ServerStreamingServer[WatchPermissionSetsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method WatchPermissionSets not implemented")
}
func (UnimplementedWatchPermissionSetsServiceServer) LookupPermissionSets(*LookupPermissionSetsRequest, grpc.ServerStreamingServer[LookupPermissionSetsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method LookupPermissionSets not implemented")
}
func (UnimplementedWatchPermissionSetsServiceServer) mustEmbedUnimplementedWatchPermissionSetsServiceServer() {
}
func (UnimplementedWatchPermissionSetsServiceServer) testEmbeddedByValue() {}

// UnsafeWatchPermissionSetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchPermissionSetsServiceServer will
// result in compilation errors.
type UnsafeWatchPermissionSetsServiceServer interface {
	mustEmbedUnimplementedWatchPermissionSetsServiceServer()
}

func RegisterWatchPermissionSetsServiceServer(s grpc.ServiceRegistrar, srv WatchPermissionSetsServiceServer) {
	// If the following call pancis, it indicates UnimplementedWatchPermissionSetsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WatchPermissionSetsService_ServiceDesc, srv)
}

func _WatchPermissionSetsService_WatchPermissionSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPermissionSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchPermissionSetsServiceServer).WatchPermissionSets(m, &grpc.GenericServerStream[WatchPermissionSetsRequest, WatchPermissionSetsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WatchPermissionSetsService_WatchPermissionSetsServer = grpc.ServerStreamingServer[WatchPermissionSetsResponse]

func _WatchPermissionSetsService_LookupPermissionSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupPermissionSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchPermissionSetsServiceServer).LookupPermissionSets(m, &grpc.GenericServerStream[LookupPermissionSetsRequest, LookupPermissionSetsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WatchPermissionSetsService_LookupPermissionSetsServer = grpc.ServerStreamingServer[LookupPermissionSetsResponse]

// WatchPermissionSetsService_ServiceDesc is the grpc.ServiceDesc for WatchPermissionSetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchPermissionSetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.materialize.v0.WatchPermissionSetsService",
	HandlerType: (*WatchPermissionSetsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPermissionSets",
			Handler:       _WatchPermissionSetsService_WatchPermissionSets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LookupPermissionSets",
			Handler:       _WatchPermissionSetsService_LookupPermissionSets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "authzed/api/materialize/v0/watchpermissionsets.proto",
}
