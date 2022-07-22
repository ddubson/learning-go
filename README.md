# 🎓 Learning Go

## 🍦 The Fundamentals

- **[Variables are mutable by default](src/02_variables_and_types/mutable_var.go)** - when a variable is declared with
  the `var` keyword, it is mutable, and can change during its own lifetime. This also applies to the variables declared
  with _short variable declaration operator_
- **Default size of types such as int and float are based on the architecture of the machine running the Go compiler** -
  for example an `int` on a 64-bit machine, defaults to `int64`, or `int32` on a 32-bit machine.
- **[`type <name> struct` lets you construct a user-defined structure](src/02_variables_and_types/structs.go)**

## 🍨 Concepts & Constructs

- [**Zero value**](https://go.dev/tour/basics/12) - when variables are declared, they are initialized to their
  respective zero value (string -> "", bool -> false). The use of keyword `var` gives you a variable initialized with a
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

## Running the samples

e.g. `go run src/01_hello_world/main.go`

## Linting, Formatting, Maintenance, and Curation

To format a module according to Go Lang style, you can run
`go fmt [module]` to do so.

e.g. `go fmt 01_hello_world`

## Packaging & Distributing

TBD

## Go Resources

Essentials for the real world. Wheels don't need to be reinvented.

### CLI frameworks and tools

- [⭐️ Cobra CLI framework](https://github.com/spf13/cobra)

### Testing frameworks and tools

- [Ginkgo - Go BDD Test Framework](https://onsi.github.io/ginkgo/)
