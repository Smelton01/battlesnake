variable "project_name" {
  type = string
}

variable "env" {
  type = string
}

variable "aws_tenant_id" {
  type = string
}

variable "user_assume_roles" {
  type = list(string)
}