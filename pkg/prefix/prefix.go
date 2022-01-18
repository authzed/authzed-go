package prefix

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKeyType string

// ctxPrefixKey is for storing/fetching from a context
var ctxPrefixKey ctxKeyType = "prefix"

// grpcMDPrefixKey is for storing/fetching from grpc MD (must be a string)
const grpcMDPrefixKey = "spicedb-prefix"

// WithUnaryClientInterceptor annotates requests with the desired prefix
func WithUnaryClientInterceptor(prefix string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, callOpts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(ctx, grpcMDPrefixKey, prefix)
		return invoker(ctx, method, req, reply, cc, callOpts...)
	}
}

// WithStreamClientInterceptor annotates requests with the desired prefix
func WithStreamClientInterceptor(prefix string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, grpcMDPrefixKey, prefix)
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// FromContext reads the prefix from the context
func FromContext(ctx context.Context) string {
	if c := ctx.Value(ctxPrefixKey); c != nil {
		return c.(string)
	}
	return ""
}

// MustFromContext reads the prefix from the context and panics if it has
// not been set on the context.
func MustFromContext(ctx context.Context) string {
	prefix := FromContext(ctx)
	if prefix == "" {
		panic("prefix missing")
	}
	return prefix
}

// UnaryServerInterceptor returns a new unary server interceptor that extracts
// the prefix from the request metadata and adds it to the context
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	prefix := metautils.ExtractIncoming(ctx).Get(grpcMDPrefixKey)
	return handler(context.WithValue(ctx, ctxPrefixKey, prefix), req)
}

// StreamServerInterceptor returns a new stream server interceptor that extracts
// the prefix from the request metadata and adds it to the context
func StreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrapper := &recvWrapper{stream, stream.Context(), stream.Context()}
	return handler(srv, wrapper)
}

type recvWrapper struct {
	grpc.ServerStream
	initialCtx context.Context
	ctx        context.Context
}

func (s *recvWrapper) Context() context.Context {
	return s.ctx
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	prefix := metautils.ExtractIncoming(s.initialCtx).Get(grpcMDPrefixKey)
	s.ctx = context.WithValue(s.initialCtx, ctxPrefixKey, prefix)
	return nil
}
