# go-enum-linter

Go Enum Linter and Styles

## Problem Statement

In Go, enums are typically implemented using type aliases or type definitions with constant values, like this:

```go
package task
type Status int
const (
    Done Status = iota
    InProgress
    ToDo
)
```

However, this implementation is not truly type-safe. The Go compiler allows several unsafe operations that can break enum semantics:

```go
package main
type MyTask struct {
    status task.Status
}
const AnotherTaskStatus task.Status = 1111 // Arbitrary values allowed
func main() {
    var taskStatusA task.Status  // Zero-value initialization
    taskStatusB := 238             
    printStatus(taskStatusB)
    printStatus(238)             // Literal values accepted as parameters
    _ = MyTask{
        status: 348123893,       // Any literal value accepted in structs
    }
}

func printStatus(taskStatus task.Status) {
    fmt.Println(taskStatus)
}
```

## Solutions: Type-Safe Enum Guidelines

`go-enum-linter` enforces type-safe enum patterns through the following rules:

1. Enum definitions must be in files with the `.enum.go` suffix
2. Only within `.enum.go` files:
   - Define enum types (via type aliases or definitions)
   - Define enum constants / vars
   - Assign literal values to enum types
   - Define constructor functions for raw value conversion (if needed)

```go
// Constructor for safe conversion from raw values
func NewStatus(value int) (Status, error) {
    s := Status(value)
    switch s {
    case Done, InProgress, ToDo:
        return s, nil
    default:
        return Status(0), fmt.Errorf("invalid status value: %d", value)
    }
}
```

3. Outside of `.enum.go` files, the following are prohibited:
   - Defining new enum values
   - Zero-value initialization of enum variables
   - Direct integer assignments to enum types
   - Using integer literals where enum types are expected
   - Converting raw values to enum types without using constructors defined in `.enum.go`

##  Linter for the Guidelines

`go-enum-linter` helps you enforce these type-safe enum guidelines in your codebase.

### Installation: with `golangci-lint` (recommended)

The linter is available as a plugin for golangci-lint. First, ensure you have golangci-lint installed:

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### Usage

1. Enable the linter in your `.golangci.yml`:

```yaml
linters:
  enable:
    - go-enum-linter

linters-settings:
  go-enum-linter:
    # configuration options if any
```

2. Run golangci-lint:

```bash
golangci-lint run
```

The linter will report violations of the enum guidelines, such as:

- Invalid enum value assignments outside of `.enum.go` files
- Direct integer conversions without constructor functions
- Zero-value initializations
- And other violations of the guidelines above

### Installation manual

TODO: provide manual CLI installation