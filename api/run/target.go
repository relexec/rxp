package apirun

import (
	apicore "github.com/relexec/rxp/api/core"
	apidomain "github.com/relexec/rxp/api/domain"
	apierrors "github.com/relexec/rxp/api/errors"
	apikindversion "github.com/relexec/rxp/api/kindversion"
	apisystem "github.com/relexec/rxp/api/system"
)

// Target uniquely identifies an Object generation that contains the
// *definition* of the thing that will be executed by Run.
type Target struct {
	// NOTE(jaypipes): Record.SystemInternalID points to the internal ID of the
	// Object+Generation record this Target points at.
	apicore.Record
	// KindVersionName is the kind and version identifier for the type of
	// Object describing the Runnable.
	KindVersionName apikindversion.Name
	// System is System in which the Target's Object resides. If empty, the
	// host System for the entity executing the API call is used.
	System *apisystem.System
	// Domain is the optional Domain the Target's Object is scoped to.
	Domain *apidomain.Domain
	// UUID is the identifier of the Object describing the Runnable.
	UUID string
	// Generation is the Object generation.
	Generation apicore.Generation
}

// Validate returns an error if the Target is not valid.
func (t Target) Validate() error {
	if t.KindVersionName == "" {
		return apierrors.ErrRunTargetKindVersionNameRequired
	}
	if t.UUID == "" {
		return apierrors.ErrRunTargetUUIDRequired
	}
	if t.Generation == 0 {
		return apierrors.ErrRunTargetGenerationRequired
	}
	if t.System != nil {
		err := t.System.Validate()
		if err != nil {
			return apierrors.ErrRunTargetSystemInvalid
		}
	}
	if t.Domain != nil {
		err := t.Domain.Validate()
		if err != nil {
			return apierrors.ErrRunTargetDomainInvalid
		}
	}
	return nil
}
