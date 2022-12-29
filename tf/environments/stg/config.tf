terraform {
  required_version = "1.3.4"

  backend "s3" {
    bucket  = "snek-stg-infra"
    key     = "tfstate/stag/terraform.tfstate"
    region  = "ap-northeast-1"
    encrypt = "true"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=4.34.0"
    }
  }
}