variable "architectures" {
  type = list(string)
}


variable "project_name" {
  type = string
}

variable "env" {
  type = string
}

variable "execution_role" {
  type = string
}

variable "image_repo" {
  type = string
}

variable "image_tag" {
  type    = string
  default = "latest"
}

variable "memory_size" {
  default = 256
}

variable "timeout" {
  default = 15
}

variable "environment_variables" {
  type    = map(string)
  default = {}
}

variable "vpc_subnet_ids" {
  type    = list(string)
  default = []
}

variable "vpc_security_group_ids" {
  type    = list(string)
  default = []
}