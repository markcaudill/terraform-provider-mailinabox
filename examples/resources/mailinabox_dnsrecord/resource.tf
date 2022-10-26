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
