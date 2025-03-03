package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

// Input Validation with regex
var validFunctionName = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
var validLogGroupPath = regexp.MustCompile(`^/aws/lambda/[a-zA-Z0-9_-]+$`)

// Constants for AWS command validation
const (
	maxFunctionNameLength = 64
	maxLogGroupLength     = 128
	awsLogsCommand        = "logs"
	awsTailCommand        = "tail"
	awsFollowFlag         = "--follow"
	awsLogGroupPrefix     = "/aws/lambda/"
)

// GetLogsCmd returns the logs command
func GetLogsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logs [function-name]",
		Short: "Fetch AWS CloudWatch logs for a Lambda function",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			functionName := args[0]

			if len(functionName) > maxFunctionNameLength {
				fmt.Printf("Function name is too long (maximum %d characters)\n", maxFunctionNameLength)
				return
			}

			if !validFunctionName.MatchString(functionName) {
				fmt.Println("Invalid function name. Allowed characters: letters, numbers, hyphens, and underscores.")
				return
			}

			logGroup := awsLogGroupPrefix + functionName
			if len(logGroup) > maxLogGroupLength {
				fmt.Printf("Log group path is too long (maximum %d characters)\n", maxLogGroupLength)
				return
			}

			if !validLogGroupPath.MatchString(logGroup) {
				fmt.Println("Invalid log group path format")
				return
			}

			fmt.Printf("Fetching logs for function: %s\n\n", functionName)

			awsPath, err := exec.LookPath("aws")
			if err != nil {
				fmt.Println("AWS CLI not found. Please install it and ensure it is in your PATH.")
				return
			}

			// Vérification supplémentaire : s'assurer que awsPath est bien un fichier exécutable
			awsPath, err = filepath.Abs(awsPath) // Obtient le chemin absolu
			if err != nil {
				fmt.Println("Failed to resolve AWS CLI path.")
				return
			}

			// Sécurisation supplémentaire pour éviter injection
			if _, err := os.Stat(awsPath); err != nil {
				fmt.Println("Invalid AWS CLI executable path.")
				return
			}

			// Création des arguments avec validation stricte
			cmdArgs := []string{
				awsLogsCommand,
				awsTailCommand,
				logGroup,
				awsFollowFlag,
			}

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			// Lancement sécurisé du processus
			cmdLog := exec.CommandContext(ctx, awsPath, cmdArgs...)
			cmdLog.Stdout = os.Stdout
			cmdLog.Stderr = os.Stderr

			if err := cmdLog.Run(); err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		},
	}
}
