.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run: tidy
	@go run main.go $(MOD)