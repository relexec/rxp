package apirun

import (
	apierrors "github.com/relexec/rxp/api/errors"
	apiobject "github.com/relexec/rxp/api/object"
	apiquery "github.com/relexec/rxp/api/query"
)

type UUIDPredicate struct {
	apiquery.BasePredicate
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
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// UUIDEqual returns an Expression that will match runs having a particular
// UUID.
func UUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match runs that have any of a set of
// specified Request UUIDs.
func UUIDIn(uuids ...string) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type RootUUIDPredicate struct {
	apiquery.BasePredicate
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
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// RootUUIDEqual returns an Expression that will match runs having a Root
// with a particular UUID.
func RootUUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: RootUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// RootUUIDIn returns an Expression that will match runs that have Roots
// with any of a set of specified UUIDs.
func RootUUIDIn(uuids ...string) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return RootUUIDEqual(uuids[0])
	}
	return apiquery.UnaryExpression{
		Predicate: RootUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type RootPredicate struct {
	apiquery.BasePredicate
}

func (p RootPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []*Run:
		return nil
	case *Run:
		return nil
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// RootEqual returns an Expression that will match runs having a particular
// Root.
func RootEqual(run *Run) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: RootPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: run,
			},
		},
	}
}

// RootIn returns an Expression that will match runs that have any of the
// supplied Roots.
func RootIn(runs ...*Run) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(runs) == 1 {
		return RootEqual(runs[0])
	}
	return apiquery.UnaryExpression{
		Predicate: RootPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: runs,
			},
		},
	}
}

type TargetUUIDPredicate struct {
	apiquery.BasePredicate
}

func (p TargetUUIDPredicate) Validate() error {
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
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// TargetUUIDEqual returns an Expression that will match runs having a Target
// with a particular UUID.
func TargetUUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: TargetUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// TargetUUIDIn returns an Expression that will match runs that have Targets
// with any of a set of specified UUIDs.
func TargetUUIDIn(uuids ...string) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return TargetUUIDEqual(uuids[0])
	}
	return apiquery.UnaryExpression{
		Predicate: TargetUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

type TargetPredicate struct {
	apiquery.BasePredicate
}

func (p TargetPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []*apiobject.Object:
		return nil
	case *apiobject.Object:
		return nil
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// TargetEqual returns an Expression that will match runs having a particular
// Target.
func TargetEqual(target *apiobject.Object) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: TargetPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: target,
			},
		},
	}
}

// TargetIn returns an Expression that will match runs that have any of the
// supplied Targets.
func TargetIn(targets ...*apiobject.Object) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(targets) == 1 {
		return TargetEqual(targets[0])
	}
	return apiquery.UnaryExpression{
		Predicate: TargetPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: targets,
			},
		},
	}
}
