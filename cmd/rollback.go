package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/cobra"
)

// GetRollbackCmd returns the rollback command
func GetRollbackCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rollback [function-name] [version]",
		Short: "Rollback an AWS Lambda function to a previous version",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			functionName := args[0]
			version := args[1]

			fmt.Printf("Rolling back function: %s to version %s...\n", functionName, version)

			cfg, err := config.LoadDefaultConfig(context.TODO())
			if err != nil {
				log.Fatalf("Failed to load AWS config: %v", err)
			}

			client := lambda.NewFromConfig(cfg)

			// Update the alias "prod" to point to the previous version
			aliasName := "prod"
			input := &lambda.UpdateAliasInput{
				FunctionName:    aws.String(functionName),
				Name:            aws.String(aliasName),
				FunctionVersion: aws.String(version),
			}

			_, err = client.UpdateAlias(context.TODO(), input)
			if err != nil {
				log.Fatalf(" Failed to rollback Lambda function: %v", err)
			}

			fmt.Printf(" Successfully rolled back %s to version %s\n", functionName, version)
		},
	}
}
