package predicate

import (
	"github.com/relexec/rxp/types"
)

// Option modifies a Predicate returned from New.
type Option func(*Predicate)

// New returns a new [Predicate].
func New(opts ...Option) Predicate {
	p := Predicate{}
	for _, opt := range opts {
		opt(&p)
	}
	return p
}

// WithOperator sets the Predicate's Operator.
func WithOperator(op types.PredicateOperator) Option {
	return func(p *Predicate) {
		p.op = op
	}
}

// WithNegated sets the Predicate's Negated to the supplied value.
func WithNegated(val bool) Option {
	return func(p *Predicate) {
		p.negated = val
	}
}

// WithValues sets the Predicate's Values to the supplied values.
func WithValues(vals ...any) Option {
	return func(p *Predicate) {
		p.values = vals
	}
}
