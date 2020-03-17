Terraform Mail-in-a-Box Provider
================================

This is an unofficial Mail-in-a-Box Terraform provider. Currently it only supports DNS record resources.

Configuration
-------------

Configuration is done either in the `provider` block or by using environment variables.

Provider block example:

```hcl
provider "mailinabox" {
  url      = "https://mail.example.com"
  username = "admin@example.com"
  password = "abc123fireme"
}
```

Environment variable example:

```shell
export MAILINABOX_URL="https://mail.example.com"
export MAILINABOX_USERNAME="admin@example.com"
export MAILINABOX_PASSWORD="abc123fireme"
```

Usage
-----

```hcl
provider "mailinabox" {}

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
