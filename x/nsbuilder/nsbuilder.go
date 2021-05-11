// Package nsbuilder implements a builder-pattern for defining Authzed
// Namespaces.
package nsbuilder

import (
	api "github.com/authzed/authzed-go/arrakisapi/api"
)

// Namespace creates a namespace definition with one or more defined relations.
func Namespace(name string, relations ...*api.Relation) *api.NamespaceDefinition {
	return &api.NamespaceDefinition{
		Name:     name,
		Relation: relations,
	}
}

// Relation creates a relation definition with an optional rewrite definition.
func Relation(name string, rewrite *api.UsersetRewrite) *api.Relation {
	return &api.Relation{
		Name:           name,
		UsersetRewrite: rewrite,
	}
}

// Union creates a rewrite definition that combines/considers usersets in all children.
func Union(firstChild *api.SetOperation_Child, rest ...*api.SetOperation_Child) *api.UsersetRewrite {
	return &api.UsersetRewrite{
		RewriteOperation: &api.UsersetRewrite_Union{
			Union: setOperation(firstChild, rest),
		},
	}
}

// Intersection creates a rewrite definition that returns/considers only usersets present in all children.
func Intersection(firstChild *api.SetOperation_Child, rest ...*api.SetOperation_Child) *api.UsersetRewrite {
	return &api.UsersetRewrite{
		RewriteOperation: &api.UsersetRewrite_Intersection{
			Intersection: setOperation(firstChild, rest),
		},
	}
}

// Exclusion creates a rewrite definition that starts with the usersets of the first child
// and iteratively removes usersets that appear in the remaining children.
func Exclusion(firstChild *api.SetOperation_Child, rest ...*api.SetOperation_Child) *api.UsersetRewrite {
	return &api.UsersetRewrite{
		RewriteOperation: &api.UsersetRewrite_Exclusion{
			Exclusion: setOperation(firstChild, rest),
		},
	}
}

func setOperation(firstChild *api.SetOperation_Child, rest []*api.SetOperation_Child) *api.SetOperation {
	children := append([]*api.SetOperation_Child{firstChild}, rest...)
	return &api.SetOperation{
		Child: children,
	}
}

// This creates a child for a set operation that references only direct usersets with the parent relation.
func This() *api.SetOperation_Child {
	return &api.SetOperation_Child{
		ChildType: &api.SetOperation_Child_XThis{},
	}
}

// ComputesUserset creates a child for a set operation that follows a relation on the given starting object.
func ComputedUserset(relation string) *api.SetOperation_Child {
	return &api.SetOperation_Child{
		ChildType: &api.SetOperation_Child_ComputedUserset{
			ComputedUserset: &api.ComputedUserset{
				Relation: relation,
			},
		},
	}
}

// TupleToUserset creates a child which first loads all tuples with the specific relation,
// and then unions all children on the usersets found by following a relation on those loaded
// tuples.
func TupleToUserset(tuplesetRelation, usersetRelation string) *api.SetOperation_Child {
	return &api.SetOperation_Child{
		ChildType: &api.SetOperation_Child_TupleToUserset{
			TupleToUserset: &api.TupleToUserset{
				Tupleset: &api.TupleToUserset_Tupleset{
					Relation: tuplesetRelation,
				},
				ComputedUserset: &api.ComputedUserset{
					Relation: usersetRelation,
					Object:   api.ComputedUserset_TUPLE_USERSET_OBJECT,
				},
			},
		},
	}
}
