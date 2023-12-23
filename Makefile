.PHONY: all
all: build run

.PHONY: build
build:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		routes/**/*.proto
	go mod tidy
	go build .

.PHONY: run
run:
	./grpc-example
