
# Allow fast rollback to previous version if a deployment fails


resource "aws_lambda_alias" "lambda_alias" {
  name             = "prod"
  function_name    = aws_lambda_function.terralambda_fn.function_name
  function_version = var.rollback_version
}