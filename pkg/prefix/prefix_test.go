package prefix

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	grpc_testing "github.com/grpc-ecosystem/go-grpc-middleware/testing"
	pb_testproto "github.com/grpc-ecosystem/go-grpc-middleware/testing/testproto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

const testPrefix = "test_prefix"

type testServer struct{}

func (t testServer) PingEmpty(ctx context.Context, empty *pb_testproto.Empty) (*pb_testproto.PingResponse, error) {
	return &pb_testproto.PingResponse{Value: MustFromContext(ctx)}, nil
}

func (t testServer) Ping(ctx context.Context, request *pb_testproto.PingRequest) (*pb_testproto.PingResponse, error) {
	return &pb_testproto.PingResponse{Value: MustFromContext(ctx)}, nil
}

func (t testServer) PingError(ctx context.Context, request *pb_testproto.PingRequest) (*pb_testproto.Empty, error) {
	return nil, fmt.Errorf("err")
}

func (t testServer) PingList(request *pb_testproto.PingRequest, server pb_testproto.TestService_PingListServer) error {
	return server.Send(&pb_testproto.PingResponse{Value: MustFromContext(server.Context())})
}

func (t testServer) PingStream(stream pb_testproto.TestService_PingStreamServer) error {
	for {
		_, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&pb_testproto.PingResponse{
			Value: MustFromContext(stream.Context()),
		}); err != nil {
			return err
		}
	}
	return nil
}

func TestConsistencyTestSuite(t *testing.T) {
	s := &PrefixMiddlewareTestSuite{
		InterceptorTestSuite: &grpc_testing.InterceptorTestSuite{
			TestService: &testServer{},
			ServerOpts: []grpc.ServerOption{
				grpc.StreamInterceptor(StreamServerInterceptor),
				grpc.UnaryInterceptor(UnaryServerInterceptor),
			},
			ClientOpts: []grpc.DialOption{
				grpc.WithUnaryInterceptor(WithUnaryClientInterceptor(testPrefix)),
				grpc.WithStreamInterceptor(WithStreamClientInterceptor(testPrefix)),
			},
		},
	}
	suite.Run(t, s)
}

var goodPing = &pb_testproto.PingRequest{Value: "something"}

type PrefixMiddlewareTestSuite struct {
	*grpc_testing.InterceptorTestSuite
}

func (s *PrefixMiddlewareTestSuite) TestValidPasses_Unary() {
	resp, err := s.Client.Ping(s.SimpleCtx(), goodPing)
	require.NoError(s.T(), err)
	require.Equal(s.T(), testPrefix, resp.Value)
}

func (s *PrefixMiddlewareTestSuite) TestValidPasses_ServerList() {
	stream, err := s.Client.PingList(s.SimpleCtx(), goodPing)
	require.NoError(s.T(), err)
	resp, err := stream.Recv()
	assert.NoError(s.T(), err, "no error on messages sent occurred")
	require.Equal(s.T(), testPrefix, resp.Value)
}

func (s *PrefixMiddlewareTestSuite) TestValidPasses_ServerStream() {
	stream, err := s.Client.PingStream(s.SimpleCtx())
	require.NoError(s.T(), err)
	for i := 0; i < 3; i++ {
		require.NoError(s.T(), stream.Send(goodPing))
		resp, err := stream.Recv()
		require.NoError(s.T(), err)
		require.Equal(s.T(), testPrefix, resp.Value)
	}
}
