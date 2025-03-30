# Task Manager

A task management system developed in Go, supporting multiple API integration methods.

> **Note**: This project is developed in collaboration with Cursor IDE for learning purposes, exploring best practices in Go development and software architecture.

## Directory Structure

```
.
├── README.md
├── go.mod
├── go.sum
└── internal/
    ├── domain/             # Domain models and core interface definitions
    │   └── task.go         # Task-related domain models and interfaces
    │
    ├── repository/         # Data storage layer implementation
    │   └── memory/         # In-memory storage implementation
    │       └── task.go
    │
    ├── service/           # Business logic layer
    │   └── task/          # Task service implementation
    │       └── service.go
    │
    └── delivery/          # API transport layer
        ├── rest/          # RESTful API implementation
        │   └── handler.go
        │
        ├── graphql/       # GraphQL API implementation
        │   └── handler.go
        │
        └── grpc/         # gRPC API implementation
            └── handler.go

└── pkg/                  # Reusable packages
    └── validator/        # Validator implementation
        ├── task.go       # Task-related validation rules
        └── validator.go  # Common validation logic

```

## Architecture Overview

- **Domain Layer**: Defines core business models and interfaces
- **Repository Layer**: Implements data persistence
- **Service Layer**: Implements core business logic
- **Delivery Layer**: Provides multiple API access methods
  - RESTful API
  - GraphQL API
  - gRPC API

## Design Principles

This project follows SOLID principles:

1. **Single Responsibility Principle (SRP)**
   - Each package has a clear, single responsibility
   - Clear separation of responsibilities between layers

2. **Open-Closed Principle (OCP)**
   - Extensions implemented through interface definitions
   - New features can be added without modifying existing code

3. **Liskov Substitution Principle (LSP)**
   - All implementations can substitute their interfaces
   - Ensures system testability

4. **Interface Segregation Principle (ISP)**
   - Interfaces are concise and focused
   - Clients only depend on interfaces they need

5. **Dependency Inversion Principle (DIP)**
   - High-level modules don't depend on low-level modules
   - All depend on abstract interfaces

## Development Plan

- [x] Basic infrastructure setup
- [x] Core domain model implementation
- [x] Storage layer implementation
- [x] RESTful API implementation
- [ ] GraphQL API implementation
- [ ] gRPC API implementation
- [ ] Unit test coverage
- [ ] Performance testing and optimization 