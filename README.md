# authzed-go

The official Go client library for Authzed.
This repository is a collection of Go packages specialized for each versions of the Authzed API.

## Example (v0 API)

This example demonstrates initializing a client and making a [Check request] to an existing [Namespace].

A more full example can be found in [v0/examples].

[Check request]:https://docs.authzed.com/concept/check
[Namespace]: https://docs.authzed.com/concept/namespaces
[v0/examples]: /v0/examples

```go
package main

import (
	"context"
	"log"

	"github.com/authzed/authzed-go/v0"
	"github.com/authzed/authzed-go/proto/authzed/api/v0"
)

func main() {
	// Create an Authzed client
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		authzed.Token("t_your_token_here_1234567deadbeef"),
		authzed.SystemCerts(authzed.VerifyCA),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	// Check if User #26 has read access to note #47
	resp, err := client.Check(context.Background(), &v0.CheckRequest{
		TestUserset: &v0.ObjectAndRelation{
			Namespace: "mynoteapp/note", // Object Type
			ObjectId:  "47",             // Unique id for a object being accessed
			Relation:  "reader"          // Relationship required for access
		},
		User: &v0.User{UserOneof: &v0.User_Userset{Userset: &v0.ObjectAndRelation{
			Namespace: "mynoteapp/user", // User Type
			ObjectId: "26",              // Unique id for a user accessing the object
			Relation: "...",
		}}},
	})
	if err != nil {
		log.Fatalf("unable to run check request: %s", err)
	}

	fmt.Println(resp.GetMembership() == v0.CheckResponse_MEMBER)
	// Outputs:
	// true
}
```
