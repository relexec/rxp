# Taxonomy

Data managed by `rxp` is uniformly organized in a common taxonomy. This
document describes how this data is defined, categorized, identified and named.

Briefly, a `Kind` identifies a *type* of a thing that is managed by `rxp`. A
`Kind` has a [`Name`][kindname] and a [`Namescope`](#namescope).

`Kinds` always have a [`System`](#system) identifier. System identifiers are
globally-unique.

[kindname]: https://github.com/relexec/rxp/blob/b5c989b0a587961dfaecf441c84dff58452fcbff/types/kind.go#L12-L60

```mermaid
erDiagram
    System ||--|{ Kind : "knows about"
    System {
        string **uuid**
        string name
    }
    Kind {
        string **system**
        string **name**
        int **namescope**
    }
```

A `KindVersion` is a **string** that uniquely identifies a *type and version*
of a thing that is managed by `rxp`.

A `Meta` contains the definition for a `KindVersion`. This definition includes
a `Schema` that defines the fields that comprise desired state for things of
that `KindVersion`.

```mermaid
erDiagram
    System ||--|{ Kind : "knows about"
    Kind ||--|{ Meta : "has a"
    System {
        string **uuid**
        string name
    }
    Kind {
        string **system**
        string **name**
        int **namescope**
    }
    Meta{
        string **kind**
        string **version**
        string **schema**
    }
```

An `Object` is an *instance* of a `KindVersion`.

`Objects` always have a `System` identifier.

`Objects` always have a UUID globally-unique identifier.

`Objects` always have a Name. An `Object`'s Name is unique within the
`Namescope` associated with the `Kind`.

If that `Namescope` is `NamescopeNamespace` or `NamescopeDomain`, the `Object`
is guaranteed to have a [`Domain`](#domain). If that `Namescope` is
`NamescopeNamespace`, the `Object` is guaranteed to have a
[`Namespace`](#namespace).

`Objects` may have zero or more `Labels` associated with them. `Labels` are
structures with a `Key` and optional `Value` that can be used to categorize
`Objects`.

```mermaid
erDiagram
    System ||--|{ Meta : "knows about"
    System ||--|{ Domain : "knows about"
    System ||--|{ Object : "knows about"
    Kind ||--|{ Meta : "has a"
    Meta ||--|{ Object : "instance of"
    Domain ||--|{ Namespace : "may have"
    Domain ||--|{ Object : "may have"
    Namespace ||--|{ Object : "may have"
    System {
        string **uuid**
        string name
    }
    Kind {
        string **name**
        int **namescope**
    }
    Meta{
        string **kind**
        string **version**
        string **schema**
    }
    Domain{
        string **system**
        string **name**
    }
    Namespace{
        string **domain**
        string **name**
    }
    Object{
        string **system**
        string **meta**
        string **uuid**
        string **name**
        string domain
        string namespace
        array labels
    }
```

## `System`

`System` represents the boundaries of an `rxp` system installation.

`System` has the following methods:

* `UUID()`: returns the globally-unique identifier.
* `Name()`: returns the optional human-readable name.

## `Domain`

`Domain` is a specialized string containing a top-level division or partition
of things managed by `rxp`.

A valid `Domain` is a DNS-formatted (RFC 1035-compliant) name less than 254
characters.

A `Domain` must be unique within the scope of the `rxp` system installation.

## `Namespace`

`Namespace` describes a logical division within a `Domain`.

A `Namespace` is typically used to segregate data by tenancy boundaries.

A valid `Namespace` is a DNS-formatted (RFC 1035-compliant) name.

Note that unlike RFC 1035, there is no 253 character size limit on `Namespace`
string length.

A `Namespace` must be unique within its containing `Domain`.

## `Kind`

[`Kind`][kind] is a specialized string containing the *type* of an `Object`.

A valid `Kind` is a DNS-formatted (RFC 1035-compliant) name of the type of
`Object`, e.g.  `flow.temporal.io`.

Conventionally, a `Kind` is specified as a singular, not plural, noun. So,
`flow`, not `flows`.

Furthermore, a `Kind` is conventionally all lower-cased, with dots separating
coarser-grained categories/groups. So, `flow.temporal.io`, not
`TemporalFlow`.

You can use only alphanumeric characters and hyphens in the `Kind` name parts,
separated by periods. Furthermore, the first character of the `Kind` must be a
letter or number, not a hyphen or period.

> Note that unlike RFC 1035, there is no 253 character size limit on the
> `Kind` string length.

A `Kind` must be unique within the scope of the `rxp` system installation,
however for any `Kind` that is intended to be used across multiple `rxp` system
installations, the `Kind` should be globally-unique.

[kind]: https://github.com/relexec/rxp/blob/main/types/kind.go

## `KindVersion`

[`KindVersion`][kindversion] is a specialized string that contains the `Kind`
and optionally a SemVer version string that uniquely identifies the exact type
of an `Object`.

A `KindVersion` string has the format `<kind>[@<version>]`, where `<kind>` is a
valid `Kind` and the optional `<version>` component must be a valid SemVer
version string.

> Note that a valid SemVer version string does *not* contain a `v` prefix.

[kindversion]: https://github.com/relexec/rxp/blob/main/types/kindversion.go

## `Namescope`

`Namescope` refers to the uniqueness constraint applied to the name of some
thing managed by `rxp`.

There are three `Namescope` values, listed here in order of specificity, from
the narrowest to broadest specificity.

* `NamescopeNamespace`: name is unique within the scope of the `Object`'s
  `System`, `Kind`, `Domain`, and `Namespace`.
* `NamescopeDomain`: name is unique within the scope of the `Object`'s
  `System`, Kind` and `Domain`.
* `NamescopeSystem`: name is unique within the scope of the `Object`'s `System`
  and `Kind`.

## `Object`

[`Object`][object] describes an *instance* of something whose lifecycle is
controlled by `rxp`.

An `Object`'s lifecycle encompasses its creation, mutation and deletion.

`Object` has the following methods:

* `System()`: returns the `System` to which the the `Meta` is known.
* `KindVersion()`: returns a unique identifier for the type and version of the
  Object.
* `UUID()`: returns the globally-unique identifier.
* `Domain()`: returns the optional `Domain`.
* `Namespace()`: returns the optional intra-`Domain` `Namespace`.
* `Name()`: returns the human-readable name.
* `Labels()`: returns the optional collection of `Label`s.
* `Generation()`: returns the number of times the `Object`'s desired state has
  changed.
* `Spec()`: returns the desired state.

When an `Object` is read, it will always have a non-zero `Generation` value.
The `Generation` represents the number of times that the desired state of the
`Object` (its `Spec` has been mutated).

[object]: https://github.com/relexec/rxp/blob/main/types/object.go

## `Meta`

`Meta` contains metadata about a versioned type of `Object`.

`Meta` has the following methods:

* `System()`: returns the `System` to which the the `Meta` is known.
* `KindVersion()`: returns the `KindVersion`
* `Version()`: returns the [`semver.Version`][semver-version] struct indicating
  the Semantic Version of the `Kind` of `Object` the `Meta` defines.
* `Namescope()`: returns the `Namescope` uniqueness constraint.
* `Schema()`: returns the [jsonschema.Schema][jsonschema-schema] describing the
  field composition of desired state.
* `SchemaJSON()`: returns a string representation of the `Schema`

When the definition of a `Kind` of `Object` changes, the `Version` is
incremented, allowing for the controlled evolution of the schema and definition
of a `Kind`.

[semver-version]: https://pkg.go.dev/github.com/Masterminds/semver/v3#Version
[jsonschema-schema]: https://github.com/google/jsonschema-go/blob/main/jsonschema/schema.go

## `Spec`

`Spec` represents the *desired state* of an `Object`.

The fields that comprise a `Spec` are defined in the `Meta`'s `Schema`.
