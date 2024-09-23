# Official SpiceDB Go Client

[![GoDoc](https://godoc.org/github.com/authzed/authzed-go?status.svg)](https://godoc.org/github.com/authzed/authzed-go)
[![Docs](https://img.shields.io/badge/docs-authzed.com-%234B4B6C "Authzed Documentation")](https://authzed.com/docs)
[![YouTube](https://img.shields.io/youtube/channel/views/UCFeSgZf0rPqQteiTQNGgTPg?color=%23F40203&logo=youtube&style=flat-square&label=YouTube "Authzed YouTube Channel")](https://www.youtube.com/channel/UCFeSgZf0rPqQteiTQNGgTPg)
[![Discord Server](https://img.shields.io/discord/844600078504951838?color=7289da&logo=discord "Discord Server")](https://authzed.com/discord)
[![Twitter](https://img.shields.io/badge/twitter-%40authzed-1D8EEE?logo=twitter "@authzed on Twitter")](https://twitter.com/authzed)

This repository houses the official Go client library for SpiceDB and Authzed services.

[SpiceDB] is an open source, [Google Zanzibar]-inspired, database system for creating and managing security-critical application permissions.

Developers create a schema that models their permissions requirements and use any of the official or community maintained [client libraries] to apply the schema to the database, insert data into the database, and query the data to efficiently check permissions in their applications.

[SpiceDB]: https://github.com/authzed/spicedb
[Google Zanzibar]: https://authzed.com/blog/what-is-zanzibar/
[client libraries]: https://github.com/authzed/awesome-spicedb#clients

Supported client API versions:
- [v1](https://buf.build/authzed/api/docs/main/authzed.api.v1)
- [v1alpha1](https://buf.build/authzed/api/docs/main/authzed.api.v1alpha1)

Have questions? Ask in our [Discord].

Looking to contribute? See [CONTRIBUTING.md].

You can find issues by priority: [Urgent], [High], [Medium], [Low], [Maybe].
There are also [good first issues].

[Discord]: https://authzed.com/discord
[CONTRIBUTING.md]: https://github.com/authzed/authzed-go/blob/main/CONTRIBUTING.md
[Urgent]: https://github.com/authzed/authzed-go/labels/priority%2F0%20urgent
[High]: https://github.com/authzed/authzed-go/labels/priority%2F1%20high
[Medium]: https://github.com/authzed/authzed-go/labels/priority%2F2%20medium
[Low]: https://github.com/authzed/authzed-go/labels/priority%2F3%20low
[Maybe]: https://github.com/authzed/authzed-go/labels/priority%2F4%20maybe
[good first issues]: https://github.com/authzed/authzed-go/labels/hint%2Fgood%20first%20issue

## Getting Started

We highly recommend following the **[Protecting Your First App]** guide to learn the latest best practice to integrate an application with Authzed.

[Protecting Your First App]: https://docs.authzed.com/guides/first-app

### Installation

If you're using a modern version of [Go], run the following commands to add dependencies to your project:

```sh
go get github.com/authzed/authzed-go
go get github.com/authzed/grpcutil
```

[grpcutil] is not _strictly_ required, but greatly reduces the boilerplate required to create a client in the general case.

[Go]: https://golang.org/dl/
[grpcutil]: https://github.com/authzed/grpcutil

### Initializing a client

The [`NewClient()`] constructor is the recommended method for creating a client.

Because this library is using [gRPC] under the hood, you are free to leverage the wealth of functionality provided via [DialOptions].

In order to successfully connect, you will have to provide a [Bearer Token] with your own API Token from the [Authzed dashboard] in place of `t_your_token_here_1234567deadbeef` in the following example:

[`NewClient()`]: https://pkg.go.dev/github.com/authzed/authzed-go/v1#NewClient
[Bearer Token]: https://datatracker.ietf.org/doc/html/rfc6750#section-2.1
[Authzed Dashboard]: https://app.authzed.com
[gRPC]: https://grpc.io
[DialOptions]: https://pkg.go.dev/google.golang.org/grpc?utm_source=godoc#DialOption

```go
import (
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

...
systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
if err != nil {
	log.Fatalf("unable to load system CA certificates: %s", err)
}

client, err := authzed.NewClient(
	"grpc.authzed.com:443",
	systemCerts,
	grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
)
if err != nil {
	log.Fatalf("unable to initialize client: %s", err)
}
```

### Performing an API call

Requests and response types are located in a package under `proto/` respective to their API version.

Because of the verbosity of these types, we recommend writing your own functions/methods to create these types from your existing application's models.

```go
package main

import (
	"context"
	"log"

	"github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

func main() {
	emilia := &v1.SubjectReference{Object: &v1.ObjectReference{
		ObjectType: "blog/user",
		ObjectId:   "emilia",
	}}

	firstPost := &v1.ObjectReference{
		ObjectType: "blog/post",
		ObjectId:   "1",
	}

	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		grpcutil.WithSystemCerts(grpcutil.VerifyCA),
		grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "read",
		Subject:    emilia,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		log.Println("allowed!")
	}
}
```

### Insecure Credentials
For contexts that don't require TLS, such as a development environment or integration
tests, it's possible to set up a client that does not use TLS:

```go
import (
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/authzed/authzed-go/v1"
)

client, err := authzed.NewClient(
    "localhost:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpcutil.WithInsecureBearerToken("some token"),
)
```
