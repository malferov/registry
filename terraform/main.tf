terraform {
  backend "remote" {}
}

variable "letsencrypt_url" {}
variable "do_token" {}

provider "acme" {
  server_url = var.letsencrypt_url
}

provider "digitalocean" {
  token = var.do_token
}
