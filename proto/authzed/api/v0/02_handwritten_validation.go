// The contents of this file are hand-written to add HandwrittenValidate to select message types

package v0

func (m *CheckRequest) HandwrittenValidate() error {
	if m.GetTestUserset() != nil && m.GetTestUserset().GetObjectId() == "*" {
		return ObjectAndRelationValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *ContentChangeCheckRequest) HandwrittenValidate() error {
	if m.GetTestUserset() != nil && m.GetTestUserset().GetObjectId() == "*" {
		return ObjectAndRelationValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *ExpandRequest) HandwrittenValidate() error {
	if m.GetUserset() != nil && m.GetUserset().GetObjectId() == "*" {
		return ObjectAndRelationValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *LookupRequest) HandwrittenValidate() error {
	if m.GetUser() != nil && m.GetUser().GetObjectId() == "*" {
		return ObjectAndRelationValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}
