package run

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/query"
)

type UUIDPredicate struct {
	query.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// UUIDEqual returns an Expression that will match runs having a particular
// UUID.
func UUIDEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		Predicate: UUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match runs that have any of a set of
// specified Request UUIDs.
func UUIDIn(uuids ...string) query.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return query.UnaryExpression{
		Predicate: UUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type RootUUIDPredicate struct {
	query.BasePredicate
}

func (p RootUUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// RootUUIDEqual returns an Expression that will match runs having a Root
// with a particular UUID.
func RootUUIDEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		Predicate: RootUUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// RootUUIDIn returns an Expression that will match runs that have Roots
// with any of a set of specified UUIDs.
func RootUUIDIn(uuids ...string) query.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return RootUUIDEqual(uuids[0])
	}
	return query.UnaryExpression{
		Predicate: RootUUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type RootPredicate struct {
	query.BasePredicate
}

func (p RootPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// RootEqual returns an Expression that will match runs having a particular
// Root.
func RootEqual(run *Run) query.Expression {
	return query.UnaryExpression{
		Predicate: RootPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: run,
			},
		},
	}
}

// RootIn returns an Expression that will match runs that have any of the
// supplied Roots.
func RootIn(runs ...*Run) query.Expression {
	// flatten IN to = when there's only one value...
	if len(runs) == 1 {
		return RootEqual(runs[0])
	}
	return query.UnaryExpression{
		Predicate: RootPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: runs,
			},
		},
	}
}

type TargetUUIDPredicate struct {
	query.BasePredicate
}

func (p TargetUUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// TargetUUIDEqual returns an Expression that will match runs having a Target
// with a particular UUID.
func TargetUUIDEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		Predicate: TargetUUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// TargetUUIDIn returns an Expression that will match runs that have Targets
// with any of a set of specified UUIDs.
func TargetUUIDIn(uuids ...string) query.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return TargetUUIDEqual(uuids[0])
	}
	return query.UnaryExpression{
		Predicate: TargetUUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type TargetPredicate struct {
	query.BasePredicate
}

func (p TargetPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// TargetEqual returns an Expression that will match runs having a particular
// Target.
func TargetEqual(target *object.Object) query.Expression {
	return query.UnaryExpression{
		Predicate: TargetPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: target,
			},
		},
	}
}

// TargetIn returns an Expression that will match runs that have any of the
// supplied Targets.
func TargetIn(targets ...*object.Object) query.Expression {
	// flatten IN to = when there's only one value...
	if len(targets) == 1 {
		return TargetEqual(targets[0])
	}
	return query.UnaryExpression{
		Predicate: TargetPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: targets,
			},
		},
	}
}
