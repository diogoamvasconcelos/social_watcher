locals {
    hello_world_lambda_file = "${var.out_dir}/lambda_artifact.zip"
    hello_world_lambda_name = "hello_world_lambda"
    hello_world_lambda_handler = "lambda_artifact"
    hello_world_lambda_role = "hello_world_lambda"
}

resource "aws_lambda_function" "hello_world_lambda" {
  filename         = "${local.hello_world_lambda_file}"
  function_name    = "${local.hello_world_lambda_name}"
  handler          = "${local.hello_world_lambda_handler}"
  role             = "${aws_iam_role.hello_world_lambda.arn}"
  runtime          = "go1.x"
  memory_size      = "512" 
  timeout          = "3"
  source_code_hash = "${filebase64sha256("${local.hello_world_lambda_file}")}"
  description      = "Hello World Lambda"
  depends_on       = ["aws_cloudwatch_log_group.hello_world_lambda", 
                      "aws_iam_role.hello_world_lambda",
                      "aws_iam_role_policy_attachment.hello_world_lambda"]
}

resource "aws_cloudwatch_log_group" "hello_world_lambda" {
  name              = "/aws/lambda/${local.hello_world_lambda_name}"
  retention_in_days = 30
}

resource "aws_iam_role" "hello_world_lambda" {
  name = "${local.hello_world_lambda_role}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "hello_world_lambda" {
  name = "${local.hello_world_lambda_name}"
  path = "/"
  description = "IAM policy for Hello World lambda - Logging"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "hello_world_lambda" {
  role = "${aws_iam_role.hello_world_lambda.name}"
  policy_arn = "${aws_iam_policy.hello_world_lambda.arn}"
}