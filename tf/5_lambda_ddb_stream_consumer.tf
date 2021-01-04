
locals {
  ddb_stream_consumer_lambda_name = "ddb_stream_consumer"
  ddb_stream_consumer_lambda_file = "${var.out_dir}/${local.ddb_stream_consumer_lambda_name}.zip"
}

resource "aws_lambda_function" "ddb_stream_consumer" {
  filename         = "${local.ddb_stream_consumer_lambda_file}"
  function_name    = "${local.ddb_stream_consumer_lambda_name}"
  handler          = "${local.ddb_stream_consumer_lambda_name}"
  role             = "${aws_iam_role.lambda_default.arn}"
  runtime          = "go1.x"
  memory_size      = "128"
  timeout          = "30"
  source_code_hash = "${filebase64sha256("${local.ddb_stream_consumer_lambda_file}")}"
  description      = "  DynamoDB Stream Consumer"
  depends_on       = ["aws_cloudwatch_log_group.ddb_stream_consumer"]
}

resource "aws_cloudwatch_log_group" "ddb_stream_consumer" {
  name              = "/aws/lambda/${local.ddb_stream_consumer_lambda_name}"
  retention_in_days = 30
}

resource "aws_lambda_event_source_mapping" "ddb_stream_consumer_lambda_mapping" {
  event_source_arn                   = aws_dynamodb_table.stored_items.stream_arn
  function_name                      = aws_lambda_function.ddb_stream_consumer.arn
  starting_position                  = "LATEST"
  batch_size                         = 5
  maximum_batching_window_in_seconds = 30
  maximum_retry_attempts             = 5
  maximum_record_age_in_seconds      = 60
}
