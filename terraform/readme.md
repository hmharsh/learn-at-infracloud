# Terraform 
## Hashicorp configuration language

provider "aws" {
access_key = ""
secret_key = ""
region = ""
}

resource "aws_instance" "example" {
ami = ""
inatance_type = ""
}

Better example with terraform and variable block at: https://www.terraform.io/docs/configuration/index.html
