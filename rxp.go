package rxp

import (
	"github.com/relexec/rxp/core"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kindversion"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/query"
	"github.com/relexec/rxp/system"
)

type Type = core.Type

const (
	TypeSystem      = core.TypeSystem
	TypeDomain      = core.TypeDomain
	TypeKind        = core.TypeKind
	TypeKindVersion = core.TypeKindVersion
	TypeObject      = core.TypeObject
)

type Scope = core.Scope

const (
	ScopeGlobal = core.ScopeGlobal
	ScopeSystem = core.ScopeSystem
	ScopeDomain = core.ScopeDomain
)

type Caller = core.Caller

var (
	CallerToContext           = core.CallerToContext
	CallerFromContext         = core.CallerFromContext
	ErrCallerInvalid          = errors.ErrCallerInvalid
	ErrCallerIdentityRequired = errors.ErrCallerIdentityRequired
)

type Generation = core.Generation
type Label = core.Label
type Labels = core.Labels

type Error = errors.Error

var (
	ErrCodeInternal           = errors.ErrCodeInternal
	ErrCodeBadRequest         = errors.ErrCodeBadRequest
	ErrCodeConflict           = errors.ErrConflict
	ErrCodeNotFound           = errors.ErrNotFound
	ErrCodePreconditionFailed = errors.ErrCodePreconditionFailed
	ErrCodeTooManyRequests    = errors.ErrCodeTooManyRequests
	ErrNotFound               = errors.ErrNotFound
	ErrMissingIdentity        = errors.ErrMissingIdentity
	ErrConflict               = errors.ErrConflict
	ErrPreconditionFailed     = errors.ErrPreconditionFailed
	ErrInvalidReadRequest     = errors.ErrInvalidReadRequest
	ErrInvalidWriteRequest    = errors.ErrInvalidWriteRequest
)

type Expression = query.Expression

var (
	And = query.And
	Or  = query.Or

	ErrQueryInvalid                    = errors.ErrQueryInvalid
	ErrQueryExpressionRequired         = errors.ErrQueryExpressionRequired
	ErrQueryInvalidExpression          = errors.ErrQueryInvalidExpression
	ErrQueryKindPredicateInObjectQuery = errors.ErrQueryKindPredicateInObjectQuery
)

type System = system.System

var (
	ErrSystemInvalid      = errors.ErrSystemInvalid
	ErrSystemUUIDRequired = errors.ErrSystemUUIDRequired
	ErrSystemUnknown      = errors.ErrSystemUnknown
)

type Domain = domain.Domain
type DomainName = domain.Name

var (
	ErrDomainInvalid                   = errors.ErrDomainInvalid
	ErrDomainParentRootRequired        = errors.ErrDomainParentRootRequired
	ErrDomainRootSystemDifferent       = errors.ErrDomainRootSystemDifferent
	ErrDomainParentSystemDifferent     = errors.ErrDomainParentSystemDifferent
	ErrDomainParentNotFound            = errors.ErrDomainParentNotFound
	ErrDomainUUIDRequired              = errors.ErrDomainUUIDRequired
	ErrDomainNameRequired              = errors.ErrDomainNameRequired
	ErrDomainNameInvalid               = errors.ErrDomainNameInvalid
	ErrDomainNameInvalidCharacters     = errors.ErrDomainNameInvalidCharacters
	ErrDomainNameMaxLengthExceeded     = errors.ErrDomainNameMaxLengthExceeded
	ErrDomainNameRepeatedPeriods       = errors.ErrDomainNameRepeatedPeriods
	ErrDomainNameInvalidFirstCharacter = errors.ErrDomainNameInvalidFirstCharacter
)

type Kind = kind.Kind
type KindName = kind.Name

var (
	ErrKindInvalid                   = errors.ErrKindInvalid
	ErrKindUUIDRequired              = errors.ErrKindUUIDRequired
	ErrKindNameRequired              = errors.ErrKindNameRequired
	ErrKindNameInvalid               = errors.ErrKindNameInvalid
	ErrKindNameEmpty                 = errors.ErrKindNameEmpty
	ErrKindNameInvalidCharacters     = errors.ErrKindNameInvalidCharacters
	ErrKindNameRepeatedPeriods       = errors.ErrKindNameRepeatedPeriods
	ErrKindNameInvalidFirstCharacter = errors.ErrKindNameInvalidFirstCharacter
	ErrKindUnknown                   = errors.ErrKindUnknown
)

type KindVersion = kindversion.KindVersion
type KindVersionName = kindversion.Name

var (
	ErrKindVersionInvalid         = errors.ErrKindVersionInvalid
	ErrKindVersionKindRequired    = errors.ErrKindVersionKindRequired
	ErrKindVersionVersionRequired = errors.ErrKindVersionVersionRequired
	ErrKindVersionSchemaRequired  = errors.ErrKindVersionSchemaRequired
	ErrKindVersionUnknown         = errors.ErrKindVersionUnknown
)

type Object = object.Object

var (
	ErrObjectInvalid             = errors.ErrObjectInvalid
	ErrObjectNil                 = errors.ErrObjectNil
	ErrObjectKindVersionRequired = errors.ErrObjectKindVersionRequired
	ErrObjectNameRequired        = errors.ErrObjectNameRequired
	ErrObjectUUIDRequired        = errors.ErrObjectUUIDRequired
	ErrObjectDomainRequired      = errors.ErrObjectDomainRequired
)
