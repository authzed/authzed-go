package authzed

import (
	"github.com/jzelinskie/stringz"
	"google.golang.org/grpc"

	v1alpha1 "github.com/authzed/authzed-go/proto/authzed/api/v1alpha1"
)

// Client represents an open connection to Authzed.
//
// Clients are backed by a gRPC client and as such are thread-safe.
type Client struct {
	v1alpha1.SchemaServiceClient
}

// NewClient initializes a brand new client for interacting with Authzed.
func NewClient(endpoint string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.Dial(
		stringz.DefaultEmpty(endpoint, "grpc.authzed.com:443"),
		opts...,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		v1alpha1.NewSchemaServiceClient(conn),
	}, nil
}
