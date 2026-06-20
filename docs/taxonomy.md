# Taxonomy

Data managed by `rxp` is organized in a uniform taxonomy.

This document describes how this data is categorized and identified.

## Categorization

Categorization refers to the ways in which things are grouped.

In `rxp`, data is categorized by *Type* and *Scope*.

Type refers to the class of thing. Types in `rxp` include System, Domain,
Kind and Object.

Scope refers to the *extent to which Names of instances of a Type of thing are
unique*. There are three scopes, shown here in decreasing order of breadth.

```mermaid
flowchart TD
    subgraph Global
        subgraph System
            subgraph Domain
```

Instances of a Type that is *globally*-scoped can *only* be identified by UUID
and never by Name.

Instances of a Type that is *system*-scoped can be identified by Name and
System.

Instances of a Type that is *domain-scoped* can be identified by Name and
Domain.

## Identification

Identification refers to the way in which a *single instance of a thing is
uniquely selected from a set of things of the same type*.

In `rxp`, you can select a single unique thing **by UUID** or **by Name**.

When selecting something by UUID, you supply a UUID value which is guaranteed
to be globally unique.

When selecting something by Name, you supply a human-readable string name along
with a System or Domain, depending on the thing's Scope.

For instance, if looking up a Domain by Name, you would supply the System in
which the Domain is found. If looking up an Object having a Kind that is scoped
to a Domain, you would supply the Domain.
