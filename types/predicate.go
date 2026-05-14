package types

// PredicateOperator is the type of comparison operator used by a Predicate.
type PredicateOperator int

const (
	PredicateOperatorEqual PredicateOperator = iota
	PredicateOperatorIn
)

// Predicate describes a single boolean expression. Predicates are used by
// `rxp` backend implementations to translate the user's matching/filtering
// intent into backend-specific filtering expressions (e.g. SQL statements).
type Predicate interface {
	Validatable
	// Operator returns the Predicate's Operator.
	Operator() PredicateOperator
	// Negated returns true if the Predicate's Operator should be negated. For
	// example, if Operator() returns PredicateOperatorEqual and Negated()
	// returns true, the Predicate will evaluate to true if the target field is
	// NOT equal to the Values.
	Negated() bool
	// Values returns the values that are compared to the target field.
	Values() []any
}
