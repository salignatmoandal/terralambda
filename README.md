# TerraLambda CLI Documentation

## Overview 
TerraLambda is a CLI tool designed to simplify the deployment and invocation of AWS Lambda functions using Terraform. It streamlines the function lifecycle by automating packaging, deployment, and execution processes, providing a seamless experience for developers working with AWS Lambda.
## Core concepts 
1. ### Infrastructure as Code (IaC) with Terraform
  TerraLambda leverages Terraform to define and manage AWS Lambda functions, ensuring a declarative and reproducible deployment process.

2. ### AWS Lambda Invocation 
- The tool compiles the source code, packages it into a ZIP file, and deploys it using Terraform.
- Manages different versions of Lambda functions to support rollback mechanisms.

3. ## AWS Lambda Invocation
- Allows invoking deployed Lambda functions via the AWS SDK.
- Supports sending payloads to test functions interactively.

4. ## Rollback Mechanisms (in progress) 
Uses AWS Lambda versioning and aliases to revert to a previous stable version.
Integrates Terraform state management to facilitate rollback operations.

# Installation 
## Prerequesistes 
- Go (latest stable version)
- Terraform (installed and configured for AWS)
- AWS CLI (configured with proper credentials)

