package thoriumfacts

type Repository interface {
	// GetFacts returns a list of facts for the given user.
	GetFacts(userID string) ([]Fact, error)

	// AddFact adds a new fact to the repository.
	AddFact(userID string, fact Fact) error

	// DeleteFact deletes a fact from the repository.
	DeleteFact(userID string, fact Fact) error
}
