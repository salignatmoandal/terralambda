package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	awslambda "github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaInvoker represents an AWS Lambda function invoker
type LambdaInvoker struct {
	ctx    context.Context
	client *awslambda.Client
}

// NewInvoker creates a new instance of LambdaInvoker
func NewInvoker(ctx context.Context) (*LambdaInvoker, error) {
	// Load AWS configuration from environment or credentials file
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create AWS Lambda client
	client := awslambda.NewFromConfig(cfg)

	return &LambdaInvoker{
		ctx:    ctx,
		client: client,
	}, nil
}

// Invoke calls the specified AWS Lambda function with the given payload
func (i *LambdaInvoker) Invoke(functionName string, payload []byte) ([]byte, error) {
	// Prepare the input parameters for Lambda invocation
	input := &awslambda.InvokeInput{
		FunctionName:   &functionName,
		Payload:        payload,
		InvocationType: "RequestResponse",
	}

	// Call the Lambda function
	resp, err := i.client.Invoke(i.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke Lambda function %s: %w", functionName, err)
	}

	// Check if the function returned an error
	if resp.FunctionError != nil {
		return nil, fmt.Errorf("Lambda function %s returned error: %s", functionName, *resp.FunctionError)
	}

	return resp.Payload, nil
}
