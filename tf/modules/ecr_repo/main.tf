resource "aws_ecr_repository" "snek-ecr-repo" {
  name = "${var.project_name}-${var.env}-app"

  image_scanning_configuration {
    scan_on_push = true
  }
}