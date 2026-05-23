package expression

// Expression describes a filtering expression. Similar to Predicates,
// Expressions are used by `rxp` backend implementations to translate the
// user's matching/filtering intent into backend-specific filtering expressions
// (e.g. SQL statements).
type Expression interface {
	// Unary returns true if the Expression can be reduced to a single
	// Predicate.
	Unary() bool
}
