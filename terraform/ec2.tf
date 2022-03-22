

variable "availability_zones" {
  default = [
    "ap-northeast-1a",
    "ap-northeast-1d"
  ]
}
# Create a VPC
resource "aws_vpc" "elb-testing" {
  cidr_block = "10.0.0.0/16"

  tags = {
    project = var.project
  }
}

resource "aws_subnet" "elb-testing-subnets" {
  count             = length(var.availability_zones)
  vpc_id            = aws_vpc.elb-testing.id
  cidr_block        = cidrsubnet("10.0.0.0/16", 8, count.index)
  availability_zone = var.availability_zones[count.index]

  tags = {
    subnet  = "${var.project}-${count.index}"   # Anti Colour render"
    project = var.project
  }
}


resource "aws_security_group" "all-ec2" {
  name        = "all-ec2"
  description = "allow all outcomming traffic to bastion ssh"
  vpc_id      = aws_vpc.elb-testing.id
  
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8443
    to_port     = 8443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    env     = "dev"
    project = var.project
  }
}

// resource "aws_instance" "elb-testing" {
//   ami           = data.aws_ami.AmazonLinux2.id
//   instance_type = "t3a.nano"

//   count                       = 2
//   associate_public_ip_address = true

//   key_name = "dyingapple-dev"

//   vpc_security_group_ids = [
//     aws_security_group.all-ec2.id
//   ]

//   subnet_id = "subnet-1d5c1a54"

//   tags = {
//     env     = "dev"
//     project = "elb-testing"
//   }
// }


# Put an instance in each subnet
resource "aws_instance" "elb-testing" {
  count         = length(var.availability_zones)
  ami           = data.aws_ami.AmazonLinux2.id
  instance_type = var.instance-type
  subnet_id     = aws_subnet.elb-testing-subnets[count.index].id

  vpc_security_group_ids = [
    aws_security_group.all-ec2.id
  ]

  key_name = "dyingapple-dev"

  tags = {
    env      = "dev"
    instance = "${var.project}-${count.index}" # Anti Colour render"
    project  = var.project
  }
}