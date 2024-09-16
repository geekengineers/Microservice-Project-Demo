variable "public_ip" {
  description = "Public ip address of the server"
}

variable "ssh_username" {
  description = "Username of the ssh server"
  default     = "root"
}

variable "ssh_private_key" {
  description = "Private key of the ssh server"
  type        = string
}

variable "instantiate_services_command" {
  default = "docker-compose up -d"
}
