variable "gce_ssh_user" {
  type        = string
  description = "GCE user"
}

variable "gce_ssh_pub_key_file1" {
  type        = string
  description = "GCE ssh public key"
  default = "ssh-key_alibabaih"
}

variable "gce_ssh_pub_key_file2" {
  type        = string
  description = "GCE ssh public key"
  default = "ssh-key_west"
}

variable "gce_ssh_pub_key_file3" {
  type        = string
  description = "GCE ssh public key"
  default = "ssh-key_west_linux"
}
