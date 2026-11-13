package validation_test

import (
	"testing"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// FuzzObjectID tests the extended object ID validation with fuzzing.
// The v1 object ID regex pattern is: ^(([a-zA-Z0-9/_|\-=+]{1,})|\*)$
// This fuzz test ensures the validation doesn't panic on arbitrary input.
func FuzzObjectID(f *testing.F) {
	// Add seed corpus with known valid and invalid object IDs
	seeds := []string{
		// Valid object IDs
		"a",
		"1",
		"a-",
		"A-A-A",
		"123e4567-e89b-12d3-a456-426614174000",
		"authn|someauthnvalue",
		"authn|",
		"authn|a-b-c-d-e",
		"authn|a_b",
		"--=base64YWZzZGZh-ZHNmZHPwn5iK8J+YivC/fmIrwn5iK==",
		"-",
		"*",
		// Invalid object IDs
		"",
		"  ",
		"a@b.com",
		"test@example.com",
		"test*",
		"*test",
		"t*t",
		"test@example.*",
		"authn|a@b.com",
		// Edge cases
		"\x00",
		"\n",
		"\t",
		"foo\x00bar",
		"\u0394\u0394\u0394\u0394",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, objectID string) {
		// Test ObjectReference validation
		objRef := &v1.ObjectReference{
			ObjectType: "testtype",
			ObjectId:   objectID,
		}
		// We only care that this doesn't panic - validation result can be either valid or invalid
		_ = objRef.Validate()

		// Test SubjectReference validation with the object ID
		subjRef := &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "testtype",
				ObjectId:   objectID,
			},
		}
		_ = subjRef.Validate()
		_ = subjRef.HandwrittenValidate()

		// Test RelationshipFilter validation
		filter := &v1.RelationshipFilter{
			ResourceType:       "testtype",
			OptionalResourceId: objectID,
		}
		_ = filter.Validate()
		_ = filter.HandwrittenValidate()

		// Test SubjectFilter validation
		subjectFilter := &v1.SubjectFilter{
			SubjectType:       "testtype",
			OptionalSubjectId: objectID,
		}
		_ = subjectFilter.Validate()
		_ = subjectFilter.HandwrittenValidate()

		// Test full Relationship validation
		rel := &v1.Relationship{
			Resource: &v1.ObjectReference{
				ObjectType: "testtype",
				ObjectId:   objectID,
			},
			Relation: "testrelation",
			Subject: &v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: "testtype",
					ObjectId:   "testsubject",
				},
			},
		}
		_ = rel.Validate()
		_ = rel.HandwrittenValidate()
	})
}

// FuzzNamespace tests the namespace/object type validation with fuzzing.
// The v1 namespace regex pattern is: ^([a-z][a-z0-9_]{1,61}[a-z0-9]/)*[a-z][a-z0-9_]{1,62}[a-z0-9]$
func FuzzNamespace(f *testing.F) {
	seeds := []string{
		// Valid namespaces
		"foo",
		"bar",
		"foo1",
		"bar1",
		"foo_bar",
		"foo/bar",
		"foo/bar/baz",
		// Invalid namespaces
		"",
		"...",
		"ab",
		"Foo1",
		"foo_bar_",
		"foo/b",
		"Foo/bar",
		// Edge cases
		"\x00",
		"\n",
		"\t",
		"\u0394\u0394\u0394",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, namespace string) {
		objRef := &v1.ObjectReference{
			ObjectType: namespace,
			ObjectId:   "testid",
		}
		_ = objRef.Validate()
	})
}

// FuzzRelation tests the relation name validation with fuzzing.
// The v1 relation regex pattern is: ^[a-z][a-z0-9_]{1,62}[a-z0-9]$
func FuzzRelation(f *testing.F) {
	seeds := []string{
		// Valid relations
		"foo",
		"bar",
		"foo1",
		"bar1",
		"foo_bar",
		"member",
		"viewer",
		"owner",
		// Invalid relations
		"",
		"ab",
		"Foo1",
		"foo_bar_",
		"foo/bar",
		// Edge cases
		"\x00",
		"\n",
		"\t",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, relation string) {
		rel := &v1.Relationship{
			Resource: &v1.ObjectReference{
				ObjectType: "testtype",
				ObjectId:   "testid",
			},
			Relation: relation,
			Subject: &v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: "testtype",
					ObjectId:   "testsubject",
				},
			},
		}
		_ = rel.Validate()

		// Also test relation in SubjectReference
		subjRef := &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "testtype",
				ObjectId:   "testid",
			},
			OptionalRelation: relation,
		}
		_ = subjRef.Validate()
	})
}

// FuzzCaveatName tests the caveat name validation with fuzzing.
func FuzzCaveatName(f *testing.F) {
	seeds := []string{
		"test",
		"my_caveat",
		"caveat123",
		"",
		"\x00",
		"\n",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, caveatName string) {
		caveat := &v1.ContextualizedCaveat{
			CaveatName: caveatName,
		}
		_ = caveat.Validate()
	})
}
