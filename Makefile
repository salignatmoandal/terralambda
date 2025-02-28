.PHONY: build test clean run install deploy logs rollback orchestrate

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

install: build
	@echo "Installing TerraLambda globally..."
	sudo cp bin/terralambda /usr/local/bin/terralambda
	sudo chmod +x /usr/local/bin/terralambda
	@echo "Installed! Now you can run 'terralambda' from anywhere."

deploy:
	@echo " Deploying AWS Lambda..."
	cd deployments/terraform && terraform init && terraform apply -auto-approve

logs:
	@echo " Fetching logs..."
	./bin/terralambda logs MyLambdaFunction

rollback:
	@echo "Rollback de la Lambda..."
	aws lambda update-alias --function-name MyLambdaFunction --name prod --function-version 1

orchestrate:
	@echo "Executing a Step Functions workflow..."
	./bin/terralambda orchestrate <state-machine-arn>