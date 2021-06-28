package pool_test

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
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

	_, err = client.Check(ctx, &v0.CheckRequest{
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
