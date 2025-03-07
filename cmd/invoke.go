package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/salignatmoandal/terralambda/internal/lambda"
	"github.com/spf13/cobra"
)

// GetInvokeCmd returns the invoke command
func GetInvokeCmd() *cobra.Command {
	var (
		payloadFlag string
		regionFlag  string
	)

	cmd := &cobra.Command{
		Use:   "invoke [function-name]",
		Short: "Invoke an AWS Lambda function",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if payloadFlag == "" {
				payloadFlag = "{}"
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			functionName := args[0]

			// Vérifier si le payload est un JSON valide
			var jsonTest map[string]interface{}
			if err := json.Unmarshal([]byte(payloadFlag), &jsonTest); err != nil {
				return fmt.Errorf("invalid JSON payload: %v", err)
			}

			payload := []byte(payloadFlag)

			// Créer un nouvel invoker
			invoker, err := lambda.NewInvoker(ctx, regionFlag)
			if err != nil {
				return fmt.Errorf("error creating invoker: %v", err)
			}

			// Exécuter l'invocation de la Lambda
			response, err := invoker.Invoke(functionName, payload)
			if err != nil {
				return fmt.Errorf("error invoking function: %v", err)
			}

			fmt.Printf("Response: %s\n", string(response))
			return nil
		},
	}

	// Important : définir les flags avant d'ajouter la commande
	cmd.Flags().StringVarP(&payloadFlag, "payload", "p", "{}", "JSON payload to send to the Lambda function")
	cmd.Flags().StringVarP(&regionFlag, "region", "r", "us-east-1", "AWS region for the Lambda function")

	return cmd
}
