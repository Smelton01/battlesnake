resource "aws_iam_role" "terraform_role" {
  name        = "${var.project_name}-${var.env}-terraform-role"
  description = "Automate deployment of battlesnake Snek to AWS"

  assume_role_policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": ["${join("\", \"", var.user_assume_roles)}"]
            },
            "Action": ["sts:AssumeRole", "sts:TagSession"]
        },
        {
            "Effect": "Allow",
            "Principal": {
                "Federated" : ["arn:aws:iam::${var.aws_tenant_id}:oidc-provider/token.actions.githubusercontent.com"]
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringLike": {
                    "token.actions.githubusercontent.com:sub": "repo:Smelton01/battlesnake:*"
                }
            }
        }
    ]
}
POLICY
}

resource "aws_iam_role_policy" "terraform_role_policy" {
  name = "${var.project_name}-${var.env}-terraform-role-iam-management-policy"
  role = aws_iam_role.terraform_role.id

  policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ReadOnlyPermissions",
            "Effect": "Allow",
            "Action": [
                "iam:Get*"
            ],
            "Resource": [
                "arn:aws:iam::${var.aws_tenant_id}:role/${var.project_name}*",
                "arn:aws:iam::${var.aws_tenant_id}:instance-profile/${var.project_name}*",
                "arn:aws:iam::${var.aws_tenant_id}:policy/${var.project_name}*"
            ]
        },
        {
            "Sid": "S3",
            "Effect": "Allow",
            "Action": [
                "s3:*"
            ],
            "Resource": [
                "arn:aws:s3:::${var.project_name}*"
            ]
        },
        {
            "Sid": "Lambda",
            "Effect": "Allow",
            "Action": [
                "lambda:*"
            ],
            "Resource": [
                "arn:aws:lambda:*:${var.aws_tenant_id}:function:${var.project_name}*"
            ]
        },
        {
            "Sid": "ECR",
            "Effect": "Allow",
            "Action": [
                "ecr:*"
            ],
            "Resource": [
                "arn:aws:ecr:*:${var.aws_tenant_id}:repository/${var.project_name}*"
            ]
        },
        {
            "Sid": "ECRAuthorizationToken",
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken"
            ],
            "Resource": "*"
        },
        {
            "Sid": "SNS",
            "Effect": "Allow",
            "Action": [
                "sns:*"
            ],
            "Resource": [
                "arn:aws:sns:*:${var.aws_tenant_id}:${var.project_name}*"
            ]
        },
        {
            "Sid": "ListAllPermissions",
            "Effect": "Allow",
            "Action": [
                "iam:List*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "CreateModifyPermissions",
            "Effect": "Allow",
            "Action": [
                "iam:Add*",
                "iam:Create*",
                "iam:Update*",
                "iam:Put*",
                "iam:Attach*",
                "iam:Detach*",
                "iam:Delete*",
                "iam:PassRole",
                "iam:GenerateServiceLastAccessedDetails",
                "iam:GenerateCredentialReport",
                "iam:Tag*"
            ],
            "Resource": [
                "arn:aws:iam::${var.aws_tenant_id}:role/${var.project_name}*",
                "arn:aws:iam::${var.aws_tenant_id}:instance-profile/${var.project_name}*",
                "arn:aws:iam::${var.aws_tenant_id}:policy/${var.project_name}*"
            ]
        },
        {
            "Sid": "DenyPermissions",
            "Effect": "Deny",
            "Action": [
                "iam:CreateUser",
                "iam:DeleteUser"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
POLICY
}