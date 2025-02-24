.PHONY: build test clean run

build:
	go build -o bin/terralambda ./cmd/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	rm -f lambda.zip
	rm -f function.zip

run: build
	./bin/terralambda 