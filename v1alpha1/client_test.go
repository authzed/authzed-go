package authzed_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/authzed/grpcutil"

	authzed "github.com/authzed/authzed-go/v1alpha1"
)

func ExampleNewClient(_ *testing.T) {
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithBearerToken("tc_my_token_deadbeefdeadbeefdeadbeef"),
		grpcutil.WithSystemCerts(false),
	)
	if err != nil {
		log.Fatalf("failed to connect to authzed: %s", err)
	}
	fmt.Println(client)
}
