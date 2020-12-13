terraform {
  required_version = ">= 0.12"
}

provider "google" {
  project     = "astute-charter-295013"
  region      = "europe-west3"
  zone    = "europe-west3-a"
  credentials = file(var.gcp_auth_file)
}
