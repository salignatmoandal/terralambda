package lambda

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

type LambdaDeployer struct {
	ctx        context.Context
	workingDir string
}

func NewDeployer(ctx context.Context, workingDir string) *LambdaDeployer {
	return &LambdaDeployer{
		ctx:        ctx,
		workingDir: workingDir,
	}
}

func (d *LambdaDeployer) Deploy(functionName string, zipPath string) error {
	operations := []struct {
		name string
		fn   func() error
	}{
		{"compile", d.compileLambda},
		{"create ZIP", d.createZip},
		{"terraform deployment", d.applyTerraform},
	}

	for _, op := range operations {
		if err := op.fn(); err != nil {
			return fmt.Errorf("error during %s: %w", op.name, err)
		}
	}

	return nil
}

func (d *LambdaDeployer) Cleanup() error {
	// Clean up temporary files
	return nil
}

func (d *LambdaDeployer) compileLambda() error {
	cmd := exec.CommandContext(d.ctx, "go", "build", "-o", "lambda", "main.go")
	cmd.Dir = d.workingDir

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("compilation failed: %v, output: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) createZip() error {
	cmd := exec.CommandContext(d.ctx, "zip", "-r", "function.zip", "lambda")
	cmd.Dir = d.workingDir

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("ZIP creation failed: %v, output: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) applyTerraform() error {
	cmd := exec.CommandContext(d.ctx, "terraform", "init")
	cmd.Dir = filepath.Join(d.workingDir, "deployments", "terraform")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform init failed: %v, output: %s", err, output)
	}

	cmd = exec.CommandContext(d.ctx, "terraform", "apply", "-auto-approve")
	cmd.Dir = filepath.Join(d.workingDir, "deployments", "terraform")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, output)
	}

	return nil
}
