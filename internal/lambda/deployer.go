package lambda

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

// LambdaDeployer handles the deployment process
type LambdaDeployer struct {
	ctx        context.Context
	workingDir string
}

// NewDeployer creates a new LambdaDeployer instance
func NewDeployer(ctx context.Context, workingDir string) *LambdaDeployer {
	return &LambdaDeployer{
		ctx:        ctx,
		workingDir: workingDir,
	}
}

// Deploy compiles, packages, and deploys the Lambda function
func (d *LambdaDeployer) Deploy(functionName string, zipPath string) error {
	if err := d.compileLambda(); err != nil {
		return fmt.Errorf("error during compilation: %w", err)
	}

	if err := d.createZip(); err != nil {
		return fmt.Errorf("error creating ZIP: %w", err)
	}

	if err := d.applyTerraform(); err != nil {
		return fmt.Errorf("error applying Terraform: %w", err)
	}

	return nil
}

// Cleanup removes temporary files
func (d *LambdaDeployer) Cleanup() error {
	// Add cleanup logic if necessary
	return nil
}

// compileLambda compiles the Go Lambda function into an executable
func (d *LambdaDeployer) compileLambda() error {
	cmd := exec.CommandContext(d.ctx, "go", "build", "-o", "bootstrap", "main.go")
	cmd.Dir = filepath.Join(d.workingDir, "lambda")

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("compilation failed: %v, output: %s", err, output)
	}
	return nil
}

// createZip packages the Lambda function into a ZIP file
func (d *LambdaDeployer) createZip() error {
	cmd := exec.CommandContext(d.ctx, "zip", "-j", "function.zip", "bootstrap")
	cmd.Dir = filepath.Join(d.workingDir, "lambda")

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("ZIP creation failed: %v, output: %s", err, output)
	}
	return nil
}

// applyTerraform executes Terraform commands to deploy the Lambda function
func (d *LambdaDeployer) applyTerraform() error {
	terraformDir := filepath.Join(d.workingDir, "deployments", "terraform")

	cmd := exec.CommandContext(d.ctx, "terraform", "init")
	cmd.Dir = terraformDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform init failed: %v, output: %s", err, output)
	}

	cmd = exec.CommandContext(d.ctx, "terraform", "apply", "-auto-approve")
	cmd.Dir = terraformDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, output)
	}

	return nil
}
