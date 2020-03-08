package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/markcaudill/gomailinabox"
	"strings"
)

func resourceMailinaboxDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceMailinaboxDNSRecordCreate,
		Read:   resourceMailinaboxDNSRecordRead,
		Update: resourceMailinaboxDNSRecordUpdate,
		Delete: resourceMailinaboxDNSRecordDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMailinaboxDNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	newRecord := &gomailinabox.Record{Domain: d.Get("domain").(string), Type: d.Get("type").(string), Value: d.Get("value").(string)}
	recs, err := client.CreateRecord(newRecord)
	if err != nil {
		return err
	}
	d.SetId(recs[0].Domain + "_" + recs[0].Type + "_" + recs[0].Value)
	return resourceMailinaboxDNSRecordRead(d, m)
}

func resourceMailinaboxDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	i := strings.Split(d.Id(), "_")
	recs, err := client.GetRecord(&gomailinabox.Record{Domain: i[0], Type: i[1], Value: i[2]})
	if err != nil {
		d.SetId("")
	} else {
		d.Set("domain", recs[0].Domain)
		d.Set("type", recs[0].Type)
		d.Set("value", recs[0].Value)
	}
	return nil
}

func resourceMailinaboxDNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	i := strings.Split(d.Id(), "_")
	recs, err := client.UpdateRecord(&gomailinabox.Record{Domain: i[0], Type: i[1], Value: d.Get("value").(string)})
	if err != nil {
		d.SetId("")
	} else {
		d.SetId(recs[0].Domain + "_" + recs[0].Type + "_" + recs[0].Value)
	}
	return resourceMailinaboxDNSRecordRead(d, m)
}

func resourceMailinaboxDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	i := strings.Split(d.Id(), "_")
	_, err := client.DeleteRecord(&gomailinabox.Record{Domain: i[0], Type: i[1], Value: i[2]})
	if err == nil {
		d.SetId("")
	}
	return nil
}
