package authzed

import (
	"testing"

	api "github.com/authzed/authzed-go/arrakisapi/api"
)

func user(namespace, object, relation string) *api.User {
	return &api.User{UserOneof: &api.User_Userset{Userset: onr(namespace, object, relation)}}
}

func onr(namespace, object, relation string) *api.ObjectAndRelation {
	return &api.ObjectAndRelation{
		Namespace: namespace,
		ObjectId:  object,
		Relation:  relation,
	}
}

func TestFlatten(t *testing.T) {
	testCases := []struct {
		name     string
		tree     *api.RelationTupleTreeNode
		expected []*api.ObjectAndRelation
	}{
		{
			"simple leaf",
			leaf(nil, user("user", "user1", "...")),
			[]*api.ObjectAndRelation{onr("user", "user1", "...")},
		},
		{
			"simple union",
			union(nil,
				leaf(nil, user("user", "user1", "...")),
				leaf(nil, user("user", "user2", "...")),
				leaf(nil, user("user", "user3", "...")),
			),
			[]*api.ObjectAndRelation{
				onr("user", "user1", "..."),
				onr("user", "user2", "..."),
				onr("user", "user3", "..."),
			},
		},
		{
			"simple intersection",
			intersection(nil,
				leaf(nil,
					user("user", "user1", "..."),
					user("user", "user2", "..."),
				),
				leaf(nil,
					user("user", "user2", "..."),
					user("user", "user3", "..."),
				),
				leaf(nil,
					user("user", "user2", "..."),
					user("user", "user4", "..."),
				),
			),
			[]*api.ObjectAndRelation{onr("user", "user2", "...")},
		},
		{
			"empty intersection",
			intersection(nil,
				leaf(nil,
					user("user", "user1", "..."),
					user("user", "user2", "..."),
				),
				leaf(nil,
					user("user", "user3", "..."),
					user("user", "user4", "..."),
				),
			),
			[]*api.ObjectAndRelation{},
		},
		{
			"simple exclusion",
			exclusion(nil,
				leaf(nil,
					user("user", "user1", "..."),
					user("user", "user2", "..."),
				),
				leaf(nil, user("user", "user2", "...")),
				leaf(nil, user("user", "user3", "...")),
			),
			[]*api.ObjectAndRelation{onr("user", "user1", "...")},
		},
		{
			"empty exclusion",
			exclusion(nil,
				leaf(nil,
					user("user", "user1", "..."),
					user("user", "user2", "..."),
				),
				leaf(nil, user("user", "user1", "...")),
				leaf(nil, user("user", "user2", "...")),
			),
			[]*api.ObjectAndRelation{},
		},
		{
			"nested union",
			union(nil,
				leaf(nil, user("user", "user1", "...")),
				union(nil, leaf(nil, user("user", "user2", "..."))),
			),
			[]*api.ObjectAndRelation{
				onr("user", "user1", "..."),
				onr("user", "user2", "..."),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var flattened userSet = make(map[string]struct{})
			flattened.add(flatten(tc.tree)...)

			for _, onr := range tc.expected {
				usr := user(onr.Namespace, onr.ObjectId, onr.Relation)
				if !flattened.contains(usr) {
					t.Fatalf("flattened tree failed to contain expected user: %s", usr)
				}
				flattened.remove(usr)
			}

			if len(flattened) != 0 {
				t.Fatalf("additional users found in flattened tree: %s", flattened)
			}
		})
	}
}
