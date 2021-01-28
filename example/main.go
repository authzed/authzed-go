package main

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
)

var (
	// Define the types of objects and their relations.
	Document = authzed.NewNamespace("yourtenant", "document", "manager", "contributor", "viewer")
	User     = authzed.NewTerminal("yourtenant", "user")
)

func main() {
	// Initialize a new gRPC Client.
	opts, err := authzed.NewGrpcClientOptions("t_my_token_1234567deadbeef")
	if err != nil {
		log.Fatalf("unable to init client options: %s", err)
	}

	client, err := authzed.NewGrpcClient(opts)
	if err != nil {
		log.Fatalf("unable to init client: %s", err)
	}

	// Create some objects to work with.
	doc1 := Document("doc_id_1")
	ceo := User("user_id_1")
	techwriter := User("user_id_2")
	reviewer := User("user_id_3")

	// Assign some relationships to these instances.
	createTuples := []*api.RelationTuple{
		ceo.Is("manager").Of(doc1),
		techwriter.Is("contributor").Of(doc1),
		reviewer.Is("viewer").Of(doc1),
	}

	var tupleUpdates []*api.RelationTupleUpdate
	for _, tpl := range createTuples {
		tupleUpdates = append(tupleUpdates, &api.RelationTupleUpdate{
			Operation: api.RelationTupleUpdate_CREATE,
			Tuple:     tpl,
		})
	}

	writeResp, err := client.Write(context.TODO(), &api.WriteRequest{
		Updates: tupleUpdates,
	})
	if err != nil {
		log.Fatalf("unable to write updates: %s", err)
	}

	// Save the revision from our previous call, so that we can guarantee future
	// queries are at least that fresh.
	//
	// We recommend saving this after every call to Write or ContentChangeCheck
	// and storing it alongside the object referenced in the write or check.
	revision := writeResp.Revision

	// Create some assertions to check.
	expected := []struct {
		Tuple   *api.RelationTuple
		Allowed bool
	}{
		{ceo.Is("viewer").Of(doc1), true},
		{ceo.Is("contributor").Of(doc1), true},
	}

	for _, tt := range expected {
		// Execute a check against the server.
		resp, err := client.Check(context.TODO(), &api.CheckRequest{
			TestUserset: tt.Tuple.ObjectAndRelation,
			User:        tt.Tuple.User,
			AtRevision:  revision,
		})
		if err != nil {
			log.Fatalf("unable to check access: %s", err)
		}

		if resp.IsMember != tt.Allowed {
			log.Fatalf("check for %#v returned wrong result: %#v", tt, resp)
		}
	}
}
