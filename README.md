# Go Lang

#### Setting up environment

-  Install latest Go Lang via `brew install go`
- `export GOPATH=/path/to/this/workspace`
- To build any given module in `src`, run `go install [module-name]`

e.g. `go install 01_hello_world`

- Run `./bin/[module-name]` to execute

#### Go Format

To format a module according to Go Lang style, you can run
`go fmt [module]` to do so.

e.g. `go fmt 01_hello_world`