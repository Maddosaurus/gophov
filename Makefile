.DEFAULT_GOAL := build

build:
	go build -o bin/phovctl cmd/phovctl/main.go

clean:
	rm -rf bin/*