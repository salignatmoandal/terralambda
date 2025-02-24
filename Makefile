.PHONY: build test clean run

build: clean
	go build -o bin/terralambda ./cmd/

test:
	go test -v ./...

clean:
	rm -rf bin/
	rm -f lambda.zip
	rm -f function.zip

run: build
	./bin/terralambda 