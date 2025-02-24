package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	return Response{Message: fmt.Sprintf("Hello, %s!", request.Name)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
