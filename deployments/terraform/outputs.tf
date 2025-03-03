# Output the Lambda function name
output "lambda_function_name" {
  description = "The name of the deployed AWS Lambda function"
  value       = aws_lambda_function.terralambda_fn.function_name
}

# Output the Lambda function ARN
output "lambda_function_arn" {
  description = "The ARN of the deployed AWS Lambda function"
  value       = aws_lambda_function.terralambda_fn.arn
}

# Output the latest deployed Lambda version
output "lambda_function_version" {
  description = "The latest version of the deployed Lambda function"
  value       = aws_lambda_function.terralambda_fn.version
}

# Output the Lambda function invoke ARN
output "lambda_invoke_arn" {
  description = "The ARN used to invoke the Lambda function"
  value       = aws_lambda_function.terralambda_fn.invoke_arn
}