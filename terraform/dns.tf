variable "ip" {}

data "digitalocean_domain" "domain" {
  name = var.domain
}

resource "digitalocean_record" "api" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = var.app
  value  = var.ip
  ttl    = 60
}
