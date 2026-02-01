# AGENTS.md - Agent Guide for AGG Repository

This document provides essential information for agentic coding agents working with the AGG (Amisgo Games) repository.

## Project Overview

AGG is a Go-based educational game platform featuring STEM games built with the amisgo framework. The project includes:
- Ball Sort Puzzle game
- Tower of Hanoi game  
- Ice Magic game

## Build and Development Commands

### Core Commands
```bash
# Run the application
make run                    # Starts the server at http://localhost:3000
go run -ldflags="-checklinkname=0" ./cmd

# Generate exports for ixgo
make gen                   # Exports hanoi package to ./internal/exported
ixgo export -outdir ./internal/exported ./pkg/export/hanoi

# Clean generated exports
make clean                 # Removes internal/exported directory

# Build
go build ./cmd

# Testing
go test ./...              # Run all tests
go test -v ./...           # Run tests with verbose output
go test -run TestName ./... # Run specific test
go test ./internal/games/... # Run tests for specific package
```

## Project Structure

```
agg/
├── cmd/main.go                 # Entry point
├── internal/
│   ├── app.go                  # Main application setup
│   ├── static/                 # Static assets
│   └── games/                  # Game implementations
│       ├── ball-sort/
│       ├── hanoi/
│       └── ice-magic/
├── pkg/                        # Shared packages
│   ├── base.go                 # Base game structure
│   ├── level.go                # Level management
│   └── export/                 # Export packages for ixgo
└── internal/exported/         # Generated exports (do not edit)
```

## Code Style Guidelines

### Import Organization
- Group imports in three sections: standard library, third-party packages, internal packages
- Use blank imports (`_`) for side effects
- Import paths use full module names (e.g., `github.com/zrcoder/agg/pkg`)

Example:
```go
import (
    "errors"
    "fmt"                    // Standard library
    
    "github.com/gorilla/websocket"
    "github.com/zrcoder/amisgo"  // Third-party
    
    "github.com/zrcoder/agg/internal/games/hanoi"
    "github.com/zrcoder/agg/pkg"  // Internal
)
```

### Naming Conventions
- **PascalCase** for exported types, functions, constants
- **camelCase** for unexported types, functions, variables
- **ALL_CAPS** for constants that are truly constant
- Use descriptive names (e.g., `HanoiCodeAction` not `hca`)
- Game-specific types should be concise but clear (e.g., `Pile`, `Disk`, `Bottle`, `Ball`)

### Type Definitions
- Embed common functionality using anonymous fields
- Use struct embedding for composition (e.g., `*Game` embedded in `Pile`)
- Define methods on receiver types that make semantic sense

### Error Handling
- Return errors explicitly, don't panic for expected errors
- Use descriptive error messages
- Handle errors at appropriate levels
- Use `errors.New()` for simple errors, `fmt.Errorf()` for formatted errors

### Constants
- Group related constants together
- Use meaningful names that describe their purpose
- Game-specific constants should use descriptive prefixes or be scoped to their package

### Function Organization
- Constructor functions use `New()` prefix
- Methods are organized by functionality and importance
- UI methods are grouped together
- Business logic methods are separated from UI logic

## Framework Conventions

### Amisgo Usage
- All games extend `*pkg.Base` for common functionality
- Use `amisgo.App` for application-level operations
- Components are built using the fluent interface pattern
- WebSocket integration is handled through the base package

### Level Management
- Implement `pkg.Level` interface for game levels
- Use `pkg.WithLevels()` option for setup
- Level data is typically stored in `levels/` subdirectories
- Support both simple and chapter-based level structures

### UI Patterns
- Use `comp.Service` for game UIs with WebSocket updates
- Implement `Main()` method for scene rendering
- Use forms for user interactions with proper submit handlers
- Follow the established pattern for level selection and reset forms

## Testing

- Test files use `_test.go` suffix
- Use table-driven tests where appropriate
- Test both happy path and error cases
- Mock external dependencies where needed
- No current test files exist - this is an area for improvement

## Development Workflow

1. Make changes to game logic or UI
2. Use `make run` to test locally at http://localhost:3000
3. If modifying export packages, run `make clean && make gen`
4. Test game functionality through the web interface
5. Ensure all games still function correctly

## Important Notes

- The `internal/exported/` directory is auto-generated - do not edit manually
- WebSocket connections are used for real-time UI updates
- The application runs on port 3000 by default
- Games share common patterns but have unique logic implementations
- The ixgo integration allows for dynamic code execution in some games

## Dependencies

- Go 1.23+
- amisgo framework for UI and web server
- ixgo for dynamic code execution
- Standard Go libraries for HTTP, WebSocket, and random generation