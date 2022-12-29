locals {
  project_name        = "snek"
  env                 = "stg"
  snek_app_image_repo = "998684527758.dkr.ecr.ap-northeast-1.amazonaws.com/snek-stg-app"
}

provider "aws" {
  region              = var.aws_region
  allowed_account_ids = [var.aws_account_id]

  default_tags {
    tags = {
      Project = local.project_name
    }
  }
}

module "terraform_role" {
  source = "../../modules/tf_role"

  aws_tenant_id = var.aws_account_id
  project_name  = local.project_name
  env           = local.env
  user_assume_roles = [
    "arn:aws:sts::998684527758:role/root"
  ]
}

module "lambda_role" {
  source = "../../modules/role"

  project_name = local.project_name
  env          = local.env
  aws_region   = var.aws_region
}

module "snek_ecr_repo" {
  source = "../../modules/ecr_repo"

  aws_region   = var.aws_region
  project_name = local.project_name
  env          = local.env
}

module "snek_app" {
  source = "../../modules/app"

  architectures  = ["arm64"]
  project_name   = local.project_name
  env            = local.env
  execution_role = module.lambda_role.role_arn
  image_repo     = local.snek_app_image_repo
  image_tag      = var.snek_image_tag
  environment_variables = {
    "APP_ENV" = local.env
  }
}

resource "aws_apigatewayv2_api" "api" {
  name          = "${local.project_name}-${local.env}-api-gw"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id      = aws_apigatewayv2_api.api.id
  name        = local.env
  auto_deploy = true
}

resource "aws_apigatewayv2_integration" "snek" {
  api_id           = aws_apigatewayv2_api.api.id
  integration_type = "AWS_PROXY"
  connection_type  = "INTERNET"

  integration_method = "POST"
  integration_uri    = module.snek_app.invoke_arn
  request_parameters = {
    "overwrite:path" = "$request.path"
  }
}



resource "aws_apigatewayv2_route" "snek" {
  api_id = aws_apigatewayv2_api.api.id

  target    = "integrations/${aws_apigatewayv2_integration.snek.id}"
  route_key = "$default"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = module.snek_app.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.api.execution_arn}/*/*"
}

resource "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"
  
  client_id_list = [
    "sts.amazonaws.com",
  ]
  thumbprint_list = [
    "6938fd4d98bab03faadb97b34396831e3780aea1",
  ]
}