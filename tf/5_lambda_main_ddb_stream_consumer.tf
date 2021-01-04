
locals {
  main_ddb_stream_consumer_lambda_name = "main_ddb_stream_consumer"
  main_ddb_stream_consumer_lambda_file = "${var.out_dir}/${local.main_ddb_stream_consumer_lambda_name}.zip"
}

resource "aws_lambda_function" "main_ddb_stream_consumer" {
  filename         = "${local.main_ddb_stream_consumer_lambda_file}"
  function_name    = "${local.main_ddb_stream_consumer_lambda_name}"
  handler          = "${local.main_ddb_stream_consumer_lambda_name}"
  role             = "${aws_iam_role.lambda_default.arn}"
  runtime          = "go1.x"
  memory_size      = "128"
  timeout          = "30"
  source_code_hash = "${filebase64sha256("${local.main_ddb_stream_consumer_lambda_file}")}"
  description      = "DynamoDB Stream Consumer for the main table"
  depends_on       = ["aws_cloudwatch_log_group.main_ddb_stream_consumer"]
}

resource "aws_cloudwatch_log_group" "main_ddb_stream_consumer" {
  name              = "/aws/lambda/${local.main_ddb_stream_consumer_lambda_name}"
  retention_in_days = 30
}

resource "aws_lambda_event_source_mapping" "main_ddb_stream_consumer_lambda_mapping" {
  event_source_arn                   = aws_dynamodb_table.main_table.stream_arn
  function_name                      = aws_lambda_function.main_ddb_stream_consumer.arn
  starting_position                  = "LATEST"
  batch_size                         = 1
  maximum_batching_window_in_seconds = 30
  maximum_retry_attempts             = 5
  maximum_record_age_in_seconds      = 60
}
