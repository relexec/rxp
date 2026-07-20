# Identifiers

*Identifiers* are strings that can be used to uniquely identify some thing that
is managed by `rxp`.

We differentiate between machine-readable and human-readable identifiers,
calling the former things `UUIDs` and the latter `Names`.

## UUIDs

A `UUID` is a *globally-unique string identifier*.

All things managed by `rxp` have a `UUID()` method that returns the thing's
globally-unique identifer:

## Names

A `Name` is a *human-readable string identifier*. All things managed by `rxp`
can have a `Name`.

The scope of uniqueness of a thing's `Name` varies.

A `System`'s `Name` is optional and therefore is not guaranteed to be unique. A
`System`'s `UUID` is its only identifier.

A `Kind`'s `Name` is guaranteed to be unique within the scope of the `Kind`'s
`System`.

However, conventionally `Kind`s for types that are intended for public APIs are
globally-unique and conventionally are valid DNS domain names (e.g.
`runnable.t2.temporal.io`).

A `Domain`'s `Name` is guaranteed to be unique within the scope of the
`Domain`'s `System`.

## Object names

An `Object`'s `Name` is guaranteed to be unique within the `Scope`
associated with the `Object`'s `Kind`.

### System-qualified names

If the `Scope` associated with an `Object`'s `Kind` is
`ScopeSystem`, `rxp` guarantees the Object's `Name` is unique within the
`Object`'s `System` and `Kind`.

Names for these Objects are called *system-qualified names*.

### Domain-qualified names

If the `Scope` associated with an `Object`'s `Kind` is
`ScopeDomain`, `rxp` guarantees the Object's `Name` is unique within the
`Object`'s `System`, `Kind` and `Domain`.

Names for these Objects are called *domain-qualified names*.
