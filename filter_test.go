package authzed

import (
	"context"
	"os"
	"testing"

	api "github.com/authzed/authzed-go/arrakisapi/api"
)

var fred = &api.User{UserOneof: &api.User_Userset{
	Userset: &api.ObjectAndRelation{
		Namespace: "petricorp_jimmy_test_dev/user",
		ObjectId:  "fred",
		Relation:  "...",
	},
}}

type doc struct {
	ID string
}

var _ Checkable = doc{}

func (d doc) AsObjectAndRelation(relation string) *api.ObjectAndRelation {
	return &api.ObjectAndRelation{
		Namespace: "petricorp_jimmy_test_dev/document",
		ObjectId:  d.ID,
		Relation:  relation,
	}
}

func TestFilterIter(t *testing.T) {
	client, err := NewClient(
		"grpc.authzed.com:443",
		Token(os.Getenv("AUTHZED_GO_TOKEN")),
		SystemCerts(VerifyCA),
	)
	if err != nil {
		t.Fatal(err)
	}

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
