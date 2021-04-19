package main

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
)

const (
	document_ns = "yourtenant/document"
	user_ns     = "yourtenant/user"
)

func main() {
	// Create an Authzed client.
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		authzed.Token("t_your_token_here_1234567deadbeef"),
		authzed.SystemCerts(authzed.VerifyCA),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	// Create some objects that will be protected by Authzed.
	aDoc := createObject(document_ns, "doc1")
	anOwner := createObject(user_ns, "theowner")("...")
	anEditor := createObject(user_ns, "userwhocanedit")("...")
	aViewer := createObject(user_ns, "viewonlyuser")("...")

	// Create some tuples that represent roles granted between users and objects.
	resp, err := client.Write(context.Background(), &api.WriteRequest{
		Updates: []*api.RelationTupleUpdate{
			createTuple(tuple(aDoc("owner"), anOwner)),
			createTuple(tuple(aDoc("contributor"), anEditor)),
			createTuple(tuple(aDoc("viewer"), aViewer)),
		}})
	if err != nil {
		log.Fatalf("unable to write tuples: %s", err)
	}

	// Save the revision from the Write for future requests in order to enforce
	// that responses are at least as fresh as our last write.
	//
	// We recommend saving this from any call to Write or ContentChangeCheck,
	// and storing it alongside the object referenced in the write or check (in this case aDoc)"
	//
	// For more info see:
	// https://docs.authzed.com/authz/new-enemy
	whenPermsChanged := resp.Revision

	// Run some permission checks on the written data.
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
		testResp, err := client.Check(context.Background(), &api.CheckRequest{
			TestUserset: test.permission,
			User: &api.User{UserOneof: &api.User_Userset{
				Userset: test.user,
			}},
			AtRevision: whenPermsChanged, // Guarantee checks occur on data fresher than the write.
		})
		if err != nil {
			log.Fatalf("unable to run check request: %s", err)
		}

		hasAccess := testResp.GetMembership() == api.CheckResponse_MEMBER
		if hasAccess != test.hasAccess {
			log.Fatalf("check returned the wrong result: %v", test)
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
