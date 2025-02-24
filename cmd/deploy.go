package main

import (
	"fmt"
	"os"

	"github.com/salignatmoandal/terralambda/internal/lambda"
	"github.com/spf13/cobra"
)

// deployCmd is the command for deploying a new version of the Lambda function
// It compiles the lambda project, zips it, and then deploys the Terraform code
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new version of the Lambda function",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error retrieving working directory:", err)
			return
		}

		deployer := lambda.NewDeployer(ctx, workingDir)
		defer deployer.Cleanup()

		if err := deployer.Deploy("", ""); err != nil {
			fmt.Println("Error deploying:", err)
			return
		}

		fmt.Println("Deployment successful.")
	},
}
