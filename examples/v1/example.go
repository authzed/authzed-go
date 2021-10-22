package main

import (
	"context"
	"log"

	"github.com/authzed/grpcutil"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
)

const (
	documentNS = "yourtenant/document"
	userNS     = "yourtenant/user"

	schema = `
	definition yourtenant/user {}

	definition yourtenant/document {
		relation viewer: yourtenant/user
		relation contributor: yourtenant/user
		relation owner: yourtenant/user

		permission read = viewer + contributor + owner
		permission write = contributor + owner
		permission delete = owner
	}
	`
)

func main() {
	// Create an Authzed client.
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithInsecureBearerToken("t_your_token_here"),
		grpcutil.WithSystemCerts(grpcutil.VerifyCA),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	// Uncomment this block to run against a local SpiceDB.
	// client, err = authzed.NewClient(
	// 	"localhost:50051",
	// 	grpcutil.WithInsecureBearerToken("testtesttesttest"),
	// 	grpc.WithInsecure(),
	// )
	// if err != nil {
	// 	log.Fatalf("unable to initialize client: %s", err)
	// }

	// Write the schema to the permissions system
	_, err = client.WriteSchema(context.Background(), &v1.WriteSchemaRequest{
		Schema: schema,
	})
	if err != nil {
		log.Fatalf("unable to write schema: %s", err)
	}

	// Create some objects that will be protected by Authzed.
	aDoc := object(documentNS, "doc1")
	anOwner := subject(userNS, "theowner")
	anEditor := subject(userNS, "userwhocanedit")
	aViewer := subject(userNS, "viewonlyuser")

	// Create some relationships that represent roles granted between users and objects.
	resp, err := client.WriteRelationships(context.Background(), &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			createRelationship(relationship(aDoc, "owner", anOwner)),
			createRelationship(relationship(aDoc, "contributor", anEditor)),
			createRelationship(relationship(aDoc, "viewer", aViewer)),
		},
	})
	if err != nil {
		log.Fatalf("unable to write tuples: %s", err)
	}

	// Save the ZedToken from the Write for future requests in order to enforce
	// that responses are at least as fresh as our last write.
	//
	// We recommend saving this from any call to WriteRelationships and storing it
	// alongside the object referenced in the write or check (in this case aDoc)"
	//
	// For more info see:
	// https://github.com/authzed/spicedb/blob/main/docs/zedtokens-and-zookies.md
	whenPermsChanged := resp.WrittenAt

	// Run some permission checks on the written data.
	aNobody := subject(userNS, "randomnobody")
	expected := []struct {
		resource   *v1.ObjectReference
		permission string
		subject    *v1.SubjectReference
		hasAccess  bool
	}{
		{aDoc, "read", anOwner, true},
		{aDoc, "write", anOwner, true},
		{aDoc, "delete", anOwner, true},
		{aDoc, "read", anEditor, true},
		{aDoc, "write", anEditor, true},
		{aDoc, "delete", anEditor, false},
		{aDoc, "read", aViewer, true},
		{aDoc, "write", aViewer, false},
		{aDoc, "delete", aViewer, false},
		{aDoc, "read", aNobody, false},
		{aDoc, "write", aNobody, false},
		{aDoc, "delete", aNobody, false},
	}

	for _, test := range expected {
		testResp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
			// Guarantee checks occur on data fresher than the write.
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_AtLeastAsFresh{
					AtLeastAsFresh: whenPermsChanged,
				},
			},
			Resource:   test.resource,
			Permission: test.permission,
			Subject:    test.subject,
		})
		if err != nil {
			log.Fatalf("unable to run check request: %s", err)
		}

		hasAccess := testResp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION
		if hasAccess != test.hasAccess {
			log.Fatalf("check returned the wrong result: %#v", test)
		}
	}
}

func object(namespace, objectID string) *v1.ObjectReference {
	return &v1.ObjectReference{
		ObjectType: namespace,
		ObjectId:   objectID,
	}
}

func subject(namespace, objectID string) *v1.SubjectReference {
	return &v1.SubjectReference{
		Object: object(namespace, objectID),
	}
}

func relationship(resource *v1.ObjectReference, relation string, subject *v1.SubjectReference) *v1.Relationship {
	return &v1.Relationship{
		Resource: resource,
		Relation: relation,
		Subject:  subject,
	}
}

func createRelationship(relationship *v1.Relationship) *v1.RelationshipUpdate {
	return &v1.RelationshipUpdate{
		Operation:    v1.RelationshipUpdate_OPERATION_CREATE,
		Relationship: relationship,
	}
}
