package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/salignatmoandal/terralambda/internal/lambda"
	"github.com/spf13/cobra"
)

// GetDeployCmd returns the deploy command
func GetDeployCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a new version of the AWS Lambda function",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			// Get the current working directory
			workingDir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error retrieving working directory:", err)
				return
			}

			// Create a new Lambda deployer instance
			deployer := lambda.NewDeployer(ctx, workingDir)
			defer func() {
				if err := deployer.Cleanup(); err != nil {
					fmt.Printf("Error during cleanup: %v\n", err)
				}
			}()

			// Execute deployment
			if err := deployer.Deploy("", ""); err != nil {
				fmt.Println("Error deploying:", err)
				return
			}

			fmt.Println("✅ Deployment successful!")
		},
	}
}
