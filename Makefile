.PHONY: run

run:
	export $$(cat .env) && \
	go run cmd/main.go

tidy:
	go mod tidy
