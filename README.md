# 🎓 Learning Go

[![Learning Go](https://github.com/ddubson/learning-go/actions/workflows/go.yml/badge.svg)](https://github.com/ddubson/learning-go/actions/workflows/go.yml)

## 🍦 The Fundamentals

- **[Variables are mutable by default](source/2_variables_and_types.go)** - when a variable is declared with
  the `var` keyword, it is mutable, and can change during its own lifetime. This also applies to the variables declared
  with _short variable declaration operator_
- **Default size of types such as int and float are based on the architecture of the machine running the Go compiler** -
  for example an `int` on a 64-bit machine, defaults to `int64`, or `int32` on a 32-bit machine.
- **[`type <name> struct` lets you construct a user-defined structure](source/7_user_defined_types.go)**
- **Go is structured by modules and packages**
- Functions are by default **package-private**.
- To export to make available externally, **capitalize** the first character of the function name
- Unused variables produce a compile-time error - you may not declare and initialize a variable without later using the
  variable.

## 🍨 Concepts & Constructs

- [**Zero value**](https://go.dev/tour/basics/12) - when variables are declared, they are initialized to their
  respective zero value (string -> `""`, bool -> `false`). The use of keyword `var` gives you a variable initialized with a
  zero value. [[1]](https://go.dev/ref/spec#The_zero_value)
- [**Short variable declaration operator**](https://go.dev/tour/basics/10) - when you use `:=` operator, Go infers the
  type being declared or throws a compilation error if it cannot infer the type.
- **Escape Analysis** - part of static code analysis done by the compiler - determines if value belongs on the stack or
  on the heap. This is to ensure that integrity is maintained as a value is accessed during its lifetime.

## 🏗 Setting up a Go environment

### Installing Go

- [On a MacOS](https://go.dev/doc/install)

### Preparing the environment for local usage

```shell
mkdir ~/go

# For ZSH users
echo "GOPATH=/Users/$(whoami)/go" >> ~/.zprofile && zsh -l
```

## Running examples

> NOTE: On Windows, run the following with `go run <dir>` (sans *.go)

```shell
go run 01-variables-and-types/*.go
go run 02-conditionals/*.go
go run 03-loops/*.go
go run 04-functions/*.go
go run 05-errors/*.go
go run 06-user-defined-types/*.go
go run 07-concurrency/*.go
go run 08-interfaces/*.go
go test 09-testing/adder/*.go
go test 09-testing/poem/*.go
go run 10-builtin-libraries/*.go
go run 11-file-io/*.go
```

## Linting, Formatting, Maintenance, and Curation

### go mod tidy

> `go mod tidy` ensures that the `go.mod` file **matches the source code in the module**. It adds any missing module
> requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules
> that don’t provide any relevant packages. It also adds any missing entries to `go.sum` and removes unnecessary entries.

```bash
go mod tidy
```

### go fmt

To format a module according to Go Lang style, you can run
`go fmt [module]` to do so.

e.g. `go fmt 01_hello_world`

## Go Resources

Essentials for the real world. Wheels don't need to be reinvented.

### CLI frameworks and tools

- [⭐️ Cobra CLI framework](https://github.com/spf13/cobra)

### Testing frameworks and tools

- [Ginkgo - Go BDD Test Framework](https://onsi.github.io/ginkgo/)
