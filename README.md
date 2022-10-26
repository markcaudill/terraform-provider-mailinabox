# Terraform Mail-in-a-Box Provider

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
