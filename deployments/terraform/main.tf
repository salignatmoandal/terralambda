provider "aws" {
  region = "us-east-1"
}

# Ajout d'une ressource null_resource pour créer le ZIP avant le déploiement
resource "null_resource" "lambda_zip" {
  provisioner "local-exec" {
    command = "cd ../../lambda && GOOS=linux GOARCH=amd64 go build -o bootstrap main.go && zip function.zip bootstrap && mv function.zip ../deployments/terraform/"
  }

  triggers = {
    always_run = "${timestamp()}"
  }
}

resource "aws_lambda_function" "terralambda_fn" {
  function_name = "TerraLambdaExample"
  role          = aws_iam_role.lambda_exec.arn
  handler       = "bootstrap"
  runtime       = "provided.al2"
  filename      = "${path.module}/function.zip"
  
  depends_on = [null_resource.lambda_zip]
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