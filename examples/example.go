package main

import (
	"context"
	"log"

	client "github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
)

const document_ns = "yourtenant/document"
const user_ns = "yourtenant/user"

func main() {
	token := "t_your_token_here_1234567deadbeef"

	options, err := client.NewClientOptions(token)
	if err != nil {
		log.Fatalf("Unable to create client options: %s", err)
	}

	client, err := client.NewClient(options)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}

	// Create some objects
	aDoc := createObject(document_ns, "doc1")
	anOwner := createObject(user_ns, "theowner")("...")
	anEditor := createObject(user_ns, "userwhocanedit")("...")
	aViewer := createObject(user_ns, "viewonlyuser")("...")

	// Create some tuples that represent roles granted between users and objects
	newTuples := []*api.RelationTupleUpdate{
		createTuple(tuple(aDoc("owner"), anOwner)),
		createTuple(tuple(aDoc("contributor"), anEditor)),
		createTuple(tuple(aDoc("viewer"), aViewer)),
	}

	req := api.WriteRequest{
		Updates: newTuples,
	}

	resp, err := client.Write(context.Background(), &req)
	if err != nil {
		log.Fatalf("Unable to write tuples: %s", err)
	}

	// Save the zookie that the call above generated to prevent new enemies
	// We recommend saving this from any call to Write or ContentChangeCheck,
	// and storing it alongside the object referenced in the write or check (in this case aDoc)"
	whenPermsChanged := resp.Revision

	// Run some checks on the written data
	aNobody := createObject(user_ns, "randomnobody")("...")
	expected := []checkData{
		{permission: aDoc("read"), user: anOwner, hasAccess: true},
		{permission: aDoc("write"), user: anOwner, hasAccess: true},
		{permission: aDoc("delete"), user: anOwner, hasAccess: true},
		{permission: aDoc("read"), user: anEditor, hasAccess: true},
		{permission: aDoc("write"), user: anEditor, hasAccess: true},
		{permission: aDoc("delete"), user: anEditor, hasAccess: false},
		{permission: aDoc("read"), user: aViewer, hasAccess: true},
		{permission: aDoc("write"), user: aViewer, hasAccess: false},
		{permission: aDoc("delete"), user: aViewer, hasAccess: false},
		{permission: aDoc("read"), user: aNobody, hasAccess: true},
		{permission: aDoc("write"), user: aNobody, hasAccess: false},
		{permission: aDoc("delete"), user: aNobody, hasAccess: false},
	}

	for _, test := range expected {
		testReq := api.CheckRequest{
			TestUserset: test.permission,
			User: &api.User{
				UserOneof: &api.User_Userset{
					Userset: test.user,
				},
			},
			AtRevision: whenPermsChanged,
		}
		testResp, err := client.Check(context.Background(), &testReq)
		if err != nil {
			log.Fatalf("Unable to run check request: %s", err)
		}

		if testResp.IsMember != test.hasAccess {
			log.Fatalf("Check returned the wrong result: %v", test)
		}
	}
}

type checkData struct {
	permission *api.ObjectAndRelation
	user       *api.ObjectAndRelation
	hasAccess  bool
}

func createObject(namespace, objectID string) func(string) *api.ObjectAndRelation {
	return func(relation string) *api.ObjectAndRelation {
		return &api.ObjectAndRelation{
			Namespace: namespace,
			ObjectId:  objectID,
			Relation:  relation,
		}
	}
}

func tuple(onr *api.ObjectAndRelation, userset *api.ObjectAndRelation) *api.RelationTuple {
	return &api.RelationTuple{
		ObjectAndRelation: onr,
		User: &api.User{
			UserOneof: &api.User_Userset{
				Userset: userset,
			},
		},
	}
}

func createTuple(tpl *api.RelationTuple) *api.RelationTupleUpdate {
	return &api.RelationTupleUpdate{
		Operation: api.RelationTupleUpdate_CREATE,
		Tuple:     tpl,
	}
}
