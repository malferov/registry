variable "domain" {}
variable "app" {}
variable "email" {}

resource "tls_private_key" "private_key" {
  algorithm   = "RSA"
  ecdsa_curve = "P256"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.private_key.private_key_pem
  email_address   = var.email
}

resource "acme_certificate" "cert" {
  account_key_pem           = acme_registration.reg.account_key_pem
  common_name               = "${var.app}.${var.domain}"
  dns_challenge {
    provider = "digitalocean"

    config = {
      DO_AUTH_TOKEN = var.do_token
    }
  }
}

output "crt" {
  sensitive = true
  value     = base64encode("${acme_certificate.cert.certificate_pem}${acme_certificate.cert.issuer_pem}")
}

output "key" {
  sensitive = true
  value     = base64encode(acme_certificate.cert.private_key_pem)
}
