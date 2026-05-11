# Identifiers

*Identifiers* are strings that can be used to uniquely identify some thing that
is managed by `rxp`.

We differentiate between machine-readable and human-readable identifiers,
calling the former things `UUIDs` and the latter `Names`.

## UUIDs

A `UUID` is a *globally-unique string identifier*. Many things managed by `rxp`
can have a `UUID`:

* [`System`][system]
* [`Object`][object]

[system]: https://github.com/relexec/rxp/blob/main/docs/taxonomy.md#system
[object]: https://github.com/relexec/rxp/blob/main/docs/taxonomy.md#object

## Names

A `Name` is a *human-readable string identifier*. All things managed by `rxp`
can have a `Name`.

The scope of uniqueness of a thing's `Name` varies.

A `System`'s `Name` is optional and therefore is not guaranteed to be unique. A
`System`'s `UUID` is its only identifier.

A `Meta`'s `Kind` is a special `Name`. It is guaranteed to be unique within the
scope of the `Meta`'s `System`. However, conventionally `Kind`s for types that
are intended for public APIs are globally-unique and conventionally are valid
DNS domain names (e.g. `flow.temporal.io`).

A `Domain`'s `Name` is guaranteed to be unique within the scope of the
`Domain`'s `System`.

A `Namespace`'s `Name` is guaranteed to be unique within the scope of the
`Namespace`'s `Domain`.

## Object names

An `Object`'s `Name` is guaranteed to be unique within the `Namescope`
associated with the `Object`'s `Meta`.

### Globally-unique names

If the `Meta.Namescope` associated with an `Object`'s `KindVersion` is
`NamescopeGlobal`, the `Object`'s `Name` is expected to be globally-unique.

There is no way to guarantee global uniqueness of human-readable string names.

### System-qualified names

If the `Meta.Namescope` associated with an `Object`'s `KindVersion` is
`NamescopeSystem`, the `Object`'s `Name` `rxp` guarantees the name is unique
within the `Object`'s `System`.

### Kind-qualified names

If the `Meta.Namescope` associated with an `Object`'s `KindVersion` is
`NamescopeKind`, the `Object`'s `Name` `rxp` guarantees the name is unique
within the `Object`'s `System` and `Kind`.

### Domain-qualified names

If the `Meta.Namescope` associated with an `Object`'s `KindVersion` is
`NamescopeDomain`, the `Object`'s `Name` `rxp` guarantees the name is  unique
within the `Object`'s `System`, `Kind` and `Domain`.

### Namespace-qualified names

If the `Meta.Namescope` associated with an `Object`'s `KindVersion` is
`NamescopeNamespace`, the `Object`'s `Name` `rxp` guarantees the name is unique
within the `Object`'s `System`, `Kind`, `Domain` and `Namespace`.
