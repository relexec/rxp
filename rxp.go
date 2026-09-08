package rxp

import (
	"github.com/relexec/rxp/core"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kindversion"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/system"
)

type Scope = core.Scope

const (
	ScopeGlobal = core.ScopeGlobal
	ScopeSystem = core.ScopeSystem
	ScopeDomain = core.ScopeDomain
)

type Caller = core.Caller
type Generation = core.Generation
type Label = core.Label
type Labels = core.Labels

type System = system.System
type Domain = domain.Domain
type DomainName = domain.Name
type Kind = kind.Kind
type KindName = kind.Name
type KindVersion = kindversion.KindVersion
type KindVersionName = kindversion.Name
type Object = object.Object
