---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mailinabox_dnsrecord Resource - terraform-provider-mailinabox"
subcategory: ""
description: |-
  A DNS record on a Mail-in-a-box server
---

# mailinabox_dnsrecord (Resource)

A DNS record on a Mail-in-a-box server

## Example Usage

```terraform
resource "mailinabox_dnsrecord" "testdomain_a_dnsrecord" {
  domain = "testdomain.example.com"
  type   = data.mailinabox_dnsrecord.bare_a_record.type
  value  = data.mailinabox_dnsrecord.bare_a_record.value
}

resource "mailinabox_dnsrecord" "testdomain_aaaa_dnsrecord" {
  domain = "testdomain.example.com"
  type   = "AAAA"
  value  = "2607:ff18:80::1308"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String)
- `type` (String)
- `value` (String)

### Read-Only

- `id` (String) The ID of this resource.

