provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "terralambda_fn" {
  function_name = "TerraLambdaExample"
  role          = aws_iam_role.lambda_exec.arn
  handler       = "main"
  runtime       = "go1.x"
  filename      = "${path.module}/../../function.zip"
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
      "Effect": "Allow",
      "Principal": {"Service": "lambda.amazonaws.com"},
      "Action": "sts:AssumeRole"
  }]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}