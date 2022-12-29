resource "aws_lambda_function" "app" {
  architectures = var.architectures
  function_name = "${var.project_name}-${var.env}-app"
  description   = "Snek lambda app"
  role          = var.execution_role
  package_type  = "Image"
  image_uri     = "${var.image_repo}:${var.image_tag}"

  memory_size = var.memory_size
  timeout     = var.timeout

  publish = false

  dynamic "environment" {
    for_each = length(var.environment_variables) > 0 ? [true] : []
    content {
      variables = var.environment_variables
    }
  }

  dynamic "vpc_config" {
    for_each = length(var.vpc_subnet_ids) > 0 ? [true] : []
    content {
      subnet_ids         = var.vpc_subnet_ids
      security_group_ids = var.vpc_security_group_ids
    }
  }
  
}