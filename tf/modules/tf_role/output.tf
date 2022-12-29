output "terraform_role_arn" {
  value = aws_iam_role.terraform_role.arn
}

output "terraform_role_name" {
  value = aws_iam_role.terraform_role.name
}