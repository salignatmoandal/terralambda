package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terralambda",
	Short: "TerraLambda is a CLI tool to manage AWS Lambda functions",
}

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(invokeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
