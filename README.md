# Learning Go

## 📚 Concepts

- [**Zero value**](https://go.dev/tour/basics/12) - when variables are declared, they are initialized to their
  respective zero value (string -> "", bool -> false). The use of keyword `var` gives you a variable initialized with a
  zero value.
- [**Short variable declaration operator**](https://go.dev/tour/basics/10) - when you use `:=` operator, Go infers the
  type being declared or throws a compilation error if it cannot infer the type.
- **Escape Analysis** - part of static code analysis done by the compiler - determines if value belongs on the stack or
  on the heap. This is to ensure that integrity is maintained as a value is accessed during its lifetime.

## 🏗 Setting up a Go environment

- Install latest Go Lang via `brew install go`
- `export GOPATH=/path/to/this/workspace`
- To build any given module in `src`, run `go install [module-name]`

e.g. `go install 01_hello_world`

- Run `./bin/[module-name]` to execute

e.g. `./bin/01_hello_world`

## Go Format

To format a module according to Go Lang style, you can run
`go fmt [module]` to do so.

e.g. `go fmt 01_hello_world`

## Go Resources

Go BDD Test Framework - https://onsi.github.io/ginkgo/