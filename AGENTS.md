# AGENTS.md

This file contains guidelines for agentic coding assistants working on the moblang codebase.

## Build, Lint, and Test Commands

### Core Commands (User CLI)
- `mob run main.mob` - Compile and execute in one step
- `mob build main.mob` - Compile to native binary, creates `./main`
- `mob serve main.mob` - Start HTTP server from entry point
- `mob lint .` - Run native linter on all .mob files
- `mob version` - Show CLI version

### Development Commands
- `make build` - Build the mob CLI binary to `bin/mob`
- `make test` - Run all tests
- `make lint` - Run golangci-lint (if available)
- `make clean` - Remove build artifacts
- `make install` - Install mob to `~/.local/bin/mob`
- `make build-all` - Build for all platforms (linux-amd64, linux-arm64)

### Linter Checks
The native linter prioritizes:
1. Detection of circular imports
2. Detection of unused imports
3. All other code quality checks

## Code Style Guidelines

### File Extensions and Structure
- All source files use `.mob` extension
- Entry point is always `main.mob` in project root
- Application code resides in `src/` directory
- Recommended structure:
  ```
  /
  ├── main.mob
  └── src/
      ├── controllers/
      ├── services/
      ├── models/
      ├── jobs/
      └── commands/
  ```

### Syntax and Formatting
- Python-like syntax with indentation-based blocks
- 4 spaces for indentation (no tabs)
- No semicolons required
- Clean, readable, predictable code
- Minimize lines of code while maintaining clarity

### Type System
- Strongly typed language with compile-time type checking
- All types must be explicitly declared
- Type annotations follow Python/TypeScript style: `name: type`
- Example: `public name: string`

### Naming Conventions
- Classes: PascalCase (e.g., `User`, `HttpRequest`)
- Functions and methods: snake_case (e.g., `get_user`, `process_request`)
- Variables: snake_case (e.g., `user_id`, `response_body`)
- Constants: UPPER_SNAKE_CASE (e.g., `MAX_RETRY_COUNT`, `API_BASE_URL`)
- Private members: prefix with underscore if applicable (e.g., `_internal_method`)
- Files: snake_case (e.g., `user_controller.mob`, `http_service.mob`)

### Object-Oriented Programming
- All code must be inside classes (except entry point)
- Use `class` keyword for class definitions
- Use `extends` for inheritance (e.g., `class User extends Model`)
- Explicit access modifiers: `public`, `private`, `protected`
- Use `new` for instantiation: `user = new User()`
- Use `this` to reference instance members
- Property/method access with dot notation: `this.name`, `user.hello()`

### Import System
- Use `import` keyword for all imports (core and project modules)
- Import core modules: `import core.http`, `import core.json`
- Import project modules: `import app.services.user`, `import src.models.user`
- Never use circular imports (linter will catch this)
- Remove unused imports (linter will catch this)

### Error Handling
- Use native error handling mechanisms (to be specified)
- Always handle potential errors in production code
- Propagate errors appropriately in service layers
- Log errors appropriately for debugging

### Async and Concurrency
- Use native async/await for asynchronous operations
- Use green threads for concurrent operations
- Use parallel execution where appropriate
- Avoid blocking operations in async contexts

### Memory Management
- Rely on garbage collector for automatic memory management
- Avoid manual memory management
- Consider memory usage for long-running applications

### Code Organization
- Keep functions small and focused on single responsibility
- Prefer composition over complex inheritance
- Organize code by feature/domain
- Separate concerns: controllers for HTTP, services for business logic, models for data

### Documentation
- Code should be self-documenting through clear naming
- Add docstrings for non-obvious functions/classes
- Document public APIs and complex algorithms

## Language Philosophy
- Legibility is paramount
- Performance is not optional
- Native tools over external dependencies
- Less configuration, more execution
- Developer comes before framework

## Testing (Once Implemented)
- Tests should be in `tests/` or `__tests__/` directories
- Test files should end with `.test.mob` or `_test.mob`
- Use native testing framework
- Test file names: `user_test.mob`, `user.test.mob`

## Performance Considerations
- Code compiles to native binaries with performance equal or superior to Go
- Minimize allocations in hot paths
- Use efficient data structures and algorithms
- Profile before optimizing critical sections
