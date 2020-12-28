terraform {
  required_providers {
    acme = {
      source = "vancluever/acme"
    }
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
    tls = {
      source = "hashicorp/tls"
    }
  }
  required_version = ">= 0.13"
}
