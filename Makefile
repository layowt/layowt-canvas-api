.PHONY: build test clean run

build:
	go build -o build/myapp ./cmd/main.go

test:
	go test ./...

clean:
	rm -f build/myapp

run: build
	./myapp