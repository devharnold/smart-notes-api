variable "region" {
  description = "The region where the S3 resource will be deployed"
  type        = string
  default     = "eu-west-2"
}

variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
  default     = "notes-bucket-arnold"
}