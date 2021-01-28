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
	client, err := authzed.NewClient("t_my_token_1235768deadbeef")
	if err != nil {
		log.Fatalf("unable to init client: %s", err)
	}

	// Create some objects to work with.
	doc1 := Document("doc_id_1")
	ceo := User("user_id_1")
	techwriter := User("user_id_2")
	reviewer := User("user_id_3")

	// Writing tuples to Authzed produces a revision value.
	// Providing these revisions to read calls is what guarantees freshness.
	//
	// We recommend saving this after every call to Write or ContentChangeCheck
	// and storing it alongside the object referenced in the write or check.
	revision, err := client.Write(authzed.CREATE,
		ceo.Is("manager").Of(doc1),
		techwriter.Is("contributor").Of(doc1),
		reviewer.Is("viewer").Of(doc1),
	)
	if err != nil {
		log.Fatalf("unable to write updates: %s", err)
	}

	// Create some assertions to check.
	expected := []struct {
		tuple   Tuple
		allowed bool
	}{
		{ceo.Is("viewer").Of(doc1), true},
		{ceo.Is("contributor").Of(doc1), true},
	}

	for _, tt := range expected {
		// Execute a check against the server.
		allowed, err := client.Check(tt.tuple, revision)
		if err != nil {
			log.Fatalf("unable to check access: %s", err)
		}

		if allowed != tt.allowed {
			log.Fatalf("check for %#v returned wrong result: %#v", tt, resp)
		}
	}
}
