build:
	@go build -o gobank
run: build
	@./gobank

test:
	@go test -v ./...

