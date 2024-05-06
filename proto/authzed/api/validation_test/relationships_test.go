package validation_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

const (
	minNamespaceLength = 3
	minRelationLength  = 3

	maxCaveatName      = 128
	maxTenantLength    = 63
	maxNamespaceLength = 64
	maxObjectIDLength  = 1024
	maxRelationLength  = 64
)

var namespaces = []struct {
	name  string
	valid bool
}{
	{"", false},
	{"...", false},
	{"foo", true},
	{"bar", true},
	{"foo1", true},
	{"bar1", true},
	{"ab", false},
	{"Foo1", false},
	{"foo_bar", true},
	{"foo_bar_", false},
	{"foo/bar", true},
	{"foo/b", false},
	{"Foo/bar", false},
	{"foo/bar/baz", true},
	{strings.Repeat("f", minNamespaceLength-1), false},
	{strings.Repeat("f", minNamespaceLength), true},
	{strings.Repeat("\u0394", minNamespaceLength), false},
	{strings.Repeat("\n", minNamespaceLength), false},
	{strings.Repeat("_", minNamespaceLength), false},
	{strings.Repeat("-", minNamespaceLength), false},
	{strings.Repeat("/", minNamespaceLength), false},
	{strings.Repeat("\\", minNamespaceLength), false},
	{strings.Repeat("f", maxNamespaceLength), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength), strings.Repeat("f", maxNamespaceLength)), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength+1), strings.Repeat("f", maxNamespaceLength)), false},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength), strings.Repeat("f", maxNamespaceLength+1)), false},
	{strings.Repeat("f", maxNamespaceLength+1), false},
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
	{strings.Repeat("f", maxObjectIDLength), true},
	{"", false},
	{"  ", false},
	{"-", true},
	{strings.Repeat("\u0394", 4), false},
	{strings.Repeat("f", maxObjectIDLength+1), false},
	{"a@b.com", false},
	{"test@example.com", false},
	{"test*", false},
	{"*test", false},
	{"t*t", false},
	{"test@example.*", false},
	{"authn|someauthnvalue", true},
	{"authn|a@b.com", false},
	{"authn|", true},
	{"authn|a-b-c-d-e", true},
	{"authn|a_b", true},
	{"--=base64YWZzZGZh-ZHNmZHPwn5iK8J+YivC/fmIrwn5iK==", true},
}

var subjectIDs = append([]struct {
	name  string
	valid bool
}{
	{"*", true},
}, objectIDs...)

var caveats = []struct {
	name    string
	context map[string]any
	valid   bool
}{
	{"test", map[string]any{}, true},
	{"", map[string]any{"a": "b"}, false},
	{strings.Repeat("f", maxCaveatName), map[string]any{"a": "b"}, true},
	{strings.Repeat("f", maxCaveatName+1), map[string]any{"a": "b"}, false},
	{"test", map[string]any{"a": "b"}, true},
	{"test", nil, true},
	{"test", generateMap(128), true},
}

type relationValidity int

const (
	alwaysInvalid relationValidity = iota
	alwaysValid
	onlySubjectValid
)

type relationEntry struct {
	name     string
	validity relationValidity
}

var knownGoodObjectRef = &v1.ObjectReference{
	ObjectType: "user",
	ObjectId:   "testuser",
}

var knownGoodSubjectRef = &v1.SubjectReference{
	Object:           knownGoodObjectRef,
	OptionalRelation: "member",
}

var relations = []relationEntry{
	{"", onlySubjectValid},
	{"foo", alwaysValid},
	{"bar", alwaysValid},
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
	{strings.Repeat("f", minRelationLength-1), alwaysInvalid},
	{strings.Repeat("f", minRelationLength), alwaysValid},
	{strings.Repeat("\u0394", minRelationLength), alwaysInvalid},
	{strings.Repeat("\n", minRelationLength), alwaysInvalid},
	{strings.Repeat("_", minRelationLength), alwaysInvalid},
	{strings.Repeat("-", minRelationLength), alwaysInvalid},
	{strings.Repeat("/", minRelationLength), alwaysInvalid},
	{strings.Repeat("\\", minRelationLength), alwaysInvalid},
	{strings.Repeat("f", maxRelationLength), alwaysValid},
	{strings.Repeat("f", maxRelationLength+1), alwaysInvalid},
}

