resource "aws_cloudwatch_event_rule" "trigger_watchers" {
  name                = "trigger_watchers"
  schedule_expression = "rate(1 hour)"
}

resource "aws_cloudwatch_event_target" "trigger_watchers" {
  rule  = aws_cloudwatch_event_rule.trigger_watchers.name
  arn   = aws_lambda_function.watch_twitter_lambda.arn
  input = "{\"Keyword\": \"pureref\"}"
}

resource "aws_lambda_permission" "allow_cloudwatch_event_lambda_invokation" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.watch_twitter_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.trigger_watchers.arn
}