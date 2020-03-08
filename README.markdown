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
export MAININABOX_PASSWORD="abc123fireme"
```

Usage
-----

```hcl
provider "mailinabox" {}

resource "mailinabox_dnsrecord" "test" {
  domain = "testdomain.example.com"
  type   = "A"
  value  = "127.0.0.1"
}
```