func TestV1CoreObjectValidity(t *testing.T) {
	for _, ns := range namespaces {
		for _, objectID := range objectIDs {
			for _, subjectID := range subjectIDs {
				for _, relation := range relations {
					testName := fmt.Sprintf("%s:%s#%s@%s:%s", ns.name, objectID.name, relation.name, ns.name, subjectID.name)
					ns := ns
					objectID := objectID
					subjectID := subjectID
					relation := relation

					t.Run(testName, func(t *testing.T) {
						t.Parallel()

						require := require.New(t)

						objRef := &v1.ObjectReference{
							ObjectType: ns.name,
							ObjectId:   objectID.name,
						}
						objRefValid := ns.valid && objectID.valid
						err := objRef.Validate()
						require.Equal(objRefValid, err == nil, "should be valid: %v %s", objRefValid, err)

						subjObjRef := &v1.ObjectReference{
							ObjectType: ns.name,
							ObjectId:   subjectID.name,
						}
						subObjRefValid := ns.valid && subjectID.valid
						err = subjObjRef.Validate()
						require.Equal(subObjRefValid, err == nil, "should be valid: %v %s | ns: %v | subj: %v", subObjRefValid, err, ns.valid, subjectID.valid)

						subRef := &v1.SubjectReference{
							Object:           subjObjRef,
							OptionalRelation: relation.name,
						}
						subjectValid := ns.valid && subjectID.valid && (relation.validity == alwaysValid || relation.validity == onlySubjectValid)
						err = subRef.Validate()
						require.Equal(subjectValid, err == nil, "should be valid: %v %s", subjectValid, err)

						asResource := &v1.Relationship{
							Resource: objRef,
							Relation: relation.name,
							Subject:  knownGoodSubjectRef,
						}
						asResourceValid := objRefValid && relation.validity == alwaysValid
						err = asResource.Validate()
						require.Equal(asResourceValid, err == nil, "should be valid: %v %s | relation: `%s`", asResourceValid, err, relation.name)

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
						filterValid := (ns.valid || ns.name == "")
						err = justNS.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						objectIDFilter := &v1.RelationshipFilter{
							ResourceType:       ns.name,
							OptionalResourceId: objectID.name,
						}
						filterValid = (ns.valid || ns.name == "") && (objectID.valid || objectID.name == "")
						err = objectIDFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						objectRelationFilter := &v1.RelationshipFilter{
							ResourceType:     ns.name,
							OptionalRelation: relation.name,
						}
						filterValid = (ns.valid || ns.name == "") && (relation.validity == alwaysValid || relation.name == "")
						err = objectRelationFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						fullObjectFilter := &v1.RelationshipFilter{
							ResourceType:       ns.name,
							OptionalResourceId: objectID.name,
							OptionalRelation:   relation.name,
						}
						filterValid = (ns.valid || ns.name == "") && (objectID.valid || objectID.name == "") &&
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
								OptionalSubjectId: subjectID.name,
							},
						}
						filterValid = ns.valid && (subjectID.valid || subjectID.name == "")
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
						filterValid = ns.valid && (relation.validity == alwaysValid || relation.validity == onlySubjectValid)
						err = subjectRelationFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						fullSubjectFilter := &v1.RelationshipFilter{
							ResourceType: knownGoodObjectRef.ObjectType,
							OptionalSubjectFilter: &v1.SubjectFilter{
								SubjectType:       ns.name,
								OptionalSubjectId: subjectID.name,
								OptionalRelation: &v1.SubjectFilter_RelationFilter{
									Relation: relation.name,
								},
							},
						}
						filterValid = ns.valid && (subjectID.valid || subjectID.name == "") && (relation.validity == alwaysValid || relation.validity == onlySubjectValid)
						err = fullSubjectFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)
					})
				}
			}
		}
	}
}

func TestV1CaveatValidity(t *testing.T) {
	for _, caveat := range caveats {
		testName := fmt.Sprintf("caveat->%s_context->%v", caveat.name, caveat.context)
		caveat := caveat
		t.Run(testName, func(t *testing.T) {
			require := require.New(t)

			strct, err := structpb.NewStruct(caveat.context)
			require.NoError(err)

			optionalCaveat := &v1.ContextualizedCaveat{
				CaveatName: caveat.name,
				Context:    strct,
			}
			err = optionalCaveat.Validate()
			require.Equal(caveat.valid, err == nil, "should be valid: %v %s", caveat.valid, err)

			rel := &v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "test",
					ObjectId:   "test",
				},
				Relation: "test",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "test",
						ObjectId:   "test",
					},
				},
				OptionalCaveat: optionalCaveat,
			}
			err = rel.Validate()
			require.Equal(caveat.valid, err == nil, "should be valid: %v %s", caveat.valid, err)
		})
	}
}

func TestWildcardSubjectRelation(t *testing.T) {
	subjObjRef := &v1.ObjectReference{
		ObjectType: "somenamespace",
		ObjectId:   "*",
	}
	subRef := &v1.SubjectReference{
		Object:           subjObjRef,
		OptionalRelation: "somerelation",
	}
	require.Error(t, subRef.HandwrittenValidate())
}

func TestWildcardSubjectRelationEmpty(t *testing.T) {
	subjObjRef := &v1.ObjectReference{
		ObjectType: "somenamespace",
		ObjectId:   "*",
	}
	subRef := &v1.SubjectReference{
		Object: subjObjRef,
	}
	require.NoError(t, subRef.HandwrittenValidate())
}

func generateMap(length int) map[string]any {
	output := make(map[string]any, length)
	for i := 0; i < length; i++ {
		random := randString(32)
		output[random] = random
	}
	return output
}

var randInput = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = randInput[rand.Intn(len(randInput))] //nolint:gosec
	}
	return string(b)
}
