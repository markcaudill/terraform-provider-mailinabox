package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/markcaudill/gomailinabox"
)

func datasourceMailinaboxDNSRecord() *schema.Resource {
	return &schema.Resource{
		Read: datasourceMailinaboxDNSRecordRead,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceMailinaboxDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	// Construct a partial gomailinabox.Record for querying the backend service
	queryRec := &gomailinabox.Record{
		Domain: d.Get("domain").(string),
		Type:   d.Get("type").(string),
	}
	recs, err := client.GetRecord(queryRec)
	if err != nil {
		d.SetId("")
		return err
	}
	if len(recs) < 1 {
		d.SetId("")
		return fmt.Errorf("No matching record found: %+v.", queryRec)
	}
	id, err := generateDNSRecordId(&recs[0])
	if err != nil {
		d.SetId("")
		return err
	}
	d.SetId(id)
	d.Set("domain", recs[0].Domain)
	d.Set("type", recs[0].Type)
	d.Set("value", recs[0].Value)
	return nil
}
