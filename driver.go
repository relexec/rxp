package rxp

import (
	"context"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/query"
	"github.com/relexec/rxp/run"
	"github.com/relexec/rxp/system"
)

// Driver is the primary interface that rxp backends implement.
type Driver interface {
	// SystemRead reads a System from persistent storage.
	SystemRead(
		context.Context,
		system.Selector,
	) (*system.System, error)
	// SystemWrite atomically writes the supplied System to persistent storage.
	SystemWrite(
		context.Context,
		*system.System,
	) error
	// SystemQuery queries zero or more Systems from persistent storage.
	SystemQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*system.System], error)
	// DomainRead reads a Domain from persistent storage.
	DomainRead(
		context.Context,
		domain.Selector,
	) (*domain.Domain, error)
	// DomainWrite atomically writes the supplied Domain to persistent storage.
	DomainWrite(
		context.Context,
		*domain.Domain,
	) error
	// DomainQuery queries zero or more Domains from persistent storage.
	DomainQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*domain.Domain], error)

	// KindRead reads a Kind from persistent storage.
	KindRead(
		context.Context,
		kind.Selector,
	) (*kind.Kind, error)
	// KindWrite atomically writes the supplied Kind to persistent storage.
	KindWrite(
		context.Context,
		*kind.Kind,
	) error
	// KindQuery queries zero or more Kinds from persistent storage.
	KindQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*kind.Kind], error)
	// KindVersionRead reads a KindVersion from persistent storage.
	KindVersionRead(
		context.Context,
		kindversion.Selector,
	) (*kindversion.KindVersion, error)
	// KindVersionWrite atomically writes the supplied KindVersion to
	// persistent storage.
	KindVersionWrite(
		context.Context,
		*kindversion.KindVersion,
	) error
	// KindVersionQuery queries zero or more KindVersions from persistent
	// storage.
	KindVersionQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*kindversion.KindVersion], error)

	// ObjectRead reads a single Object from persistent storage.
	ObjectRead(
		context.Context,
		api.KindVersionName,
		object.Selector,
	) (*object.Object, error)
	// ObjectWrite persists a single supplied Object to backend storage, Note
	// that on successful write, the newly-created or updated Object is
	// returned.
	ObjectWrite(
		context.Context,
		object.Object,
	) (*object.Object, error)
	// ObjectQuery queries zero or more Objects of a specified kind or
	// kindversion from persistent storage.
	ObjectQuery(
		context.Context,
		api.KindVersionName,
		query.Expression,
		...query.Option,
	) (*query.Result[*object.Object], error)

	// RunRead reads a single Run from persistent storage.
	RunRead(
		context.Context,
		api.KindVersionName,
		run.Selector,
	) (*run.Run, error)
	// RunWrite persists a single supplied Run to backend storage, Note that on
	// successful write, the newly-created or updated Run is returned.
	RunWrite(
		context.Context,
		run.Run,
	) (*run.Run, error)
	// RunQuery queries zero or more Runs from persistent storage.
	RunQuery(
		context.Context,
		query.Expression,
		...query.Option,
	) (*query.Result[*run.Run], error)
}
