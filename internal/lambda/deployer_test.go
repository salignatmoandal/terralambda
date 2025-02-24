package lambda

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestDeployer_Deploy(t *testing.T) {
	// Créer une structure de test temporaire
	testDir := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	ctx := context.Background()
	deployer := NewDeployer(ctx, testDir)

	err := deployer.Deploy("test-function", "test.zip")
	if err != nil {
		t.Fatalf("Deploy failed: %v", err)
	}
}

func setupTestEnvironment(t *testing.T) string {
	testDir, err := os.MkdirTemp("", "terralambda-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Structure des répertoires
	createDirectories(t, testDir)

	// Fichiers Lambda
	createLambdaFiles(t, testDir)

	// Fichiers Terraform
	createTerraformFiles(t, testDir)

	return testDir
}

func createDirectories(t *testing.T, testDir string) {
	dirs := []string{
		filepath.Join(testDir, "lambda"),
		filepath.Join(testDir, "deployments", "terraform"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}
}

func createLambdaFiles(t *testing.T, testDir string) {
	// main.go pour Lambda avec aws-lambda-go
	lambdaContent := []byte(`package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string ` + "`json:\"message\"`" + `
}

func HandleRequest(ctx context.Context) (Response, error) {
	return Response{
		Message: "Hello from Lambda!",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}`)

	mainPath := filepath.Join(testDir, "lambda", "main.go")
	if err := os.WriteFile(mainPath, lambdaContent, 0644); err != nil {
		t.Fatalf("Failed to create main.go: %v", err)
	}

	// go.mod avec les dépendances requises
	goModContent := []byte(`module testlambda

go 1.21

require github.com/aws/aws-lambda-go v1.41.0

require (
	github.com/aws/aws-sdk-go-v2 v1.36.2
	github.com/aws/aws-sdk-go-v2/config v1.29.7
)
`)

	goModPath := filepath.Join(testDir, "lambda", "go.mod")
	if err := os.WriteFile(goModPath, goModContent, 0644); err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Initialiser le module Go
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = filepath.Join(testDir, "lambda")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to run go mod tidy: %v, output: %s", err, output)
	}
}

func createTerraformFiles(t *testing.T, testDir string) {
	// Copie de la configuration Terraform de base
	terraformContent := []byte(`
provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "test" {
  filename         = "../../function.zip"
  function_name    = "test-function"
  role            = aws_iam_role.lambda_exec.arn
  handler         = "main"
  runtime         = "provided.al2"
}

resource "aws_iam_role" "lambda_exec" {
  name = "test_lambda_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}
`)

	tfPath := filepath.Join(testDir, "deployments", "terraform", "main.tf")
	if err := os.WriteFile(tfPath, terraformContent, 0644); err != nil {
		t.Fatalf("Failed to create main.tf: %v", err)
	}
}
