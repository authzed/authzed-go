package authzed_test

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"testing"

	"github.com/authzed/grpcutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
)

var fullyConsistent = &v1.Consistency{Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true}}

func ExampleNewClient() {
	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		log.Fatalf("failed to load system certs: %s", err)
	}
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithBearerToken("tc_my_token_deadbeefdeadbeefdeadbeef"),
		systemCerts,
	)
	if err != nil {
		log.Fatalf("failed to connect to authzed: %s", err)
	}
	log.Println(client)
}

func randomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer)[:length], nil
}

func testClient(t *testing.T) *authzed.Client {
	t.Helper()
	token, err := randomString(12)
	require.NoError(t, err)
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken(token),
	)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Close()) })
	return client
}

func TestBasicSchema(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	schema := `
	definition document {
		relation reader: user
	}
	definition user {}
	`

	writeResponse, err := client.SchemaServiceClient.WriteSchema(ctx, &v1.WriteSchemaRequest{Schema: schema})
	require.NoError(err)
	require.NotEmpty(writeResponse.WrittenAt.String())

	readResponse, err := client.SchemaServiceClient.ReadSchema(ctx, &v1.ReadSchemaRequest{})
	require.NoError(err)
	require.Contains(readResponse.SchemaText, "definition document")
	require.Contains(readResponse.SchemaText, "definition user")
}

func TestSchemaWithCaveats(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	err := WriteTestSchema(client)
	require.NoError(err)
}

func TestCheck(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	emilia, beatrice, postOne, _, err := WriteTestTuples(client)
	require.NoError(err)

	firstResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "view",
		Subject:     emilia,
		Consistency: fullyConsistent,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, firstResponse.Permissionship)

	secondResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "write",
		Subject:     emilia,
		Consistency: fullyConsistent,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, secondResponse.Permissionship)

	thirdResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "view",
		Subject:     beatrice,
		Consistency: fullyConsistent,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, thirdResponse.Permissionship)

	fourthResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "write",
		Subject:     beatrice,
		Consistency: fullyConsistent,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION, fourthResponse.Permissionship)
}

func TestCaveatedCheck(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	_, beatrice, postOne, _, err := WriteTestTuples(client)
	require.NoError(err)

	// Likes Harry Potter
	likesContext, err := structpb.NewStruct(map[string]any{"likes": true})
	require.NoError(err)
	firstResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "view_as_fan",
		Subject:     beatrice,
		Consistency: fullyConsistent,
		Context:     likesContext,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, firstResponse.Permissionship)

	// No longer likes Harry Potter
	dislikesContext, err := structpb.NewStruct(map[string]any{"likes": false})
	require.NoError(err)
	secondResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "view_as_fan",
		Subject:     beatrice,
		Consistency: fullyConsistent,
		Context:     dislikesContext,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION, secondResponse.Permissionship)

	// Fandom is in question
	require.NoError(err)
	thirdResponse, err := client.PermissionsServiceClient.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Resource:    postOne,
		Permission:  "view_as_fan",
		Subject:     beatrice,
		Consistency: fullyConsistent,
	})
	require.NoError(err)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_CONDITIONAL_PERMISSION, thirdResponse.Permissionship)
	require.Contains(thirdResponse.PartialCaveatInfo.MissingRequiredContext, "likes")
}

func TestLookupResources(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	emilia, _, postOne, postTwo, err := WriteTestTuples(client)
	require.NoError(err)

	// NOTE: setting a page limit and then cursoring over that limit is entirely overkill
	// for this case, where we know how many results we're expecting. This is meant as an
	// example to demonstrate a real-world lookupResources usage.
	pageLimit := 50
	// Where the result buffer is where we'll concatenate the page buffers together
	// to get a final set of results
	resultBuffer := make([]string, 0)

	for {
		response, err := client.PermissionsServiceClient.LookupResources(ctx, &v1.LookupResourcesRequest{
			ResourceObjectType: "post",
			Permission:         "write",
			Subject:            emilia,
			Consistency:        fullyConsistent,
		})
		require.NoError(err)

		// The page buffer is where we'll store individual results from the stream
		pageBuffer := make([]string, 0, pageLimit)
		for {
			item, err := response.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			require.NoError(err)
			pageBuffer = append(pageBuffer, item.ResourceObjectId)
		}

		resultBuffer = append(resultBuffer, pageBuffer...)
		resultCount := len(pageBuffer)

		// If there are no results or the number of results is less than the page limit,
		// we know that we've exhausted the pages of results.
		if resultCount == 0 || resultCount < pageLimit {
			break
		}
	}
	require.Contains(resultBuffer, postOne.ObjectId)
	require.Contains(resultBuffer, postTwo.ObjectId)
	require.Len(resultBuffer, 2)
}

