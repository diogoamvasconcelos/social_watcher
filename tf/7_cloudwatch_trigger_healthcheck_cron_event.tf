resource "aws_cloudwatch_event_rule" "trigger_healthcheck" {
  name                = "trigger_healthcheck"
  schedule_expression = "rate(5 minutes)"
}

resource "aws_cloudwatch_event_target" "trigger_pureref_healthcheck" {
  rule  = aws_cloudwatch_event_rule.trigger_healthcheck.name
  arn   = aws_lambda_function.healthcheck_website.arn
  input = "{\"Website\": \"http://pureref.com\"}"
}

resource "aws_lambda_permission" "allow_healthcheck_cloudwatch_event_lambda_invokation" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.healthcheck_website.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.trigger_healthcheck.arn
}
