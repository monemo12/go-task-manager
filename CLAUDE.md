# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Development
```bash
# Run the application
go run main.go

# Build the application
go build -o task-manager main.go

# Install dependencies
go mod tidy

# Format code
go fmt ./...

# Lint code
go vet ./...
golangci-lint run  # (requires installation: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
```

### Testing
```bash
# Run tests (when implemented)
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/service/task/
```

## Architecture Overview

This is a **Clean Architecture** Go application implementing a task management system with multiple API delivery mechanisms.

### Layer Structure
- **Domain Layer** (`internal/domain/`): Core business models and interfaces
- **Service Layer** (`internal/service/`): Business logic implementation
- **Repository Layer** (`internal/repository/`): Data access layer
- **Delivery Layer** (`internal/delivery/`): API handlers (REST, GraphQL, gRPC)
- **Validation Package** (`pkg/validator/`): Reusable validation logic

### Key Architectural Patterns

#### Dependency Injection
All dependencies are manually injected in `main.go`:
```go
repo := repo_task_memory.NewRepository()
taskValidator := validator.NewTaskValidator()
notifier := &SimpleNotifier{}
taskService := service_task.NewService(repo, taskValidator, notifier)
```

#### Interface-Driven Design
All major components depend on interfaces defined in the domain layer:
- `TaskRepository` - Data access abstraction
- `TaskService` - Business logic abstraction  
- `TaskValidator` - Validation abstraction
- `TaskNotifier` - Notification abstraction

#### Repository Pattern
- Thread-safe in-memory implementation using `sync.RWMutex`
- Interface defined in domain, implementation in `internal/repository/task/memory/`
- Context propagation throughout all operations

### Domain Model

#### Core Types
- `Task`: Main entity with ID, Title, Description, Priority, Status, timestamps
- `Priority`: Enum (Low, Medium, High) with custom JSON marshaling
- `Status`: Enum (Todo, InProgress, Done) with custom JSON marshaling

#### Business Rules
- UUIDs for task identification
- Validation rules enforced through validator package
- Time-based operations (CreatedAt, UpdatedAt, DueDate)

### Validation System

Located in `pkg/validator/`, uses a composable rule-based approach:
- `BaseValidator` with rule composition
- `TaskValidator` implements domain-specific validation
- Extensible rule system (`NotEmptyRule`, `MaxLengthRule`, `FutureDateRule`)

### API Delivery

#### REST API (Implemented)
- Gin framework on port 8080
- Full CRUD operations: `GET|POST /tasks`, `GET|PUT|DELETE /tasks/:id`
- JSON serialization with custom marshaling for enums

#### GraphQL & gRPC (Planned)
- Handler skeletons exist in `internal/delivery/graphql/` and `internal/delivery/grpc/`
- Same service interface dependency pattern

### Error Handling

Custom error types in `pkg/errors/`:
- `AppError` interface for consistent error handling
- Predefined error codes and messages
- Structured error responses

### Development Notes

#### Code Style
- Mixed Chinese and English comments (domain layer uses Chinese)
- Standard Go conventions for package organization
- Interface names follow Go idioms

#### Current Limitations
- No test files exist yet
- In-memory storage only (no persistence)
- Hard-coded configuration values
- Basic logging implementation

#### Extension Points
- Add new delivery mechanisms by implementing handlers that depend on `domain.TaskService`
- Add new storage backends by implementing `domain.TaskRepository`
- Add new validation rules by implementing `validator.ValidationRule`
- Add new notification channels by implementing `domain.TaskNotifier`

### Testing Strategy

When implementing tests:
- Use table-driven tests for validation rules
- Mock interfaces for unit testing service layer
- Test repository concurrency safety
- Integration tests for API handlers

### Performance Considerations

- Repository uses RWMutex for concurrent read access
- Context cancellation should be respected in long-running operations
- Consider adding database connection pooling when implementing persistent storage