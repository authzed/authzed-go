// Package pool implements experimental pooling of Authzed clients.
//
// Without pooling, Authzed clients are thread-safe and support concurrent
// requests. Thus, pooling should only be used after instrumentation has proven
// that the application would benefit from the additional throughput.
package pool

import (
	"context"
	"time"

	"github.com/jzelinskie/stringz"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
)

// Client implements an Authzed client with an additional Close() method that
// returns the client connection to a connection pool.
type Client struct {
	*authzed.Client
	closefn func() error
}

// Close returns the client connection to its connection pool.
func (c *Client) Close() error {
	return c.closefn()
}

// ClientPool represents a pool of reusable client connections to Authzed.
//
// Without pooling, Authzed clients are thread-safe and support concurrent
// requests. Thus, pooling should only be used after instrumentation has proven
// that the application would benefit from the additional throughput.
type ClientPool struct {
	connPool *grpcpool.Pool
}

// Get fetches a connection from the pool and allocates a new client using it.
//
// The returned client should have its Close() method called to return the
// connection to the pool.
func (p *ClientPool) Get(ctx context.Context) (*Client, error) {
	conn, err := p.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		&authzed.Client{
			api.NewACLServiceClient(conn),
			api.NewNamespaceServiceClient(conn),
		},
		conn.Close,
	}, nil
}

// NewClientPool returns a new pool of reusable Authzed clients.
func NewClientPool(endpoint string, poolSize int, opts ...grpc.DialOption) (*ClientPool, error) {
	var factory grpcpool.Factory
	factory = func() (*grpc.ClientConn, error) {
		return grpc.Dial(stringz.DefaultEmpty(endpoint, "grpc.authzed.com:443"), opts...)
	}

	pool, err := grpcpool.New(factory, poolSize, poolSize, time.Second)
	if err != nil {
		return nil, err
	}

	return &ClientPool{pool}, nil
}
