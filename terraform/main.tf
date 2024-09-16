terraform {
  required_providers {
    ssh = {
      source  = "loafee/ssh"
      version = "~> 2.7.0"
    }
  }
}

resource "ssh_resource" "server_1" {
  connection {
    host        = var.public_ip
    user        = var.ssh_username
    private_key = var.ssh_private_key
    agent       = false
  }

  provisioner "remote-exec" {
    inline = ["${var.instantiate_services_command}"]
  }

  destroy {
    pre_destroy = ["${var.destroy_services_command}"]
  }
}
