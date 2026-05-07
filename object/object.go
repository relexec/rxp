package object

import (
	"encoding/json"

	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/system"
	"github.com/relexec/rxp/types"
)

// Object is a base struct from which all things that implement [types.Object]
// derive.
type Object struct {
	// kindVersion is the kind and version identifier for the type of Object.
	kindVersion types.KindVersion
	// system contains the system identifier for the Object.
	system types.System
	// uuid is the globally-unique string identifier.
	uuid string
	// domain is the optional Domain.
	domain types.Domain
	// namespace is the optional Namespace.
	namespace types.Namespace
	// name is the Name.
	name string
	// labels is the collection of Labels.
	labels types.Labels
	// generation contains the generation of the Object's desired state.
	generation types.Generation
	// spec contains the Object's desired state encoded as a JSON string.
	spec string
}

// Kind returns the DNS-formatted name of the Kind of Object, e.g.
// `flow.temporal.io`.
func (o Object) Kind() types.Kind {
	return o.kindVersion.Kind()
}

// KindVersion returns the KindVersion of the Object. This string uniquely
// identifies the type of an Object.
func (o Object) KindVersion() types.KindVersion {
	return o.kindVersion
}

// System returns the System of the Object.
func (o Object) System() types.System {
	return o.system
}

// SetSystem sets the System of Object.
func (o *Object) SetSystem(system types.System) {
	o.system = system
}

// UUID returns the globally-unique string identifier.
func (o Object) UUID() string {
	return o.uuid
}

// SetUUID sets the globally-unique string identifier.
func (o *Object) SetUUID(uuid string) {
	o.uuid = uuid
}

// Domain returns the optional Domain.
func (o Object) Domain() types.Domain {
	return o.domain
}

// SetDomain sets the Domain.
func (o *Object) SetDomain(domain types.Domain) {
	o.domain = domain
}

// Namespace returns optional Namespace.
func (o Object) Namespace() types.Namespace {
	return o.namespace
}

// SetNamespace sets the Namespace.
func (o *Object) SetNamespace(ns types.Namespace) {
	o.namespace = ns
}

// Name returns the name. NameScope can be used to determine whether the name
// is unique globally, or within a Kind + Domain or within a Kind + Domain +
// Namespace.
func (o Object) Name() string {
	return o.name
}

// SetName sets the name.
func (o *Object) SetName(name string) {
	o.name = name
}

// Labels returns the collection of Labels.
func (o Object) Labels() types.Labels {
	return o.labels
}

// SetLabels sets the collection of Labels.
func (o *Object) SetLabels(labels types.Labels) {
	o.labels = labels
}

// Generation returns the Object's Generation, which represents the number of
// mutations to the Object's desired state.
func (o Object) Generation() types.Generation {
	return o.generation
}

// Spec returns the Object's desired state as a JSON-encoded string.
func (o Object) Spec() string {
	return o.spec
}

type jsonObject struct {
	KindVersion string       `json:"kind_version"`
	System      string       `json:"system"`
	UUID        string       `json:"uuid"`
	Domain      string       `json:"domain,omitempty"`
	Namespace   string       `json:"namespace,omitempty"`
	Name        string       `json:"name"`
	Labels      types.Labels `json:"labels,omitempty"`
	Generation  int          `json:"generation"`
	Spec        string       `json:"spec"`
}

// MarshalJSON serializes the Object to a JSON bytestring.
func (o Object) MarshalJSON() ([]byte, error) {
	jo := jsonObject{
		KindVersion: string(o.kindVersion),
		UUID:        o.uuid,
		Name:        o.name,
		Labels:      o.labels,
		Generation:  int(o.generation),
		Spec:        o.spec,
	}
	if o.system != nil {
		jo.System = o.system.UUID()
	}
	if o.domain != nil {
		jo.Domain = string(o.domain.Name())
	}
	if o.namespace != nil {
		jo.Namespace = string(o.namespace.Name())
	}
	return json.Marshal(&jo)
}

// UnmarshalJSON constructs the Object from a JSON bytestring.
func (o *Object) UnmarshalJSON(text []byte) error {
	var jo jsonObject
	if err := json.Unmarshal(text, &jo); err != nil {
		return err
	}
	o.kindVersion = types.KindVersion(jo.KindVersion)
	o.uuid = jo.UUID
	o.name = jo.Name
	o.labels = jo.Labels
	o.generation = types.Generation(jo.Generation)
	o.spec = jo.Spec
	if jo.System != "" {
		o.system = system.New(
			system.WithUUID(jo.System),
		)
	}
	if jo.Domain != "" {
		o.domain = domain.New(
			domain.WithSystem(o.system),
			domain.WithName(types.DomainName(jo.Domain)),
		)
	}
	if jo.Namespace != "" {
		o.namespace = namespace.New(
			namespace.WithDomain(o.domain),
			namespace.WithName(types.NamespaceName(jo.Namespace)),
		)
	}
	return nil
}
