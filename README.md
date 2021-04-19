# authzed-go

The official Go client library for Authzed.

## Example

This example demonstrates initializing a client and making a [Check request] to an existing [Namespace].

A more full example can be found in the [examples directory].

[Check request]:https://docs.authzed.com/concept/check
[Namespace]: https://docs.authzed.com/concept/namespaces
[examples directory]: examples

```go
package main

import (
	"context"
	"log"

	"github.com/authzed/authzed-go"
	api "github.com/authzed/authzed-go/arrakisapi/api"
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
	resp, err := client.Check(context.Background(), &api.CheckRequest{
		TestUserset: &api.ObjectAndRelation{
			Namespace: "mynoteapp/note", // Object Type
			ObjectId:  "47",             // Unique id for a object being accessed
			Relation:  "reader"          // Relationship required for access
		},
		User: &api.User{UserOneof: &api.User_Userset{Userset: &api.ObjectAndRelation{
			Namespace: "mynoteapp/user", // User Type
			ObjectId: "26",              // Unique id for a user accessing the object
			Relation: "...",
		}}},
	})
	if err != nil {
		log.Fatalf("unable to run check request: %s", err)
	}

	fmt.Println(resp.GetMembership() == api.CheckResponse_MEMBER)
	// Outputs:
	// true
}
```
