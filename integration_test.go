// +build integration

package authzed

import (
	"context"
	"sync"
	"testing"

	"google.golang.org/grpc"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
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
	fred = &v0.User{UserOneof: &v0.User_Userset{
		Userset: &v0.ObjectAndRelation{
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
		if _, err := client.WriteConfig(context.Background(), &v0.WriteConfigRequest{Configs: []*v0.NamespaceDefinition{
			namespaceUser,
			namespaceDoc,
		}}); err != nil {
			t.Fatal(err)
		}

		_, err := client.Write(context.Background(), &v0.WriteRequest{
			Updates: []*v0.RelationTupleUpdate{
				{
					Operation: v0.RelationTupleUpdate_CREATE,
					Tuple: &v0.RelationTuple{
						ObjectAndRelation: &v0.ObjectAndRelation{
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
