package authzed_test

import (
	"log"

	"github.com/authzed/grpcutil"

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
