.PHONY: build test clean run download

build:
	go build -o build/myapp ./cmd/main.go

test:
	go test ./...

clean:
	rm -f build/myapp

run: build
	./build/myapp

download:
	go mod download