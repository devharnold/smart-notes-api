provider "aws" {
    region = var.region
}

# s3 bucket resource
resource "aws_s3_bucket" "notes" {
    bucket = var.bucket_name
}