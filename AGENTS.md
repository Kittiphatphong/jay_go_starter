# Agent Guidelines for Go Starter Project

## Build, Test, and Run Commands

```bash
# Build the application
go build -o bin/app .

# Run the application
go run .

# Run all tests
go test ./...

# Run a single test
go test -run TestFunctionName ./package

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Docker build
docker build -t go-starter .
```

## Project Architecture

This is a Go web API using **Fiber framework** with layered architecture:

- **controllers/**: HTTP handlers, request/response handling
- **services/**: Business logic layer
- **repositories/**: Data access layer (GORM)
- **routes/**: Route definitions and grouping
- **errs/**: Custom error types
- **logs/**: Zap logger wrapper
- **trails/**: Utility functions
- **config/**: Configuration management (Viper)
- **database/**: Database connections (Postgres, Redis)

## Code Style Guidelines

### Imports
Order imports as follows (separated by blank lines):
1. Standard library packages
2. Third-party packages
3. Local project packages (e.g., `go_starter/...`)

### Naming Conventions
- **Interfaces**: PascalCase (e.g., `Controller`, `Service`, `Repository`)
- **Structs**: lowercase for implementations (e.g., `controller`, `service`)
- **Constructor functions**: `New` + InterfaceName (e.g., `NewController`, `NewService`)
- **Packages**: lowercase, singular (e.g., `controllers`, `services`)
- **Variables**: camelCase for unexported, PascalCase for exported

### Struct Pattern
```go
type Service interface {
    MethodName() error
}

type service struct {
    repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
    return &service{repository: repository}
}
```

### Error Handling
- Use custom `AppError` from `errs` package for HTTP errors
- Wrap errors appropriately with context
- Log errors using `logs.Error()` before returning
- HTTP status codes: use `errs.ErrorBadRequest()`, `errs.ErrorInternalServerError()`, etc.

### HTTP Responses
Use helper functions from `controllers/handle.go`:
- `NewSuccessResponse(ctx, data)` - 200 OK with data
- `NewSuccessMsg(ctx, msg)` - 200 OK with message
- `NewCreateSuccessResponse(ctx, data)` - 201 Created
- `NewErrorResponses(ctx, err)` - Error response with proper status
- `NewErrorValidate(ctx, data)` - 422 validation errors

### Logging
- Use `logs.Info()`, `logs.Error()`, `logs.Debug()` from `logs` package
- Pass errors directly to `logs.Error()`

### Database
- Use GORM for database operations
- Repository pattern for data access
- Use `db.AutoMigrate()` for schema migrations in development
- Use `db.Migrator().DropTable()` before migrations when needed

### Configuration
- Use `config.Env("key")` for required config values
- Use `config.GetEnv("key", "default")` for optional values with defaults
- Configuration is loaded from `config.yaml`

## Common Patterns

### Route Registration
```go
func (r routesStruct) Install(app *fiber.App) {
    route := app.Group("prefix/", middleware)
    route.Get("endpoint", r.controller.Method)
}
```

### Fiber Context
- Always use `*fiber.Ctx` as first parameter
- Return `error` from all handlers
- Use `ctx.Next()` in middleware

## Testing
- No tests currently exist in this codebase
- When adding tests, follow standard Go testing patterns
- Place test files next to source files (e.g., `service_test.go` next to `service.go`)

## Dependencies

Key libraries used:
- `github.com/gofiber/fiber/v2` - Web framework
- `gorm.io/gorm` - ORM
- `github.com/spf13/viper` - Configuration
- `go.uber.org/zap` - Logging
- `github.com/go-playground/validator/v10` - Validation

## Environment

- Go version: 1.17+
- Module name: `go_starter`
- Default port: 9000
- Timezone: Asia/Bangkok
