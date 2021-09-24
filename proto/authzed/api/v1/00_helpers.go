package v1

// ConsistencyFull configures a request to be performed with the most recent
// data in the database.
func ConsistencyFull() *Consistency {
	return &Consistency{Requirement: &Consistency_FullyConsistent{true}}
}

// ConsistencyAtLeast configures a request to be performed with data at least
// as fresh as the event that produced the provided ZedToken.
func ConsistencyAtLeast(zedtoken string) *Consistency {
	return &Consistency{Requirement: &Consistency_AtLeastAsFresh{&ZedToken{Token: zedtoken}}}
}

// ConsistencyMinLatency configures a request to be performed with a revision
// determined by the server to optimized for performance.
//
// When no explicit consistency is provided to a request, this is the default.
func ConsistencyMinLatency() *Consistency {
	return &Consistency{Requirement: &Consistency_MinimizeLatency{true}}
}

// ConsistencyExact configures a request to be performed with the data at the
// event that produced the provided ZedToken.
func ConsistencyExact(zedtoken string) *Consistency {
	return &Consistency{Requirement: &Consistency_AtExactSnapshot{&ZedToken{Token: zedtoken}}}
}
