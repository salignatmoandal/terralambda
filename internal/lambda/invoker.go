package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaInvoker is responsible for invoking AWS Lambda functions
type LambdaInvoker struct {
	ctx    context.Context
	client *lambda.Client
}

// NewInvoker creates a new LambdaInvoker instance
func NewInvoker(ctx context.Context) (*LambdaInvoker, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := lambda.NewFromConfig(cfg)
	return &LambdaInvoker{
		ctx:    ctx,
		client: client,
	}, nil
}

// Invoke calls an AWS Lambda function
func (i *LambdaInvoker) Invoke(functionName string, payload []byte) ([]byte, error) {
	input := &lambda.InvokeInput{
		FunctionName:   &functionName,
		Payload:        payload,
		InvocationType: "RequestResponse",
	}

	resp, err := i.client.Invoke(i.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke Lambda function: %w", err)
	}

	return resp.Payload, nil
}
