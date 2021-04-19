package authzed

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/jzelinskie/stringz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client represents an open connection to Authzed.
//
// Clients are backed by a gRPC client and as such are thread-safe.
type Client struct {
	api.ACLServiceClient
	api.NamespaceServiceClient
}

type grpcMetadataCreds map[string]string

func (gmc grpcMetadataCreds) RequireTransportSecurity() bool { return true }
func (gmc grpcMetadataCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return gmc, nil
}

// CertVerification is an enumeration of how secure TLS should be configured.
type CertVerification bool

const (
	// VerifyCA will verify the certificate authority has been verified.
	VerifyCA CertVerification = false

	// SkipVerifyCA will not verify the certificate authority when using TLS.
	SkipVerifyCA CertVerification = true
)

// Token is the client option that is used for authenticating to Authzed.
func Token(token string) grpc.DialOption {
	return grpc.WithPerRPCCredentials(grpcMetadataCreds{"authorization": "Bearer " + token})
}

// SystemCerts is the client option that is used for establish a secure
// connection to Authzed.
func SystemCerts(v CertVerification) grpc.DialOption {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: bool(v),
	}))
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
		api.NewACLServiceClient(conn),
		api.NewNamespaceServiceClient(conn),
	}, nil
}
