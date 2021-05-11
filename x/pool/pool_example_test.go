package pool_test

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/authzed/authzed-go/x/pool"
)

func ExampleNewClientPool() {
	clientPool, err := pool.NewClientPool(
		"grpc.authzed.com:443",
		10,
		authzed.Token("my_token_deadbeefdeadbeefdeadbeef"),
		authzed.SystemCerts(authzed.VerifyCA),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := clientPool.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close() // Returns client to pool; does not disconnect client.

	_, err = client.Check(ctx, &api.CheckRequest{
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
