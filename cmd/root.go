package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the main command for the TerraLambda CLI
var RootCmd = &cobra.Command{
	Use:   "terralambda",
	Short: "TerraLambda is a CLI tool to manage AWS Lambda functions",
	Long: `TerraLambda is a command-line interface (CLI) that simplifies AWS Lambda deployments 
by integrating with Terraform, Step Functions, and monitoring tools.`,
}

// Execute is the entry point for the CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize all subcommands
func init() {
	RootCmd.AddCommand(GetDeployCmd()) // Deploy command
	RootCmd.AddCommand(GetInvokeCmd()) // Invoke command
	RootCmd.AddCommand(GetLogsCmd())   // CloudWatch Logging
}
