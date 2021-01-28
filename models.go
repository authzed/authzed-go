package authzed

import (
	"strings"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"github.com/jzelinskie/stringz"
)

func NewNamespace(tenant, namespace string, relations ...string) func(string) Namespace {
	return func(objectId string) Namespace {
		return Namespace{
			namespace: strings.Join([]string{tenant, namespace}, "/"),
			objectId:  objectId,
			relations: relations,
		}
	}
}

type Namespace struct {
	namespace string
	objectId  string
	relations []string
}

type PartialTuple struct {
	terminal Terminal
	relation string
}

type Tuple struct {
	namespace Namespace
	relation  string
	terminal  Terminal
}

func (t Tuple) asProto() *api.RelationTuple {
	return &api.RelationTuple{
		ObjectAndRelation: &api.ObjectAndRelation{
			Namespace: t.namespace.namespace,
			ObjectId:  t.namespace.objectId,
			Relation:  t.relation,
		},
		User: &api.User{
			UserOneof: &api.User_Userset{
				Userset: &t.terminal.ObjectAndRelation,
			},
		},
	}
}

func (p PartialTuple) Of(n Namespace) *api.RelationTuple {
	if !stringz.SliceContains(n.relations, p.relation) {
		panic("unexpected relation for namespace: " + p.relation)
	}

	return &api.RelationTuple{
		ObjectAndRelation: &api.ObjectAndRelation{
			Namespace: n.namespace,
			ObjectId:  n.objectId,
			Relation:  p.relation,
		},
		User: &api.User{
			UserOneof: &api.User_Userset{
				Userset: &p.terminal.ObjectAndRelation,
			},
		},
	}
}

type Terminal struct {
	api.ObjectAndRelation
}

func (t Terminal) Is(relation string) PartialTuple {
	return PartialTuple{
		terminal: t,
		relation: relation,
	}
}

func NewTerminal(tenant, namespace string) func(string) Terminal {
	return func(objectId string) Terminal {
		return Terminal{
			api.ObjectAndRelation{
				Namespace: strings.Join([]string{tenant, namespace}, "/"),
				ObjectId:  objectId,
				Relation:  "...",
			},
		}
	}
}
