
# Allow fast rollback to previous version if a deployment fails

variable "rollback_version" {
  description = "The Lambda version to rollback to"
  type        = string
}

resource "aws_lambda_alias" "lambda_alias" {
  name             = "prod"
  function_name    = aws_lambda_function.terralambda_fn.function_name
  function_version = var.rollback_version
}