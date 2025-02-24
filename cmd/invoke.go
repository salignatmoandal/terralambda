package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"

	"github.com/spf13/cobra"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke [lambda-name] [pa]",
	Short: "Invoke the Lambda function",
	Run: func(cmd *cobra.Command, args []string) {
		lambdaName := args[0]
		payload := args[1]

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}
		client := lambda.NewFromConfig(cfg)

		input := &lambda.InvokeInput{
			FunctionName:   &lambdaName,
			Payload:        []byte(payload),
			InvocationType: "RequestResponse",
		}

		resp, err := client.Invoke(context.TODO(), input)
		if err != nil {
			log.Fatalf("Failed to invoke Lambda function: %v", err)

		}
		fmt.Printf("Lambda function invoked successfully. Response: %s\n", string(resp.Payload))
		os.Exit(0)
	},
}
