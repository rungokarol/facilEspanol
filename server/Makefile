PHONY=clean build run

all: build

clean:
	rm -fr bin/

build:
	go build -o bin/server

run: build
	./bin/server

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
