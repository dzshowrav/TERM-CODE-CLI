---
name: clean-architecture
description: Clean Architecture in Go with domain-driven design, dependency inversion, and layered separation of concerns
license: MIT
compatibility: opencode
metadata:
  audience: developers
  pattern: architecture
---

## Layer Structure
```
internal/
  domain/       # Enterprise business rules (entities, value objects, interfaces)
  application/  # Use case / application services
  adapters/     # Interface adapters (controllers, presenters, gateways)
  infrastructure/ # Frameworks, drivers, external tools
```

## Dependency Rule
- Dependencies point inward: domain has zero dependencies
- Application depends on domain
- Adapters depend on application (via interfaces)
- Infrastructure implements adapter interfaces
- Use dependency injection (constructor injection) at composition root

## Domain Layer
- Entities: core business objects with identity
- Value Objects: immutable, no identity, equality by value
- Domain Events: captured state changes within the domain
- Repository interfaces: defined by domain needs, not data model
- No framework imports, no database concerns, no HTTP concerns

## Application Layer
- Use cases: single-responsibility application operations
- Input/Output ports: DTOs or parameter objects
- Orchestration: coordinate domain objects and infrastructure
- Transaction boundaries: wrap use cases in units of work

## Interface Adapters
- HTTP handlers: parse request, call use case, format response
- CLI commands: parse flags, call use case, render output
- Presenters: format domain data for external consumption

## Testing
- Domain: pure unit tests, no mocks
- Application: test with interface fakes
- Adapters: integration tests with real infrastructure
- Architecture tests: enforce dependency rules with `aspecttest`
