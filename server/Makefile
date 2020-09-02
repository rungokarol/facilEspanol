PHONY=clean build run

all: build

clean:
	rm -fr bin/
	rm -fr mocks/

build:
	go build -o bin/server

run: build
	./bin/server

mocks:
	./scripts/generate_mocks.sh

test: mocks
	go test ./...

coverage: mocks
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
