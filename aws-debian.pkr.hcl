packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

variable "ssh_username" {
  type    = string
  default = "admin"
}

variable "AMI_USERS" {
  type    = list(string)
  default = ["924297369992"]
}

variable "aws_region" {
  type    = string
  default = "us-west-1"
}

source "amazon-ebs" "debian" {
  ami_name      = "webapp-api-debian-aws_${formatdate("YYYY_MM_DD_hh_mm_ss", timestamp())}"
  instance_type = "t2.micro"

  region = "${var.aws_region}"

  ami_users = "${var.AMI_USERS}"

  source_ami_filter {
    filters = {
      name                = "debian-12-amd64-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["amazon"]
  }
  ssh_username = "${var.ssh_username}"
}


build {
  name    = "ami builder"
  sources = ["source.amazon-ebs.debian"]

  provisioner "file" {
    source      = "./webapplication_go.zip"
    destination = "~/webapplication_go.zip"
  }

  provisioner "shell" {
    inline = [
      "sudo apt-get update -y",
      "sudo apt-get upgrade -y",
      "sudo apt-get install -y unzip",
      "cd",
      "unzip webapplication_go.zip -d /home/admin/",
    ]
  }

  // provisioner "shell" {
  //   inline = [

  //     "cp /tmp/webapp.zip ~/",
  //     "unzip ~/webapp.zip -d ~/"

  //   ]
  // }

  provisioner "shell" {
    inline = [
      "ls -ltr",
      "cd /home/admin",
      "chmod +x /home/admin/setup.sh",
      "/home/admin/setup.sh"

    ]
  }

}