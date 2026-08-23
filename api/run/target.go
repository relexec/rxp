package apirun

import (
	"github.com/relexec/rxp/api"
	apicore "github.com/relexec/rxp/api/core"
	"github.com/relexec/rxp/errors"
)

// Target uniquely identifies an Object generation that contains the
// *definition* of the thing that will be executed by Run.
type Target struct {
	// NOTE(jaypipes): Record.SystemInternalID points to the internal ID of the
	// Object+Generation record this RunTarget points at.
	apicore.Record
	// KindVersionName is the kind and version identifier for the type of
	// Object describing the Runnable.
	KindVersionName api.KindVersionName
	// System is System in which the target's Object resides. If empty, the
	// host System for the entity executing the API call is used.
	System *api.System
	// Domain is the optional Domain the target's Object is scoped to.
	Domain *api.Domain
	// UUID is the identifier of the Object describing the Runnable.
	UUID string
	// Generation is the Object generation.
	Generation api.Generation
}

// Validate returns an error if the Target is not valid.
func (t Target) Validate() error {
	if t.KindVersionName == "" {
		return errors.ErrRunTargetKindVersionNameRequired
	}
	if t.UUID == "" {
		return errors.ErrRunTargetUUIDRequired
	}
	if t.Generation == 0 {
		return errors.ErrRunTargetGenerationRequired
	}
	if t.System != nil {
		err := t.System.Validate()
		if err != nil {
			return errors.ErrRunTargetSystemInvalid
		}
	}
	if t.Domain != nil {
		err := t.Domain.Validate()
		if err != nil {
			return errors.ErrRunTargetDomainInvalid
		}
	}
	return nil
}
