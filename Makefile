.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run: tidy
	@go run main.go $(MOD)


.PHONY: test
test: tidy
	@pushd source && go test