// +build integration

package authzed

import (
	"context"
	"sync"
	"testing"

	"google.golang.org/grpc"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/authzed/authzed-go/x/nsbuilder"
)

var setupOnce sync.Once

var (
	namespaceUser = nsbuilder.Namespace("test/user")
	namespaceDoc  = nsbuilder.Namespace("test/document",
		nsbuilder.Relation("writer", nil),
		nsbuilder.Relation("reader", nsbuilder.Union(nsbuilder.This(), nsbuilder.ComputedUserset("writer"))),
	)
)

var (
	fred = &api.User{UserOneof: &api.User_Userset{
		Userset: &api.ObjectAndRelation{
			Namespace: "test/user",
			ObjectId:  "fred",
			Relation:  "...",
		},
	}}
)

func setupTenant(t *testing.T) *Client {
	client, err := NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	setupOnce.Do(func() {
		for _, nsdef := range []*api.NamespaceDefinition{namespaceUser, namespaceDoc} {
			if _, err := client.WriteConfig(context.Background(), &api.WriteConfigRequest{Config: nsdef}); err != nil {
				t.Fatal(err)
			}
		}

		_, err := client.Write(context.Background(), &api.WriteRequest{
			Updates: []*api.RelationTupleUpdate{
				{
					Operation: api.RelationTupleUpdate_CREATE,
					Tuple: &api.RelationTuple{
						ObjectAndRelation: &api.ObjectAndRelation{
							Namespace: "test/document",
							ObjectId:  "firstdoc",
							Relation:  "writer",
						},
						User: fred,
					},
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	return client
}
