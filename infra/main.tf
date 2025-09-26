provider "aws" {
    region = "us-east-1"
}

# s3 bucket resource
resource "aws_s3_bucket" "notes" {
    bucket = "notes-bucket"
}

