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
	// Compile the lambda
	if err := d.compileLambda(); err != nil {
		return fmt.Errorf("erreur de compilation: %w", err)
	}

	// Création du ZIP
	if err := d.createZip(); err != nil {
		return fmt.Errorf("erreur de création du ZIP: %w", err)
	}

	// Déploiement Terraform
	if err := d.applyTerraform(); err != nil {
		return fmt.Errorf("erreur de déploiement Terraform: %w", err)
	}

	return nil
}

func (d *LambdaDeployer) Cleanup() error {
	// Nettoyage des fichiers temporaires
	return nil
}

func (d *LambdaDeployer) compileLambda() error {
	cmd := exec.CommandContext(d.ctx, "go", "build", "-o", "lambda", "main.go")
	cmd.Dir = d.workingDir

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("échec de la compilation: %v, sortie: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) createZip() error {
	cmd := exec.CommandContext(d.ctx, "zip", "-r", "function.zip", "lambda")
	cmd.Dir = d.workingDir

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("échec de la création du ZIP: %v, sortie: %s", err, output)
	}
	return nil
}

func (d *LambdaDeployer) applyTerraform() error {
	cmd := exec.CommandContext(d.ctx, "terraform", "init")
	cmd.Dir = filepath.Join(d.workingDir, "deployments", "terraform")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("échec de terraform init: %v, sortie: %s", err, output)
	}

	cmd = exec.CommandContext(d.ctx, "terraform", "apply", "-auto-approve")
	cmd.Dir = filepath.Join(d.workingDir, "deployments", "terraform")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("échec de terraform apply: %v, sortie: %s", err, output)
	}

	return nil
}
