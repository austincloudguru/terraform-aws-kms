variable "region" {
  description = "AWS Region"
  type        = string
  default     = "us-west-2"
}

variable "key_suffix" {
  description = "Suffix for the key to allow for multiple runs"
  type        = string
  default     = "123456"
}

variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default = {
    Environment = "terratest"
    Terraform   = "true"
  }
}
