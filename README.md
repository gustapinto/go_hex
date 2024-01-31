# Go HEX

A simple project to train hexagonal architecture with golang, with some elements of domain-driven design.

## Main folders and packages

- `cmds`: the application modules, they act as the transport layer
- `internal/core`: the business logic of the project, act as the core of the architecture
- `internal/datasource`: the datasources of the project, they provide a implementation for the `internal/core` defined repositories
- `pkg`: the util packages, containg functions and structs that can be used by any layer of teh project and also by other projects
