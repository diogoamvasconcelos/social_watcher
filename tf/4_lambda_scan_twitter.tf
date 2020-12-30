
locals {
  watch_twitter_lambda_name = "watch_twitter_lambda"
  watch_twitter_lambda_file = "${var.out_dir}/${local.watch_twitter_lambda_name}.zip"
}

resource "aws_lambda_function" "watch_twitter_lambda" {
  filename         = "${local.watch_twitter_lambda_file}"
  function_name    = "${local.watch_twitter_lambda_name}"
  handler          = "${local.watch_twitter_lambda_name}"
  role             = "${aws_iam_role.lambda_default.arn}"
  runtime          = "go1.x"
  memory_size      = "512"
  timeout          = "3"
  source_code_hash = "${filebase64sha256("${local.watch_twitter_lambda_file}")}"
  description      = "Watch Twitter Lambda"
  depends_on       = ["aws_cloudwatch_log_group.watch_twitter_lambda"]
}

resource "aws_cloudwatch_log_group" "watch_twitter_lambda" {
  name              = "/aws/lambda/${local.watch_twitter_lambda_name}"
  retention_in_days = 30
}
