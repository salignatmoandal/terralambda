package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new version of the Lambda function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compiling lambda project...")
		if err := exec.Command("go", "build", "-o", "lambda", "main.go").Run(); err != nil {
			fmt.Println("Error compiling lambda project:", err)
			return
		}
		exec.Command("zip", "-r", "lambda.zip", "lambda").Run()

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
