// Package parallel implements experimental utilities for performing parallel
// client interactions with the Authzed API.
package parallel

import (
	"context"

	"google.golang.org/grpc"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/authzed/authzed-go/internal/ctxgroup"
)

// Check performs the provided list of CheckRequests in parallel and returns
// a slice of responses that are stored at their respective indices.
func Check(ctx context.Context, client api.ACLServiceClient, in []*api.CheckRequest, opts ...grpc.CallOption) ([]*api.CheckResponse, error) {
	results := make([]*api.CheckResponse, len(in))
	g := ctxgroup.WithContext(ctx)
	for i, request := range in {
		i, request := i, request // close over the value that will not change
		g.GoCtx(func(ctx context.Context) (err error) {
			results[i], err = client.Check(ctx, request, opts...)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// ContentChangeCheck performs the provided list of ContentChangeCheckRequests
// in parallel and returns a slice of responses that are stored at their
// respective indices.
func ContentChangeCheck(ctx context.Context, client api.ACLServiceClient, in []*api.ContentChangeCheckRequest, opts ...grpc.CallOption) ([]*api.CheckResponse, error) {
	results := make([]*api.CheckResponse, len(in))
	g := ctxgroup.WithContext(ctx)
	for i, request := range in {
		i, request := i, request // close over the value that will not change
		g.GoCtx(func(ctx context.Context) (err error) {
			results[i], err = client.ContentChangeCheck(ctx, request, opts...)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// AllChecksAreMember returns true if every provided CheckResponse has
// membership.
func AllChecksAreMember(resps []*api.CheckResponse) bool {
	for _, resp := range resps {
		if resp.Membership != api.CheckResponse_MEMBER {
			return false
		}
	}
	return true
}
