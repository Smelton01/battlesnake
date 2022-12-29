resource "aws_iam_role" "lambda_role" {
  name = "${var.project_name}-${var.env}-lambda-role"

  assume_role_policy = jsonencode(
    {
      Version : "2012-10-17",
      Statement : [
        {
          Sid : "",
          Effect : "Allow",
          Principal : {
            Service : "lambda.amazonaws.com"
          },
          Action : "sts:AssumeRole"
        }
      ]
    }
  )
}

resource "aws_iam_role_policy" "lambda_role_policy" {
  name = "${var.project_name}-${var.env}-lambda-role-policy"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode(
    {
      Version : "2012-10-17",
      Statement : [
        {
          Sid : "CloudWatchLogs",
          Effect : "Allow",
          Action : [
            "logs:*"
          ],
          Resource : "*"
        },
        {
            Sid: "AllowExecutionFromAPIGateway",
            Effect: "Allow",
            Action: [
                "lambda:InvokeFunction"
            ],
            Resource: "*"
        }
      ]
    }
  )
}