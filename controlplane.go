package rxp

import (
	"context"

	"github.com/relexec/rxp/api"
	apidomain "github.com/relexec/rxp/api/domain"
	apikind "github.com/relexec/rxp/api/kind"
	apirun "github.com/relexec/rxp/api/run"
	apisystem "github.com/relexec/rxp/api/system"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/query"
)

// ControlPlane is the interface that rxp backends implement for control plane
// operations.
type ControlPlane interface {
	// SystemRead reads a System from persistent storage.
	SystemRead(
		context.Context,
		apisystem.Selector,
	) (*apisystem.System, error)
	// SystemWrite atomically writes the supplied System to persistent storage.
	SystemWrite(
		context.Context,
		apisystem.System,
	) error
	// SystemQuery queries zero or more Systems from persistent storage.
	SystemQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*apisystem.System], error)
	// DomainRead reads a Domain from persistent storage.
	DomainRead(
		context.Context,
		apidomain.Selector,
	) (*apidomain.Domain, error)
	// DomainWrite atomically writes the supplied Domain to persistent storage.
	DomainWrite(
		context.Context,
		apidomain.Domain,
	) error
	// DomainQuery queries zero or more Domains from persistent storage.
	DomainQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*apidomain.Domain], error)

	// KindRead reads a Kind from persistent storage.
	KindRead(
		context.Context,
		apikind.Selector,
	) (*apikind.Kind, error)
	// KindWrite atomically writes the supplied Kind to persistent storage.
	KindWrite(
		context.Context,
		apikind.Kind,
	) error
	// KindQuery queries zero or more Kinds from persistent storage.
	KindQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*apikind.Kind], error)
	// KindVersionRead reads a KindVersion from persistent storage.
	KindVersionRead(
		context.Context,
		kindversion.Selector,
	) (*api.KindVersion, error)
	// KindVersionWrite atomically writes the supplied KindVersion to
	// persistent storage.
	KindVersionWrite(
		context.Context,
		api.KindVersion,
	) error
	// KindVersionQuery queries zero or more KindVersions from persistent
	// storage.
	KindVersionQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*api.KindVersion], error)

	// ObjectRead reads a single Object from persistent storage.
	ObjectRead(
		context.Context,
		api.KindVersionName,
		object.Selector,
	) (*api.Object, error)
	// ObjectWrite persists a single supplied Object to backend storage, Note
	// that on successful write, the newly-created or updated Object is
	// returned.
	ObjectWrite(
		context.Context,
		api.Object,
	) (*api.Object, error)
	// ObjectQuery queries zero or more Objects of a specified kind or
	// kindversion from persistent storage.
	ObjectQuery(
		context.Context,
		api.KindVersionName,
		query.Expression,
		...query.Option,
	) (*query.Result[*api.Object], error)

	// RunRead reads a single Run from persistent storage.
	RunRead(
		context.Context,
		apirun.Target,
		apirun.Selector,
	) (*apirun.Run, error)
	// RunWrite persists a single supplied Run to backend storage, Note
	// that on successful write, the newly-created or updated Run is
	// returned.
	RunWrite(
		context.Context,
		apirun.Run,
	) (*apirun.Run, error)
	// RunQuery queries zero or more Runs from persistent storage.
	RunQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*apirun.Run], error)
}
