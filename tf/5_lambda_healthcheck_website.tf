
locals {
  healthcheck_website_lambda_name = "healthcheck_website_lambda"
  healthcheck_website_lambda_file = "${var.out_dir}/${local.healthcheck_website_lambda_name}.zip"
}

resource "aws_lambda_function" "healthcheck_website_lambda" {
  filename         = "${local.healthcheck_website_lambda_file}"
  function_name    = "${local.healthcheck_website_lambda_name}"
  handler          = "${local.healthcheck_website_lambda_name}"
  role             = "${aws_iam_role.lambda_default.arn}"
  runtime          = "go1.x"
  memory_size      = "512"
  timeout          = "3"
  source_code_hash = "${filebase64sha256("${local.healthcheck_website_lambda_file}")}"
  description      = "Healthcheck Website Lambda"
  depends_on       = ["aws_cloudwatch_log_group.healthcheck_website_lambda"]

  environment {
    variables = {
      "MAIN_TABLE_NAME" = "${aws_dynamodb_table.main_table.name}"
    }
  }
}

resource "aws_cloudwatch_log_group" "healthcheck_website_lambda" {
  name              = "/aws/lambda/${local.healthcheck_website_lambda_name}"
  retention_in_days = 30
}
