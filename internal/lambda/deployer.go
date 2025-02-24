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
	// Compile dans le répertoire lambda existant
	if err := d.compileLambda(); err != nil {
		return fmt.Errorf("erreur de compilation: %w", err)
	}

	// Créer le ZIP dans le répertoire de travail
	if err := d.createZip(); err != nil {
		return fmt.Errorf("erreur de création du ZIP: %w", err)
	}

	// Déployer avec Terraform
	if err := d.applyTerraform(); err != nil {
		return fmt.Errorf("erreur de déploiement Terraform: %w", err)
	}

	return nil
}

func (d *LambdaDeployer) Cleanup() error {
	// Clean up temporary files
	return nil
}

func (d *LambdaDeployer) compileLambda() error {
	cmd := exec.CommandContext(d.ctx, "go", "build", "-o", "main")
	cmd.Dir = filepath.Join(d.workingDir, "lambda") // Utilise le répertoire lambda existant

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("compilation failed: %v, output: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) createZip() error {
	// Se déplacer dans le répertoire lambda
	cmd := exec.CommandContext(d.ctx, "zip", "-r", "../function.zip", "main")
	cmd.Dir = filepath.Join(d.workingDir, "lambda")

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("ZIP creation failed: %v, output: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) applyTerraform() error {
	terraformDir := filepath.Join(d.workingDir, "deployments", "terraform")

	// Initialize Terraform
	initCmd := exec.CommandContext(d.ctx, "terraform", "init")
	initCmd.Dir = terraformDir
	if output, err := initCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform init failed: %v, output: %s", err, output)
	}

	// Apply Terraform
	applyCmd := exec.CommandContext(d.ctx, "terraform", "apply", "-auto-approve")
	applyCmd.Dir = terraformDir
	if output, err := applyCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, output)
	}

	return nil
}
