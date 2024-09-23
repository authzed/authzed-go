package authzed_test

import (
	"context"
	"log"
	"testing"

	"github.com/authzed/grpcutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
)

func ExampleNewClient() {
	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		log.Fatalf("failed to load system certs: %s", err)
	}
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithBearerToken("tc_my_token_deadbeefdeadbeefdeadbeef"),
		systemCerts,
	)
	if err != nil {
		log.Fatalf("failed to connect to authzed: %s", err)
	}
	log.Println(client)
}

func TestWriteSchemaCall(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	// TODO: should we get a handle on the connection in order to be able to close it?
	// It should only matter in testing, but it could still be a problem.
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("some token"),
	)
	require.NoError(err)

	schema := `
	definition document {
		relation reader: user
	}
	definition user {}
	`

	_, err = client.SchemaServiceClient.WriteSchema(context.Background(), &v1.WriteSchemaRequest{Schema: schema})
	require.NoError(err)
}
