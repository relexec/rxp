# `rxp` - Reliable eXecution Primitives

[![Go Reference](https://pkg.go.dev/badge/github.com/relexec/rxp.svg)](https://pkg.go.dev/github.com/relexec/rxp)
[![Test Status](https://github.com/relexec/rxp/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/relexec/rxp/actions)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

`rxp` provides building blocks used to construct a Reliable Execution platform.

This source code repository contains the core `rxp` library. There are separate
source code repositories for `rxp` backend implementations -- for example, the
[`rxp-pg`][rxp-pg] repository contains the `rxp` implementation using
PostgreSQL as the primary persistence store.

At its core, a Reliable Execution platform must be able to:

* Guarantee uniqueness of names within some scope
* Safely evolve the definition of a thing
* Safely mutate desired state of a thing
* Safely archive managed things
* Provide auditability for managed things

[rxp-pg]: https://github.com/relexec/rxp-pg

## Name uniqueness

All things managed by `rxp` have a *name*.

A thing's name is a special attribute.

Because names are easier to remember than globally-unique string identifiers
like UUIDs, things are often looked up by their name. This is why names must be
treated with care by the system.

Human-readable names for various things in the `rxp` system are guaranteed to
be unique within a particular [`Namescope`][namescope].

[namescope]: https://github.com/relexec/rxp/blob/main/docs/taxonomy.md#namescope

### Renaming

Humans change their mind, sometimes often. Because of this, the ability to
rename something is a critical piece of functionality for long-lived systems.

`rxp` treats the action of renaming something as a special operation. When a
thing is renamed, `rxp` guarantees that the renaming of done in a safe, audited
and complete fashion. Partly, this means that renaming is isolated to its own
mutation event. `rxp` ensures that you *cannot modify the desired state of
something and change its name at the same time*.

## Safe definition evolution

As business requirements change, developers need to evolve the definition and
schema of structures used by their applications.

Reliable Execution platforms faciliate the execution of code that is designed
to potentially run over very long periods of time. `rxp` understands when the
definition of something changes and provides guardrails that prevent unsafe
changes to that definition from potentially breaking application clients.

## Safe desired state mutation

## Safe archival

Things managed by `rxp` do not get deleted. This is so that `rxp` can ensure
that the history of changes to some thing under its management are always
persisted in order to restore a particular version or generation of the thing
at some point in the future.

Instead of deletion, things managed by `rxp` can be *archived*. Archival marks
the thing as not being in the *active data set*, which means the thing will no
longer be returned from Read operations unless the caller indicates `rxp`
should include archived items in its lookup actions.

## Auditability
