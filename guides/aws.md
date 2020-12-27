# Regions
- calulate using: https://calculator.aws/
- eu-north-1 (stockholm): cheapest in EU

# Terraform

## Guides:
* https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
* https://learn.hashicorp.com/terraform/getting-started/build
* https://ordina-jworks.github.io/cloud/2019/01/14/Infrastructure-as-code-with-terraform-and-aws-serverless.html#api-gateway  (api-gateway-openapi/swagger)

## AWS Credentials
* Create user for pragamatic access (user admin) using the AWS Console > IAM

* Put its credentials in ~/.aws/credentials (by ruuning the awscli cmd: `aws configure`)

## Terraform
All `.tf` configuration files are in ./tf and the load order is alphabetical
```
cd into `./tf`
run: `terraform init` to generate the terraform state file
run: `terraform plan / apply` to deploy!
```

### Notes
* apigateway cloudwatch group must be named like this "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.main_api.id}/${aws_api_gateway_deployment.main_api.stage_name"