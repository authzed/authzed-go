package authzed_test

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
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

	_, err = client.Check(context.Background(), &api.CheckRequest{
		TestUserset: &api.ObjectAndRelation{
			Namespace: "mytenant/document",
			ObjectId:  "readme",
			Relation:  "viewer",
		},
		User: &api.User{UserOneof: &api.User_Userset{
			Userset: &api.ObjectAndRelation{
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
