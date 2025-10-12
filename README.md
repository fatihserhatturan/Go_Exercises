# Go Language Learning Repository

A comprehensive collection of Go language examples covering fundamental concepts, concurrency patterns, and advanced topics for learning and reference.

## Project Structure

### Basic Concepts
- **main**: Hello World starter program
- **variables**: Variable declarations with type reflection
- **data types**: Working with maps and collections
- **arrays & slices**: Array and slice operations
- **conditional expressions**: If-else statements and logic
- **loops**: For loops, range iterations, and accumulation
- **functions**: Single and multiple return values
- **structs**: Person struct definitions and initialization
- **pointers**: Pointer operations and memory addresses

### Interfaces & Modules
- **Interface**: Interface implementation with Greeter and Fareweller examples
- **error handling**: Error creation, wrapping, and file operations
- **defer & panic**: Defer statements, panic recovery, and safe division
- **function literals**: Anonymous functions, closures, and inline operations
- **variable scopes**: Global, local, block-level, and shadowing examples
- **packages and modules**: Custom mathutils package with arithmetic operations

### Asynchronous Programming
- **go routines**: Concurrent task execution with WaitGroups and timing
- **unbuffered channels**: Synchronous channel communication patterns
- **buffered channels**: Asynchronous channel operations with timeouts
- **context & monitoring**: Context-based cancellation and worker lifecycle management

## Getting Started

Navigate to any example directory and run:
```bash
go run main.go
```

For package examples with dependencies:
```bash
go mod download
go run main.go
```

## Requirements

- Go 1.21 or higher (examples use Go 1.22-1.24)
- Basic understanding of programming concepts

## Learning Path

1. **Basics**: Start with variables, data types, and control structures
2. **Functions & Structs**: Learn function patterns and data modeling
3. **Interfaces & Modules**: Understand abstraction and code organization
4. **Concurrency**: Master goroutines, channels, and context patterns

## Key Examples

**Concurrency Patterns**: WaitGroup synchronization, buffered/unbuffered channels, select statements, and context-based cancellation.

**Error Handling**: Multiple approaches including error returns, panic/recover, and error wrapping.

**Module Organization**: Custom package creation demonstrated with mathutils arithmetic library.
