### Common ### 


variable "env" {
  description = "Define Environment"
  type        = string
  default     = "dev"
}

variable "instance-type" {
  description = "Instance Type for Server"
  type        = string
  default     = "t3a.nano"
}
