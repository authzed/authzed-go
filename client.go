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

type Client struct {
	api.ACLServiceClient
	api.NamespaceServiceClient
}

type grpcMetadataCreds map[string]string

func (gmc grpcMetadataCreds) RequireTransportSecurity() bool { return true }
func (gmc grpcMetadataCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return gmc, nil
}

type CertVerification bool

const (
	VerifyCA     CertVerification = false
	SkipVerifyCA CertVerification = true
)

func Token(token string) grpc.DialOption {
	return grpc.WithPerRPCCredentials(grpcMetadataCreds{"authorization": "Bearer " + token})
}

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
