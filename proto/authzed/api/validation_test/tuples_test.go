package validation_test

import (
	"fmt"
	"strings"
	"testing"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/stretchr/testify/require"
)

var namespaces = []struct {
	name  string
	valid bool
}{
	{"", false},
	{"...", false},
	{"foo", false},
	{"bar", false},
	{"foo1", true},
	{"bar1", true},
	{"ab", false},
	{"Foo1", false},
	{"foo_bar", true},
	{"foo_bar_", false},
	{"foo/bar", false},
	{"foo/b", false},
	{"Foo/bar", false},
	{"foo/bar/baz", false},
	{strings.Repeat("f", 3), false},
	{strings.Repeat("f", 4), true},
	{strings.Repeat("\u0394", 4), false},
	{strings.Repeat("\n", 4), false},
	{strings.Repeat("_", 4), false},
	{strings.Repeat("-", 4), false},
	{strings.Repeat("/", 4), false},
	{strings.Repeat("\\", 4), false},
	{strings.Repeat("f", 64), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", 63), strings.Repeat("f", 63)), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", 64), strings.Repeat("f", 64)), false}, // This doesn't match the code
	{fmt.Sprintf("%s/%s", strings.Repeat("f", 65), strings.Repeat("f", 64)), false},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", 64), strings.Repeat("f", 65)), false},
	{strings.Repeat("f", 65), false},
}

var objectIDs = []struct {
	name  string
	valid bool
}{
	{"a", true},
	{"1", true},
	{"a-", true},
	{"A-A-A", true},
	{"123e4567-e89b-12d3-a456-426614174000", true},
	{strings.Repeat("f", 64), true},
	{"", false},
	{"  ", false},
	{"-", false},
	{strings.Repeat("\u0394", 4), false},
	{strings.Repeat("f", 65), false},
}

type relationValidity int

const (
	alwaysInvalid relationValidity = iota
	alwaysValid
	validV0SubjectOnly
	validV1SubjectOnly
)

type relationEntry struct {
	name     string
	validity relationValidity
}

var knownGoodONR *v0.ObjectAndRelation = &v0.ObjectAndRelation{
	Namespace: "user",
	ObjectId:  "testuser",
	Relation:  "member",
}

var knownGoodObjectRef *v1.ObjectReference = &v1.ObjectReference{
	ObjectType: "user",
	ObjectId:   "testuser",
}

var knownGoodSubjectRef *v1.SubjectReference = &v1.SubjectReference{
	Object:           knownGoodObjectRef,
	OptionalRelation: "member",
}

var relations = []relationEntry{
	{"...", validV0SubjectOnly},
	{"", validV1SubjectOnly},
	{"foo", alwaysInvalid},
	{"bar", alwaysInvalid},
	{"foo1", alwaysValid},
	{"bar1", alwaysValid},
	{"ab", alwaysInvalid},
	{"Foo1", alwaysInvalid},
	{"foo_bar", alwaysValid},
	{"foo_bar_", alwaysInvalid},
	{"foo/bar", alwaysInvalid},
	{"foo/b", alwaysInvalid},
	{"Foo/bar", alwaysInvalid},
	{"foo/bar/baz", alwaysInvalid},
	{strings.Repeat("f", 3), alwaysInvalid},
	{strings.Repeat("f", 4), alwaysValid},
	{strings.Repeat("\u0394", 4), alwaysInvalid},
	{strings.Repeat("\n", 4), alwaysInvalid},
	{strings.Repeat("_", 4), alwaysInvalid},
	{strings.Repeat("-", 4), alwaysInvalid},
	{strings.Repeat("/", 4), alwaysInvalid},
	{strings.Repeat("\\", 4), alwaysInvalid},
	{strings.Repeat("f", 64), alwaysValid},
	{strings.Repeat("f", 65), alwaysInvalid},
}

func TestV0CoreObjectValidity(t *testing.T) {
	for _, ns := range namespaces {
		for _, objectID := range objectIDs {
			for _, relation := range relations {
				testName := fmt.Sprintf("%s:%s#%s", ns.name, objectID.name, relation.name)
				t.Run(testName, func(t *testing.T) {
					require := require.New(t)

					v0Valid := ns.valid && objectID.valid &&
						(relation.validity == alwaysValid ||
							relation.validity == validV0SubjectOnly)

					onr := &v0.ObjectAndRelation{
						Namespace: ns.name,
						ObjectId:  objectID.name,
						Relation:  relation.name,
					}
					err := onr.Validate()
					require.Equal(v0Valid, err == nil, "should be valid: %v %s", v0Valid, err)

					asObject := &v0.RelationTuple{
						ObjectAndRelation: onr,
						User: &v0.User{
							UserOneof: &v0.User_Userset{
								Userset: knownGoodONR,
							},
						},
					}
					err = asObject.Validate()
					require.Equal(v0Valid, err == nil, "should be valid: %v %s", v0Valid, err)

					asSubject := &v0.RelationTuple{
						ObjectAndRelation: onr,
						User: &v0.User{
							UserOneof: &v0.User_Userset{
								Userset: knownGoodONR,
							},
						},
					}
					err = asSubject.Validate()
					require.Equal(v0Valid, err == nil, "should be valid: %v %s", v0Valid, err)

					// test all of the types of tupleset filters
					justNS := &v0.RelationTupleFilter{
						Namespace: ns.name,
					}
					filterValid := ns.valid
					err = justNS.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					objectIDFilter := &v0.RelationTupleFilter{
						Namespace: ns.name,
						ObjectId:  objectID.name,
						Filters: []v0.RelationTupleFilter_Filter{
							v0.RelationTupleFilter_OBJECT_ID,
						},
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "")
					err = objectIDFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					objectRelationFilter := &v0.RelationTupleFilter{
						Namespace: ns.name,
						Relation:  relation.name,
						Filters: []v0.RelationTupleFilter_Filter{
							v0.RelationTupleFilter_RELATION,
						},
					}
					filterValid = ns.valid && (relation.validity == alwaysValid || relation.name == "")
					err = objectRelationFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					fullObjectFilter := &v0.RelationTupleFilter{
						Namespace: ns.name,
						ObjectId:  objectID.name,
						Relation:  relation.name,
						Filters: []v0.RelationTupleFilter_Filter{
							v0.RelationTupleFilter_OBJECT_ID,
							v0.RelationTupleFilter_RELATION,
						},
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "") &&
						(relation.validity == alwaysValid || relation.name == "")
					err = fullObjectFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					subjectFilter := &v0.RelationTupleFilter{
						Namespace: knownGoodObjectRef.ObjectType,
						Userset:   onr,
						Filters: []v0.RelationTupleFilter_Filter{
							v0.RelationTupleFilter_USERSET,
						},
					}
					err = subjectFilter.Validate()
					require.Equal(v0Valid, err == nil, "should be valid: %v %s", v0Valid, err)
				})
			}
		}
	}
}

