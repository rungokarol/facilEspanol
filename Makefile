PHONY=clean build run

clean:
		rm -fr bin/

build:
		go build -o bin/server

run: build
	./bin/server
