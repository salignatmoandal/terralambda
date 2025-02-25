.PHONY: build test clean run

build:
	@echo "🚀 Building TerraLambda CLI..."
	mkdir -p bin
	go build -o bin/terralambda main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	rm -f lambda.zip
	rm -f function.zip

run: build
	./bin/terralambda