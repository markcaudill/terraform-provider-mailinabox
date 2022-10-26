package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/markcaudill/gomailinabox"
)

func resourceMailinaboxDNSRecord() *schema.Resource {
	return &schema.Resource{
		Description: "A DNS record on a Mail-in-a-box server",

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
	// Create a new gomailinabox.Record
	newRecord := &gomailinabox.Record{
		Domain: d.Get("domain").(string),
		Type:   d.Get("type").(string),
		Value:  d.Get("value").(string),
	}
	// Create the record
	recs, err := client.CreateRecord(newRecord)
	if err != nil {
		return err
	}
	// Generate an id
	id, err := generateDNSRecordId(&recs[0])
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceMailinaboxDNSRecordRead(d, m)
}

func resourceMailinaboxDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	// Parse the id into a gomailinabox.Record
	queryRecord, err := parseDNSRecordId(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	// Fetch any records that match queryRecord
	recs, err := client.GetRecord(queryRecord)
	if err != nil {
		d.SetId("")
		return err
	}
	// Update the values of this resource
	d.Set("domain", recs[0].Domain)
	d.Set("type", recs[0].Type)
	d.Set("value", recs[0].Value)
	return nil
}

func resourceMailinaboxDNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	// Parse this id into a gomailinabox.Record
	record, err := parseDNSRecordId(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	// Update the record value
	record.Value = d.Get("value").(string)
	recs, err := client.UpdateRecord(record)
	if err != nil {
		d.SetId("")
		return err
	}
	// Generate the new id
	id, err := generateDNSRecordId(&recs[0])
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceMailinaboxDNSRecordRead(d, m)
}

func resourceMailinaboxDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	client := gomailinabox.NewClient(m.(*gomailinabox.Config))
	// Parse this id into a gomailinabox.Record
	record, err := parseDNSRecordId(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	// Delete the record
	_, err = client.DeleteRecord(record)
	if err != nil {
		d.SetId("")
		return err
	}
	return nil
}
