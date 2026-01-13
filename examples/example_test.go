//go:build examples
// +build examples

// Package examples contains runnable examples demonstrating usage of the authzed-go client library.
//
// These examples require a running SpiceDB instance and are excluded from normal test runs.
// You can start SpiceDB with:
//
//	docker run --rm -p 50051:50051 authzed/spicedb serve --grpc-preshared-key "somerandomkeyhere" --datastore-engine memory
//
// Then run the local examples with the build tag:
//
//	go test -v -tags=examples -run 'Example_connectToSpiceDB|Example_write|Example_check|Example_lookup|Example_read|Example_delete|Example_caveat|Example_watch' ./examples/...
//
// Note: Example_connectToAuthzed requires a valid Authzed API token and connects to the hosted service.
package examples

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
)

// Example_connectToAuthzed demonstrates how to connect to Authzed's hosted service.
func Example_connectToAuthzed() {
	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		log.Fatalf("failed to load system certs: %s", err)
	}

	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		systemCerts,
		grpcutil.WithBearerToken("tc_my_token_deadbeefdeadbeefdeadbeef"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	fmt.Println("Connected to Authzed")
}

// Example_connectToSpiceDB demonstrates how to connect to a local SpiceDB instance
// without TLS (for development/testing).
func Example_connectToSpiceDB() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	fmt.Println("Connected to SpiceDB")
}

// Example_writeSchema demonstrates how to write a permission schema to SpiceDB.
func Example_writeSchema() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Define a schema with users and documents
	schema := `
		definition user {}

		definition document {
			relation owner: user
			relation editor: user
			relation viewer: user

			permission edit = owner + editor
			permission view = owner + editor + viewer
		}
	`

	resp, err := client.WriteSchema(context.Background(), &v1.WriteSchemaRequest{
		Schema: schema,
	})
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}

	fmt.Printf("Schema written at: %s\n", resp.WrittenAt.Token)
}

// Example_writeRelationships demonstrates how to create relationships between objects.
func Example_writeRelationships() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Create relationships: alice is owner of doc1, bob is viewer of doc1
	resp, err := client.WriteRelationships(context.Background(), &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: &v1.ObjectReference{
						ObjectType: "document",
						ObjectId:   "doc1",
					},
					Relation: "owner",
					Subject: &v1.SubjectReference{
						Object: &v1.ObjectReference{
							ObjectType: "user",
							ObjectId:   "alice",
						},
					},
				},
			},
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: &v1.ObjectReference{
						ObjectType: "document",
						ObjectId:   "doc1",
					},
					Relation: "viewer",
					Subject: &v1.SubjectReference{
						Object: &v1.ObjectReference{
							ObjectType: "user",
							ObjectId:   "bob",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to write relationships: %s", err)
	}

	fmt.Printf("Relationships written at: %s\n", resp.WrittenAt.Token)
}

// Example_checkPermission demonstrates how to check if a subject has a permission on a resource.
func Example_checkPermission() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Check if alice can edit doc1
	resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource: &v1.ObjectReference{
			ObjectType: "document",
			ObjectId:   "doc1",
		},
		Permission: "edit",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "alice",
			},
		},
		// Use full consistency for accurate results (slower)
		// Omit for eventual consistency (faster, may be stale)
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	switch resp.Permissionship {
	case v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION:
		fmt.Println("alice CAN edit doc1")
	case v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION:
		fmt.Println("alice CANNOT edit doc1")
	case v1.CheckPermissionResponse_PERMISSIONSHIP_CONDITIONAL_PERMISSION:
		fmt.Println("alice's permission is CONDITIONAL")
	}
}

// Example_checkBulkPermissions demonstrates how to check multiple permissions in a single request.
func Example_checkBulkPermissions() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Check multiple permissions at once
	resp, err := client.CheckBulkPermissions(context.Background(), &v1.CheckBulkPermissionsRequest{
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
		Items: []*v1.CheckBulkPermissionsRequestItem{
			{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "doc1",
				},
				Permission: "view",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "alice",
					},
				},
			},
			{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "doc1",
				},
				Permission: "edit",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "bob",
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to check bulk permissions: %s", err)
	}

	for i, pair := range resp.Pairs {
		if err := pair.GetError(); err != nil {
			fmt.Printf("Check %d: error - %v\n", i+1, err)
			continue
		}
		item := pair.GetItem()
		if item != nil {
			fmt.Printf("Check %d: %v\n", i+1, item.Permissionship)
		}
	}
}

// Example_lookupResources demonstrates how to find all resources a subject has access to.
// For streaming calls, use a context with timeout or cancellation to avoid hanging indefinitely.
func Example_lookupResources() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Use a context with timeout for streaming calls to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all documents alice can view
	stream, err := client.LookupResources(ctx, &v1.LookupResourcesRequest{
		ResourceObjectType: "document",
		Permission:         "view",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "alice",
			},
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to lookup resources: %s", err)
	}

	fmt.Println("Documents alice can view:")
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive: %s", err)
		}
		fmt.Printf("  - %s\n", resp.ResourceObjectId)
	}
}

