package authzed

import (
	"fmt"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

// FlattenExpand reduces an ExpandResponse into the slice of Users present in
// the expansion.
//
// Notably, this removes the context of which relations caused users to be
// included in the expansion, but often you only need to know which users are
// present.
func FlattenExpand(resp *v0.ExpandResponse) ([]*v0.User, error) { return flatten(resp.TreeNode), nil }

func flatten(node *v0.RelationTupleTreeNode) []*v0.User {
	switch typed := node.NodeType.(type) {
	case *v0.RelationTupleTreeNode_IntermediateNode:
		switch typed.IntermediateNode.Operation {
		case v0.SetOperationUserset_UNION:
			return flattenUnion(typed.IntermediateNode.ChildNodes)
		case v0.SetOperationUserset_INTERSECTION:
			return flattenIntersection(typed.IntermediateNode.ChildNodes)
		case v0.SetOperationUserset_EXCLUSION:
			return flattenExclusion(typed.IntermediateNode.ChildNodes)
		}
	case *v0.RelationTupleTreeNode_LeafNode:
		users := newUserSet()
		users.add(typed.LeafNode.Users...)
		return users.toSlice()
	}
	return nil
}

func flattenUnion(children []*v0.RelationTupleTreeNode) []*v0.User {
	users := newUserSet()
	for _, child := range children {
		users.add(flatten(child)...)
	}
	return users.toSlice()
}

func flattenIntersection(children []*v0.RelationTupleTreeNode) []*v0.User {
	firstChildChildren := flatten(children[0])

	if len(children) == 1 {
		return firstChildChildren
	}

	inOthers := newUserSet()
	inOthers.add(flattenIntersection(children[1:])...)

	maxChildren := len(firstChildChildren)
	if len(inOthers) < maxChildren {
		maxChildren = len(inOthers)
	}

	toReturn := make([]*v0.User, 0, maxChildren)
	for _, child := range firstChildChildren {
		if inOthers.contains(child) {
			toReturn = append(toReturn, child)
		}
	}

	return toReturn
}

func flattenExclusion(children []*v0.RelationTupleTreeNode) []*v0.User {
	firstChildChildren := flatten(children[0])

	if len(children) == 1 || len(firstChildChildren) == 0 {
		return firstChildChildren
	}

	users := newUserSet()
	users.add(firstChildChildren...)
	for _, child := range children[1:] {
		users.remove(flatten(child)...)
	}

	return users.toSlice()
}

func leaf(start *v0.ObjectAndRelation, children ...*v0.User) *v0.RelationTupleTreeNode {
	return &v0.RelationTupleTreeNode{
		NodeType: &v0.RelationTupleTreeNode_LeafNode{
			LeafNode: &v0.DirectUserset{
				Users: children,
			},
		},
		Expanded: start,
	}
}

func setResult(
	op v0.SetOperationUserset_Operation,
	start *v0.ObjectAndRelation,
	children []*v0.RelationTupleTreeNode,
) *v0.RelationTupleTreeNode {
	return &v0.RelationTupleTreeNode{
		NodeType: &v0.RelationTupleTreeNode_IntermediateNode{
			IntermediateNode: &v0.SetOperationUserset{
				Operation:  op,
				ChildNodes: children,
			},
		},
		Expanded: start,
	}
}

func union(start *v0.ObjectAndRelation, children ...*v0.RelationTupleTreeNode) *v0.RelationTupleTreeNode {
	return setResult(v0.SetOperationUserset_UNION, start, children)
}

func intersection(start *v0.ObjectAndRelation, children ...*v0.RelationTupleTreeNode) *v0.RelationTupleTreeNode {
	return setResult(v0.SetOperationUserset_INTERSECTION, start, children)
}

func exclusion(start *v0.ObjectAndRelation, children ...*v0.RelationTupleTreeNode) *v0.RelationTupleTreeNode {
	return setResult(v0.SetOperationUserset_EXCLUSION, start, children)
}

type userSet map[string]struct{}

func newUserSet() userSet {
	return make(map[string]struct{})
}

func (us userSet) add(users ...*v0.User) {
	for _, usr := range users {
		us[toKey(usr)] = struct{}{}
	}
}

func (us userSet) contains(usr *v0.User) bool {
	_, ok := us[toKey(usr)]
	return ok
}

func (us userSet) remove(users ...*v0.User) {
	for _, usr := range users {
		delete(us, toKey(usr))
	}
}

func (us userSet) toSlice() []*v0.User {
	users := make([]*v0.User, 0, len(us))
	for key := range us {
		users = append(users, fromKey(key))
	}
	return users
}

func toKey(usr *v0.User) string {
	return fmt.Sprintf("%s %s %s", usr.GetUserset().Namespace, usr.GetUserset().ObjectId, usr.GetUserset().Relation)
}

func fromKey(key string) *v0.User {
	userset := &v0.ObjectAndRelation{}
	fmt.Sscanf(key, "%s %s %s", &userset.Namespace, &userset.ObjectId, &userset.Relation)
	return &v0.User{
		UserOneof: &v0.User_Userset{Userset: userset},
	}
}
