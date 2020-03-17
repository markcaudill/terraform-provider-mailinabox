package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/markcaudill/gomailinabox"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAILINABOX_URL", nil),
				Description: "The mailinabox base URL (e.g. https://mail.example.com).",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAILINABOX_USERNAME", nil),
				Description: "The email address of an admin user.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAILINABOX_PASSWORD", nil),
				Description: "The password of an admin user.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mailinabox_dnsrecord": resourceMailinaboxDNSRecord(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := gomailinabox.Config{
		URL:      d.Get("url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return &config, nil
}