func TestV1CoreObjectValidity(t *testing.T) {
	for _, ns := range namespaces {
		for _, objectID := range objectIDs {
			for _, relation := range relations {
				testName := fmt.Sprintf("%s:%s#%s", ns.name, objectID.name, relation.name)
				t.Run(testName, func(t *testing.T) {
					require := require.New(t)

					objRef := &v1.ObjectReference{
						ObjectType: ns.name,
						ObjectId:   objectID.name,
					}
					objRefValid := ns.valid && objectID.valid
					err := objRef.Validate()
					require.Equal(objRefValid, err == nil, "should be valid: %v %s", objRefValid, err)

					subRef := &v1.SubjectReference{
						Object:           objRef,
						OptionalRelation: relation.name,
					}
					subjectValid := ns.valid && objectID.valid &&
						(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
					err = subRef.Validate()
					require.Equal(subjectValid, err == nil, "should be valid: %v %s", subjectValid, err)

					asResource := &v1.Relationship{
						Resource: objRef,
						Relation: relation.name,
						Subject:  knownGoodSubjectRef,
					}
					asResourceValid := objRefValid && relation.validity == alwaysValid
					err = asResource.Validate()
					require.Equal(asResourceValid, err == nil, "should be valid: %v %s", asResourceValid, err)

					asSubject := &v1.Relationship{
						Resource: knownGoodObjectRef,
						Relation: knownGoodSubjectRef.OptionalRelation,
						Subject:  subRef,
					}
					err = asSubject.Validate()
					require.Equal(subjectValid, err == nil, "should be valid: %v %s", subjectValid, err)

					// Test all the components of a filter
					justNS := &v1.RelationshipFilter{
						ResourceType: ns.name,
					}
					filterValid := ns.valid
					err = justNS.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					objectIDFilter := &v1.RelationshipFilter{
						ResourceType:       ns.name,
						OptionalResourceId: objectID.name,
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "")
					err = objectIDFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					objectRelationFilter := &v1.RelationshipFilter{
						ResourceType:     ns.name,
						OptionalRelation: relation.name,
					}
					filterValid = ns.valid && (relation.validity == alwaysValid || relation.name == "")
					err = objectRelationFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					fullObjectFilter := &v1.RelationshipFilter{
						ResourceType:       ns.name,
						OptionalResourceId: objectID.name,
						OptionalRelation:   relation.name,
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "") &&
						(relation.validity == alwaysValid || relation.name == "")
					err = fullObjectFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					bothTypesFilter := &v1.RelationshipFilter{
						ResourceType: knownGoodObjectRef.ObjectType,
						OptionalSubjectFilter: &v1.SubjectFilter{
							SubjectType: ns.name,
						},
					}
					filterValid = ns.valid
					err = bothTypesFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					subjectIDFilter := &v1.RelationshipFilter{
						ResourceType: knownGoodObjectRef.ObjectType,
						OptionalSubjectFilter: &v1.SubjectFilter{
							SubjectType:       ns.name,
							OptionalSubjectId: objectID.name,
						},
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "")
					err = subjectIDFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					subjectRelationFilter := &v1.RelationshipFilter{
						ResourceType: knownGoodObjectRef.ObjectType,
						OptionalSubjectFilter: &v1.SubjectFilter{
							SubjectType: ns.name,
							OptionalRelation: &v1.SubjectFilter_RelationFilter{
								Relation: relation.name,
							},
						},
					}
					filterValid = ns.valid &&
						(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
					err = subjectRelationFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

					fullSubjectFilter := &v1.RelationshipFilter{
						ResourceType: knownGoodObjectRef.ObjectType,
						OptionalSubjectFilter: &v1.SubjectFilter{
							SubjectType:       ns.name,
							OptionalSubjectId: objectID.name,
							OptionalRelation: &v1.SubjectFilter_RelationFilter{
								Relation: relation.name,
							},
						},
					}
					filterValid = ns.valid && (objectID.valid || objectID.name == "") &&
						(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
					err = fullSubjectFilter.Validate()
					require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)
				})
			}
		}
	}
}
