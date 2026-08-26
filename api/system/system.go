package apisystem

import (
	apicore "github.com/relexec/rxp/api/core"
	apierrors "github.com/relexec/rxp/api/errors"
)

// System represents the boundaries of an rxp system installation.
//
// An rxp system installation represents an independent division of rxp-managed
// data.
type System struct {
	apicore.Record
	// UUID contains the System's globally-unique identifier.
	UUID string `json:"uuid"`
	// Tag contains an optional string tag for the System. Note this is not
	// called "name" because a Name in rxp has a specific semantic meaning that
	// reflects the uniqueness constraint its value. Tags have no such
	// uniqueness constraint.
	Tag string `json:"tag,omitempty"`
}

// Validate returns an error if the System is invalid.
func (s System) Validate() error {
	if s.UUID == "" {
		return apierrors.ErrSystemUUIDRequired
	}
	return nil
}
