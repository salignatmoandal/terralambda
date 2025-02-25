package cmd

import (
	"context"
	"fmt"

	"github.com/salignatmoandal/terralambda/internal/lambda"
	"github.com/spf13/cobra"
)

// GetInvokeCmd returns the invoke command
func GetInvokeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "invoke [function-name]",
		Short: "Invoke an AWS Lambda function",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			functionName := args[0]

			invoker, err := lambda.NewInvoker(ctx)
			if err != nil {
				fmt.Printf("Error creating invoker: %v\n", err)
				return
			}

			payload := []byte("{}") // Default payload
			response, err := invoker.Invoke(functionName, payload)
			if err != nil {
				fmt.Printf("Error invoking function: %v\n", err)
				return
			}

			fmt.Printf("Response: %s\n", string(response))
		},
	}
}
