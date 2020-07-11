PHONY=clean build run

all: build

clean:
	rm -fr bin/

build:
	go build -o bin/server

run: build
	./bin/server
