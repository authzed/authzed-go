package authzed

import (
	"context"
	"fmt"
	"reflect"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/authzed/authzed-go/x/parallel"
)

const (
	initialBatchSize = 8
	growthFactor     = 2
)

// NewFilterIter returns a `FilterIter` that outputs only items from the
// provided slice that are accessible:
// - by the provided user
// - via the provided relation
// - at the optional revision
//
// The provided slice value must be `[]T` where `T` implements `Checkable` or
// `CheckableAtRevision`. If `CheckableAtRevision` is implemented, the revision
// returned by `Revision()` is used instead of the optional one provided as an
// arugment to this function.
func (c *Client) NewFilterIter(slice interface{}, user *api.User, relation string, optionalRevision *api.Zookie) FilterIter {
	return &iter{
		client:   c,
		user:     user,
		relation: relation,
		revision: optionalRevision,

		batchSize:  initialBatchSize,
		batchIndex: 0,
		unfiltered: reflect.Indirect(reflect.ValueOf(slice)),
		filtered:   nil,
	}
}

// Checkable represents any object that can be represented as an
// ObjectAndRelation.
type Checkable interface {
	AsObjectAndRelation(relation string) *api.ObjectAndRelation
}

// CheckableAtRevision represents any object that can be represented at a
// specific revision.
type CheckableAtRevision interface {
	Checkable
	Revision() *api.Zookie
}

// FilterIter represents an iterator over a list of values that have been
// filtered based on access.
type FilterIter interface {
	Next(context.Context) bool
	Err() error
	Item() interface{}
}

type iter struct {
	client   *Client
	user     *api.User
	relation string
	revision *api.Zookie

	batchSize  int
	batchIndex int
	unfiltered reflect.Value
	filtered   []interface{}
	err        error
}

func (it *iter) Next(ctx context.Context) bool {
	if it.err != nil {
		return false
	}

	if len(it.filtered) == 0 && it.batchIndex >= it.unfiltered.Len() {
		return false
	}

	for len(it.filtered) == 0 && it.batchIndex < it.unfiltered.Len() {
		batchStartIndex := it.batchIndex
		batchEndIndex := min(it.unfiltered.Len()-it.batchIndex, it.batchIndex+it.batchSize)

		it.batchIndex = batchEndIndex
		it.batchSize = it.batchSize * growthFactor

		reqs := make([]*api.CheckRequest, 0, batchEndIndex-batchStartIndex)
		for i := batchStartIndex; i < batchEndIndex; i++ {
			req, err := intoRequest(it.unfiltered.Index(i).Interface(), it.user, it.relation, it.revision)
			if err != nil {
				it.err = err
				return false
			}
			reqs = append(reqs, req)
		}

		resps, err := parallel.Check(ctx, it.client, reqs)
		if err != nil {
			it.err = err
			return false
		}

		for i, resp := range resps {
			if resp.Membership == api.CheckResponse_MEMBER {
				it.filtered = append(it.filtered, it.unfiltered.Index(i+batchStartIndex).Interface())
			}
		}
	}

	return len(it.filtered) > 0
}

func (it *iter) Item() interface{} {
	if it.err != nil {
		panic("call to Item() when FilterIter was in an errored state")
	}

	if len(it.filtered) == 0 {
		panic("call to exhausted FilterIter; use Next() before calling")
	}

	head := it.filtered[0]
	it.filtered = it.filtered[1:]
	return head
}

func (it *iter) Err() error { return it.err }

func intoRequest(i interface{}, user *api.User, relation string, rev *api.Zookie) (*api.CheckRequest, error) {
	switch x := i.(type) {
	case CheckableAtRevision:
		return &api.CheckRequest{
			TestUserset: x.AsObjectAndRelation(relation),
			User:        user,
			AtRevision:  x.Revision(),
		}, nil
	case Checkable:
		return &api.CheckRequest{
			TestUserset: x.AsObjectAndRelation(relation),
			User:        user,
			AtRevision:  rev,
		}, nil
	default:
		return nil, fmt.Errorf("type provided was not Checkable: %t", x)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
