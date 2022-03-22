// variable "ami" {
//   description = "ami"
//   type        = string
//   default     = "ami-0726c1c7e01759156"
// }

data "aws_ami" "AmazonLinux2" {
  most_recent = true

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-2.0.20220218.3-x86_64-gp2"]
    #values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["137112412989"] # Amazon
}