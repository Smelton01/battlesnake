resource "aws_apigatewayv2_api" "lambda" {
  name          = "${var.project_name}-${var.env}-api-gw"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.lambda.id

  name        = "${var.project_name}-${var.env}-lambda-stage"
  auto_deploy = true
}

resource "aws_apigatewayv2_integration" "snek" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = var.lambda_function
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "hello_world" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "GET /status"
  target    = "integrations/${aws_apigatewayv2_integration.hello_world.id}"
}

# resource "aws_cloudwatch_log_group" "api_gw" {
#   name = "/aws/api_gw/${aws_apigatewayv2_api.lambda.name}"

#   retention_in_days = 30
# }

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.hello_world.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}
