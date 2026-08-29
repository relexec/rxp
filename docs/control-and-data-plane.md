This document discusses the difference between a control plane and a data
plane.

Understanding these differences helps explain the design of `rxp` and
provides the foundation of best practices in architecting applications that use
`rxp`.

# Overview

Software systems provide functionality to users via a set of Application
Programming Interfaces (APIs).

These APIs expose two types of operations: *control plane* and *data plane*.

# Control plane

Users call a control plane operation to *manage some resource* that the
software system uses to model or group some system functionality.

A control plane resource is a logical entity or configuration element managed
by the control plane of a software system.

Managing a resource entails creating a record of the resource, updating the
desired state of the resource, deleting or archiving the resource, renaming
or transferring ownership of the resource, etc.

# Data plane

Users call a data plane operation to *use capabilities exposed by an instance
of a resource*.

Say your software system provided some cloud-based database functionality. You
might have control plane operations to create and delete a database. On the
other hand, querying records in tables in those databases would be a data plane
operation.

If your software system provided some workflow orchestration functionality, you
might have control plane operations to update the definition of those workflows
and data plane operations that executed a workflow.
