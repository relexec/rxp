# Architecture

`rxp` provides low-level functionality for applications and services that
provide durable/reliable execution platforms.

## Repository structure

The [`types`][pkg-rxptypes] package contains interfaces and type definitions
referenced throughout the `rxp` library and associated `rxp` backend
implementations.

[pkg-rxptypes]: https://github.com/relexec/rxp/tree/main/types

## Design principles

* Well-documented code and plenty of example code

  The code itself should be well-documented with lots of usage examples.

* Interfaces should be consistent across modules

  Each module in the library should be structured in a consistent fashion, and
  the structs returned by various library functions should have consistent
  attribute and method names.

* Just the right amount of abstraction

  Developers should not need to wade through needless layers of abstraction to
  properly use the library. Interfaces should be small, with a minimal surface
  area driven by the consumer/caller of the interface's methods. No
  "AbstractFactoryBuilder" Java-esque stuff.

* Safety first, performance second

  The focus on the library should be on enforcing the safety and durability
  constraints that a Reliable Execution platform requires, not on raw
  performance. Performance optimization should come only after safety is
  guaranteed.

* Design for small to large scale

  The library should be capable of handling small (less than 10GB) to large
  (greater than 100TB) active data set sizes without rearchitecting.
  This means that structures managed by the library are designed to be
  partition-aware and advertise [name uniqueness constraints][#namescope].

## Out of scope

* Bindings or SDKs for programming languages other than Go.
