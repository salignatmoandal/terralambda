package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// deployCmd is the command for deploying a new version of the Lambda function
// It compiles the lambda project, zips it, and then deploys the Terraform code
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new version of the Lambda function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compiling lambda project...")

		// Compile the Go code for AWS Lambda
		if err := exec.Command("go", "build", "-o", "lambda", "main.go").Run(); err != nil {
			fmt.Println("Error compiling lambda project:", err)
			return
		}
		exec.Command("zip", "-r", "lambda.zip", "lambda").Run()

		// Execute the Terraform code
		fmt.Println("Deploying Terraform code...")
		tf := exec.Command("terraform", "apply", "-auto-approve")
		tf.Stdout = os.Stdout
		tf.Stderr = os.Stderr
		if err := tf.Run(); err != nil {
			fmt.Println("Error deploying Terraform code:", err)
			return
		}
		fmt.Println("Terraform code deployed successfully.")

	},
}
