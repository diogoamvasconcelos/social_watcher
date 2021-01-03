resource "aws_cloudwatch_event_rule" "trigger_watchers" {
  name                = "trigger_watchers"
  schedule_expression = "rate(1 hour)"
}

resource "aws_cloudwatch_event_target" "trigger_watchers_pureref" {
  rule  = aws_cloudwatch_event_rule.trigger_watchers.name
  arn   = aws_lambda_function.watch_twitter.arn
  input = "{\"Keyword\": \"pureref\"}"
}

resource "aws_lambda_permission" "allow_watcher_cloudwatch_event_lambda_invokation" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.watch_twitter.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.trigger_watchers.arn
}
