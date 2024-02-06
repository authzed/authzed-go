package validation_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

func TestV1UpdateValidity(t *testing.T) {
	tcs := []struct {
		name         string
		operation    v1.RelationshipUpdate_Operation
		relationship *v1.Relationship
		isValid      bool
	}{
		{
			"valid create",
			v1.RelationshipUpdate_OPERATION_CREATE,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "foo",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			true,
		},
		{
			"valid touch",
			v1.RelationshipUpdate_OPERATION_TOUCH,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "foo",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			true,
		},
		{
			"valid delete",
			v1.RelationshipUpdate_OPERATION_DELETE,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "foo",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			true,
		},
		{
			"invalid create",
			v1.RelationshipUpdate_OPERATION_CREATE,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			false,
		},
		{
			"invalid touch",
			v1.RelationshipUpdate_OPERATION_TOUCH,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			false,
		},
		{
			"invalid delete",
			v1.RelationshipUpdate_OPERATION_DELETE,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			false,
		},
		{
			"invalid operation",
			v1.RelationshipUpdate_OPERATION_UNSPECIFIED,
			&v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "document",
					ObjectId:   "foo",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   "sarah",
					},
				},
			},
			false,
		},
		{
			"empty relationship",
			v1.RelationshipUpdate_OPERATION_UNSPECIFIED,
			&v1.Relationship{},
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require := require.New(t)
			update := &v1.RelationshipUpdate{
				Operation:    tc.operation,
				Relationship: tc.relationship,
			}
			err := update.Validate()

			if tc.isValid {
				require.NoError(err)
			} else {
				require.NotNil(err)
			}
		})
	}
}

func TestV1PreconditionValidity(t *testing.T) {
	tcs := []struct {
		name      string
		operation v1.Precondition_Operation
		filter    *v1.RelationshipFilter
		isValid   bool
	}{
		{
			"valid must match",
			v1.Precondition_OPERATION_MUST_MATCH,
			&v1.RelationshipFilter{
				ResourceType: "document",
			},
			true,
		},
		{
			"valid must not match",
			v1.Precondition_OPERATION_MUST_NOT_MATCH,
			&v1.RelationshipFilter{
				ResourceType: "document",
			},
			true,
		},
		{
			"invalid operation",
			v1.Precondition_OPERATION_UNSPECIFIED,
			&v1.RelationshipFilter{
				ResourceType: "document",
			},
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require := require.New(t)
			update := &v1.Precondition{
				Operation: tc.operation,
				Filter:    tc.filter,
			}
			err := update.Validate()

			if tc.isValid {
				require.NoError(err)
			} else {
				require.NotNil(err)
			}
		})
	}
}
