package object

import (
	"encoding/json"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/system"
)

// Object is an *instance* of a KindVersion.
//
// Each Object has a UUID globally-unique identifier.
//
// Objects have a `Name`. An Object's `Name` is unique within the
// Scope associated with the Object's Kind.
//
// If that Scope is `ScopeDomain`, the Object will have a Domain.
//
// Objects may have zero or more `Labels` associated with them. `Labels` are
// structures with a `Key` and optional `Value` that can be used to categorize
// Objects and filter them in query operations.
type Object struct {
	// kindVersionName is the kind and version identifier for the type of
	// Object.
	kindVersionName api.KindVersionName
	// system contains the system identifier for the Object.
	system *api.System
	// uuid is the globally-unique string identifier.
	uuid string
	// domain is the optional Domain.
	domain *api.Domain
	// name is the Name.
	name string
	// labels is the collection of Labels.
	labels api.Labels
	// generation contains the generation of the Object's desired state.
	generation api.Generation
	// spec contains the Object's desired state encoded as a JSON string.
	spec string
}

// KindName returns the DNS-formatted name of the Kind of Object, e.g.
// `flow.temporal.io`.
func (o Object) KindName() api.KindName {
	return o.kindVersionName.Kind()
}

// KindVersionName returns the KindVersionName of the Object. This string
// uniquely identifies the type of an Object.
func (o Object) KindVersionName() api.KindVersionName {
	return o.kindVersionName
}

// SetKindVersionName sets the Object's kind version name which uniquely
// identifies the type of an Object.
func (o *Object) SetKindVersionName(kvn api.KindVersionName) {
	o.kindVersionName = kvn
}

// System returns the System of the Object.
func (o Object) System() *api.System {
	return o.system
}

// SetSystem sets the System of Object.
func (o *Object) SetSystem(system *api.System) {
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
func (o Object) Domain() *api.Domain {
	return o.domain
}

// SetDomain sets the Domain.
func (o *Object) SetDomain(domain *api.Domain) {
	o.domain = domain
}

// Name returns the name. The Scope of the Object's Kind is used to determine
// whether the name is unique globally, within a Kind + System or within a Kind
// + Domain.
func (o Object) Name() string {
	return o.name
}

// SetName sets the name.
func (o *Object) SetName(name string) {
	o.name = name
}

// Labels returns the collection of Labels.
func (o Object) Labels() api.Labels {
	return o.labels
}

// SetLabels sets the collection of Labels.
func (o *Object) SetLabels(labels api.Labels) {
	o.labels = labels
}

// Generation returns the Object's Generation, which represents the number of
// mutations to the Object's desired state.
func (o Object) Generation() api.Generation {
	return o.generation
}

// SetGeneration sets the Object's Generation, which represents the number of
// mutations to the Object's desired state.
func (o *Object) SetGeneration(generation api.Generation) {
	o.generation = generation
}

// Spec returns the Object's desired state as a JSON-encoded string.
func (o Object) Spec() string {
	return o.spec
}

// SetSpec sets the Object's desired state as a JSON-encoded string.
func (o *Object) SetSpec(spec string) {
	o.spec = spec
}

type jsonObject struct {
	KindVersionName string     `json:"kind_version_name"`
	System          string     `json:"system"`
	UUID            string     `json:"uuid"`
	Domain          string     `json:"domain,omitempty"`
	Name            string     `json:"name"`
	Labels          api.Labels `json:"labels,omitempty"`
	Generation      int        `json:"generation"`
	Spec            string     `json:"spec"`
}

// MarshalJSON serializes the Object to a JSON bytestring.
func (o Object) MarshalJSON() ([]byte, error) {
	jo := jsonObject{
		KindVersionName: string(o.kindVersionName),
		UUID:            o.uuid,
		Name:            o.name,
		Labels:          o.labels,
		Generation:      int(o.generation),
		Spec:            o.spec,
	}
	if o.system != nil {
		jo.System = o.system.UUID()
	}
	if o.domain != nil {
		jo.Domain = string(o.domain.Name())
	}
	return json.Marshal(&jo)
}

// UnmarshalJSON constructs the Object from a JSON bytestring.
func (o *Object) UnmarshalJSON(text []byte) error {
	var jo jsonObject
	if err := json.Unmarshal(text, &jo); err != nil {
		return err
	}
	o.kindVersionName = api.KindVersionName(jo.KindVersionName)
	o.uuid = jo.UUID
	o.name = jo.Name
	o.labels = jo.Labels
	o.generation = api.Generation(jo.Generation)
	o.spec = jo.Spec
	if jo.System != "" {
		o.system = system.New(
			system.WithUUID(jo.System),
		)
	}
	if jo.Domain != "" {
		o.domain = domain.New(
			domain.WithSystem(o.system),
			domain.WithName(api.DomainName(jo.Domain)),
		)
	}
	return nil
}

// Clone returns a copy of the Object.
func (o Object) Clone() Object {
	return o
}
