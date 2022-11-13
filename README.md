# Terraform Mail-in-a-Box Provider

![GitHub release (latest by date)](https://img.shields.io/github/v/release/markcaudill/terraform-provider-mailinabox)
![GitHub](https://img.shields.io/github/license/markcaudill/terraform-provider-mailinabox)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/markcaudill/terraform-provider-mailinabox)
![Codecov](https://img.shields.io/codecov/c/github/markcaudill/terraform-provider-mailinabox)

This is an unofficial Mail-in-a-Box Terraform provider. Currently it only
supports DNS record resources.

## Using the provider

```terraform
terraform {
  required_providers {
    mailinabox = {
      source = "markcaudill/mailinabox"
      version >= "2.0"
    }
  }
}

provider "mailinabox" {
  url      = "https://mail.example.com"
  username = "admin@example.com"
  password = "abc123fireme"
}

data "mailinabox_dnsrecord" "bare_a_record" {
  domain = "example.com"
  type = "A"
}

resource "mailinabox_dnsrecord" "testdomain_a_dnsrecord" {
  domain = "testdomain.example.com"
  type   = data.mailinabox_dnsrecord.bare_a_record.type
  value  = data.mailinabox_dnsrecord.bare_a_record.value
}

resource "mailinabox_dnsrecord" "testdomain_aaaa_dnsrecord" {
  domain = "testdomain.example.com"
  type = "AAAA"
  value = "2607:ff18:80::1308"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine.

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
