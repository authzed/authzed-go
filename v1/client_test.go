package authzed_test

import (
	"log"

	"github.com/authzed/grpcutil"

	authzed "github.com/authzed/authzed-go/v1"
)

func ExampleNewClient() {
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithBearerToken("tc_my_token_deadbeefdeadbeefdeadbeef"),
		grpcutil.WithSystemCerts(grpcutil.VerifyCA),
	)
	if err != nil {
		log.Fatalf("failed to connect to authzed: %s", err)
	}
	log.Println(client)
}
