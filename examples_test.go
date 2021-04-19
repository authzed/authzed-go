package authzed_test

import (
	"log"

	"github.com/authzed/authzed-go"
)

func ExampleNewClient() {
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		authzed.Token("my_token_deadbeefdeadbeefdeadbeef"),
		authzed.SystemCerts(authzed.VerifyCA),
	)
	if err != nil {
		log.Fatal(err)
	}
}
