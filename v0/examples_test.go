package authzed_test

import (
	"context"
	"log"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	"github.com/authzed/authzed-go/v0"
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

	_, err = client.Check(context.Background(), &v0.CheckRequest{
		TestUserset: &v0.ObjectAndRelation{
			Namespace: "mytenant/document",
			ObjectId:  "readme",
			Relation:  "viewer",
		},
		User: &v0.User{UserOneof: &v0.User_Userset{
			Userset: &v0.ObjectAndRelation{
				Namespace: "mytenant/user",
				ObjectId:  "jimmy",
				Relation:  "...",
			},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
}
