package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request defines the input structure
type Request struct {
	Name string `json:"name"`
}

// Response defines the output structure
type Response struct {
	Message string `json:"message"`
}

// HandleRequest is the AWS Lambda function handler
func HandleRequest(ctx context.Context, request Request) (Response, error) {
	if request.Name == "" {
		return Response{}, fmt.Errorf("name field is required")
	}

	return Response{Message: fmt.Sprintf("Hello, %s!", request.Name)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
