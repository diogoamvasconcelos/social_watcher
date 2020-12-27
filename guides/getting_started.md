# Credentials

## AWS account
* email:diogo.am.vasconcelos+sw@gmail.com
* pass: 3.aws
* username: diogo.vasconcelos.sw

# Configuration (installation - Ubuntu)

## Terraform
Extract the .zip from the latest terraform build and move it to `/usr/bin`. Try running `terraform` in the terminal. If you can't, make sure the `/usr/bin` was added to the #PATH.

## AWS CLI
```
apt install awscli
```

## GO
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
```
go get github.com/aws/aws-lambda-go/lambda
```