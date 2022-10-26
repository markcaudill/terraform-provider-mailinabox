package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/markcaudill/gomailinabox"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
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
			DataSourcesMap: map[string]*schema.Resource{
				"mailinabox_dnsrecord": datasourceMailinaboxDNSRecord(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"mailinabox_dnsrecord": resourceMailinaboxDNSRecord(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(c context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		config := gomailinabox.Config{
			URL:      d.Get("url").(string),
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
		}
		return &config, nil
	}
}
