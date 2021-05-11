// +build integration

package authzed

import (
	"context"
	"testing"

	api "github.com/authzed/authzed-go/arrakisapi/api"
)

type doc struct {
	ID string
}

var _ Checkable = doc{}

func (d doc) AsObjectAndRelation(relation string) *api.ObjectAndRelation {
	return &api.ObjectAndRelation{
		Namespace: "test/document",
		ObjectId:  d.ID,
		Relation:  relation,
	}
}

func TestFilterIter(t *testing.T) {
	client := setupTenant(t)
	docs := []doc{{ID: "firstdoc"}, {ID: "seconddoc"}}
	iter := client.NewFilterIter(docs, fred, "reader", nil)

	var allowedDocs []doc
	for iter.Next(context.Background()) {
		allowedDocs = append(allowedDocs, iter.Item().(doc))
	}
	if err := iter.Err(); err != nil {
		t.Fatal(err)
	}

	if len(allowedDocs) != 1 {
		t.Fatal("expected fred to only have access one item")
	}

	if allowedDocs[0].ID != "firstdoc" {
		t.Fatal("expected fred to only have access to firstdoc")
	}
}
