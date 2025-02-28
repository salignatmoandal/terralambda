package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaRollback handles AWS Lambda rollbacks
type LambdaRollback struct {
	ctx    context.Context
	client *lambda.Client
}

// NewRollback initializes a new LambdaRollback manager
func NewRollback(ctx context.Context) (*LambdaRollback, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := lambda.NewFromConfig(cfg)

	return &LambdaRollback{
		ctx:    ctx,
		client: client,
	}, nil
}

// Rollback reverts a Lambda function to a previous version
func (r *LambdaRollback) Rollback(functionName, version string) error {
	fmt.Printf("Rolling back function: %s to version %s...\n", functionName, version)

	aliasName := "prod"
	input := &lambda.UpdateAliasInput{
		FunctionName:    aws.String(functionName),
		Name:            aws.String(aliasName),
		FunctionVersion: aws.String(version),
	}

	_, err := r.client.UpdateAlias(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("Failed to rollback Lambda function: %v", err)
	}

	fmt.Printf("Successfully rolled back %s to version %s\n", functionName, version)
	return nil
}