// Example_lookupSubjects demonstrates how to find all subjects with access to a resource.
// For streaming calls, use a context with timeout or cancellation to avoid hanging indefinitely.
func Example_lookupSubjects() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Use a context with timeout for streaming calls to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all users who can view doc1
	stream, err := client.LookupSubjects(ctx, &v1.LookupSubjectsRequest{
		Resource: &v1.ObjectReference{
			ObjectType: "document",
			ObjectId:   "doc1",
		},
		Permission:        "view",
		SubjectObjectType: "user",
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to lookup subjects: %s", err)
	}

	fmt.Println("Users who can view doc1:")
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive: %s", err)
		}
		fmt.Printf("  - %s\n", resp.Subject.SubjectObjectId)
	}
}

// Example_readRelationships demonstrates how to read existing relationships.
// For streaming calls, use a context with timeout or cancellation to avoid hanging indefinitely.
func Example_readRelationships() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Use a context with timeout for streaming calls to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Read all relationships for doc1
	stream, err := client.ReadRelationships(ctx, &v1.ReadRelationshipsRequest{
		RelationshipFilter: &v1.RelationshipFilter{
			ResourceType:       "document",
			OptionalResourceId: "doc1",
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to read relationships: %s", err)
	}

	fmt.Println("Relationships for doc1:")
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive: %s", err)
		}
		rel := resp.Relationship
		fmt.Printf("  %s:%s#%s@%s:%s\n",
			rel.Resource.ObjectType,
			rel.Resource.ObjectId,
			rel.Relation,
			rel.Subject.Object.ObjectType,
			rel.Subject.Object.ObjectId,
		)
	}
}

// Example_deleteRelationships demonstrates how to delete relationships.
func Example_deleteRelationships() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Delete bob's viewer relationship on doc1
	resp, err := client.DeleteRelationships(context.Background(), &v1.DeleteRelationshipsRequest{
		RelationshipFilter: &v1.RelationshipFilter{
			ResourceType:       "document",
			OptionalResourceId: "doc1",
			OptionalRelation:   "viewer",
			OptionalSubjectFilter: &v1.SubjectFilter{
				SubjectType:       "user",
				OptionalSubjectId: "bob",
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to delete relationships: %s", err)
	}

	fmt.Printf("Deleted at: %s\n", resp.DeletedAt.Token)
}

// Example_caveatedRelationship demonstrates how to use caveats (conditional permissions).
func Example_caveatedRelationship() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// First, write a schema with a caveat
	schema := `
		caveat is_weekday(day_of_week string) {
			day_of_week != "saturday" && day_of_week != "sunday"
		}

		definition user {}

		definition document {
			relation viewer: user
			relation weekday_viewer: user with is_weekday

			permission view = viewer + weekday_viewer
		}
	`
	_, err = client.WriteSchema(context.Background(), &v1.WriteSchemaRequest{Schema: schema})
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}

	// Create a caveated relationship
	_, err = client.WriteRelationships(context.Background(), &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: &v1.ObjectReference{
						ObjectType: "document",
						ObjectId:   "doc1",
					},
					Relation: "weekday_viewer",
					Subject: &v1.SubjectReference{
						Object: &v1.ObjectReference{
							ObjectType: "user",
							ObjectId:   "charlie",
						},
					},
					OptionalCaveat: &v1.ContextualizedCaveat{
						CaveatName: "is_weekday",
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to write relationship: %s", err)
	}

	// Check permission with context
	weekdayContext, err := structpb.NewStruct(map[string]any{
		"day_of_week": "monday",
	})
	if err != nil {
		log.Fatalf("failed to create context struct: %s", err)
	}

	resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource: &v1.ObjectReference{
			ObjectType: "document",
			ObjectId:   "doc1",
		},
		Permission: "view",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "charlie",
			},
		},
		Context: weekdayContext,
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		fmt.Println("charlie CAN view doc1 on monday")
	}

	// Check permission without context (will be conditional)
	respNoContext, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource: &v1.ObjectReference{
			ObjectType: "document",
			ObjectId:   "doc1",
		},
		Permission: "view",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "charlie",
			},
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	if respNoContext.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_CONDITIONAL_PERMISSION {
		fmt.Printf("charlie's permission is CONDITIONAL, missing context: %v\n",
			respNoContext.PartialCaveatInfo.MissingRequiredContext)
	}
}

// Example_watchForChanges demonstrates how to watch for relationship changes.
// Watch is a long-running stream - use context cancellation to stop it gracefully.
func Example_watchForChanges() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer client.Close()

	// Use a cancellable context - cancel when you want to stop watching
	// In production, tie this to application shutdown signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Always clean up

	// Start watching for changes (requires a ZedToken to start from)
	// In production, you'd store and use the last seen token
	stream, err := client.Watch(ctx, &v1.WatchRequest{
		OptionalObjectTypes: []string{"document"},
	})
	if err != nil {
		log.Fatalf("failed to start watch: %s", err)
	}

	fmt.Println("Watching for document changes...")
	// In a real application, you'd process these in a loop:
	// for {
	//     resp, err := stream.Recv()
	//     if err != nil {
	//         if errors.Is(err, context.Canceled) {
	//             break // Normal shutdown
	//         }
	//         log.Fatalf("watch error: %s", err)
	//     }
	//     // Process resp.Updates...
	// }
	_ = stream
}