func TestLookupSubjects(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	emilia, beatrice, postOne, _, err := WriteTestTuples(client)
	require.NoError(err)

	// NOTE: we do a more naive approach here because the LookupSubjects API
	// doesn't support cursoring.
	resultBuffer := make([]string, 0)

	response, err := client.PermissionsServiceClient.LookupSubjects(ctx, &v1.LookupSubjectsRequest{
		SubjectObjectType: "user",
		Permission:        "view",
		Resource:          postOne,
		Consistency:       fullyConsistent,
	})
	require.NoError(err)

	for {
		item, err := response.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		require.NoError(err)
		resultBuffer = append(resultBuffer, item.Subject.SubjectObjectId)
	}

	require.Contains(resultBuffer, emilia.Object.ObjectId)
	require.Contains(resultBuffer, beatrice.Object.ObjectId)
	require.Len(resultBuffer, 2)
}

func TestCheckBulkPermissions(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	emilia, _, postOne, _, err := WriteTestTuples(client)
	require.NoError(err)

	response, err := client.PermissionsServiceClient.CheckBulkPermissions(ctx, &v1.CheckBulkPermissionsRequest{
		Consistency: fullyConsistent,
		Items: []*v1.CheckBulkPermissionsRequestItem{
			{
				Resource:   postOne,
				Permission: "view",
				Subject:    emilia,
			},
			{
				Resource:   postOne,
				Permission: "write",
				Subject:    emilia,
			},
		},
	})
	require.NoError(err)

	require.Len(response.Pairs, 2)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, response.Pairs[0].GetItem().Permissionship)
	require.Equal(v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, response.Pairs[1].GetItem().Permissionship)
}

func TestBulkExportImport(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client := testClient(t)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err := WriteTestSchema(client)
	require.NoError(err)
	_, _, _, _, err = WriteTestTuples(client)
	require.NoError(err)

	// Validate export
	exportResponse, err := client.PermissionsServiceClient.ExportBulkRelationships(ctx, &v1.ExportBulkRelationshipsRequest{
		Consistency: fullyConsistent,
	})
	require.NoError(err)

	exportResults := make([]*v1.Relationship, 0)
	for {
		item, err := exportResponse.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		require.NoError(err)
		exportResults = append(exportResults, item.Relationships...)
	}

	require.Len(exportResults, 4)

	// Note that this has a different preshared key
	// Validate import
	emptyClient := testClient(t)
	err = WriteTestSchema(emptyClient)
	require.NoError(err)

	stream, err := emptyClient.PermissionsServiceClient.ImportBulkRelationships(ctx)
	require.NoError(err)
	err = stream.Send(&v1.ImportBulkRelationshipsRequest{
		Relationships: exportResults,
	})
	require.NoError(err)
	importResponse, err := stream.CloseAndRecv()
	require.NoError(err)
	require.Equal(uint64(4), importResponse.NumLoaded)

	// Validate that things were loaded
	exportAfterImportResponse, err := emptyClient.PermissionsServiceClient.ExportBulkRelationships(ctx, &v1.ExportBulkRelationshipsRequest{
		Consistency: fullyConsistent,
	})
	require.NoError(err)

	exportAfterImportResults := make([]*v1.Relationship, 0)
	for {
		item, err := exportAfterImportResponse.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		require.NoError(err)
		exportAfterImportResults = append(exportAfterImportResults, item.Relationships...)
	}

	require.Len(exportAfterImportResults, 4)
}

func WriteTestTuples(client *authzed.Client) (emilia *v1.SubjectReference, beatrice *v1.SubjectReference, postOne *v1.ObjectReference, postTwo *v1.ObjectReference, err error) {
	emilia = &v1.SubjectReference{Object: &v1.ObjectReference{ObjectType: "user", ObjectId: "emilia"}}
	beatrice = &v1.SubjectReference{Object: &v1.ObjectReference{ObjectType: "user", ObjectId: "beatrice"}}
	postOne = &v1.ObjectReference{ObjectType: "post", ObjectId: "post-one"}
	postTwo = &v1.ObjectReference{ObjectType: "post", ObjectId: "post-two"}
	_, err = client.PermissionsServiceClient.WriteRelationships(context.Background(), &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: postOne,
					Relation: "writer",
					Subject:  emilia,
				},
			},
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: postTwo,
					Relation: "writer",
					Subject:  emilia,
				},
			},
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: postOne,
					Relation: "reader",
					Subject:  beatrice,
				},
			},
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource:       postOne,
					Relation:       "caveated_reader",
					Subject:        beatrice,
					OptionalCaveat: &v1.ContextualizedCaveat{CaveatName: "likes_harry_potter"},
				},
			},
		},
	})
	return
}

func WriteTestSchema(client *authzed.Client) error {
	schema := `
            caveat likes_harry_potter(likes bool) {
              likes == true
            }

            definition post {
                relation writer: user
                relation reader: user
                relation caveated_reader: user with likes_harry_potter

                permission write = writer
                permission view = reader + writer
                permission view_as_fan = caveated_reader + writer
            }
            definition user {}
        `
	_, err := client.SchemaServiceClient.WriteSchema(context.Background(), &v1.WriteSchemaRequest{Schema: schema})
	return err
}
