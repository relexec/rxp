package rxp

import (
	"context"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/query"
	"github.com/relexec/rxp/query/expression"
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
		expression.Expression,
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
		expression.Expression,
		...query.Option,
	) (*query.Result[*domain.Domain], error)
	// NamespaceRead reads a Namespace from persistent storage.
	NamespaceRead(
		context.Context,
		namespace.Selector,
	) (*namespace.Namespace, error)
	// NamespaceWrite atomically writes the supplied Namespace to persistent
	// storage.
	NamespaceWrite(
		context.Context,
		*namespace.Namespace,
	) error
	// NamespaceQuery queries zero or more Namespaces from persistent storage.
	NamespaceQuery(
		context.Context,
		expression.Expression,
		...query.Option,
	) (*query.Result[*namespace.Namespace], error)

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
		expression.Expression,
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
		expression.Expression,
		...query.Option,
	) (*query.Result[*kindversion.KindVersion], error)

	// ObjectRead reads a single object from persistent storage.
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
		expression.Expression,
		...query.Option,
	) (*query.Result[*object.Object], error)
}
