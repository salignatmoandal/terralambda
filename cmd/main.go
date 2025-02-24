package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "terralambda",
		Short: "TerraLambda is a CLI tool to manage AWS Lambda functions",
	}

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(invokeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
