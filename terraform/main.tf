terraform {
  required_providers {
    ssh = {
      source  = "loafoe/ssh"
      version = "2.7.0"
    }
  }
}

provider "ssh" {}

resource "ssh_resource" "default_server" {
  host        = var.public_ip
  user        = var.ssh_username
  private_key = file(var.ssh_private_key)
  agent       = false

  when = "create"

  commands = [
    "${var.instantiate_services_command}"
  ]
}

output "result" {
  value = ssh_resource.default_server.result
}
