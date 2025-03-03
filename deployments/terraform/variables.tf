# AWS Region
variable "aws_region" {
  description = "AWS region where Lambda is deployed"
  type        = string
  default     = "us-east-1"
}

# Lambda Function Name
variable "lambda_function_name" {
  description = "Name of the AWS Lambda function"
  type        = string
  default     = "MyLambdaFunction"
}

# Memory Size (MB)
variable "lambda_memory_size" {
  description = "Memory allocation for Lambda function"
  type        = number
  default     = 128
}

# Execution Timeout (seconds)
variable "lambda_timeout" {
  description = "Execution timeout for Lambda function"
  type        = number
  default     = 3
}

variable "rollback_version" {
  description = "The Lambda version to rollback to"
  type        = string
  default     = "$LATEST"
}