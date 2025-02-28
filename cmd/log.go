package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// GetLogsCmd returns the logs command
func GetLogsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logs [function-name]",
		Short: "Fetch AWS CloudWatch logs for a Lambda function",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			functionName := args[0]
			logGroup := fmt.Sprintf("/aws/lambda/%s", functionName)

			fmt.Printf("Fetching logs for function: %s\n\n", functionName)

			// Exécute la commande AWS CLI pour récupérer les logs
			cmdLog := exec.Command("aws", "logs", "tail", logGroup, "--follow")
			cmdLog.Stdout = cmd.OutOrStdout()
			cmdLog.Stderr = cmd.OutOrStderr()

			if err := cmdLog.Run(); err != nil {
				fmt.Printf("Error fetching logs: %v\n", err)
			}
		},
	}
}
