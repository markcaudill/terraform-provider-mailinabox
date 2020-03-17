package main

import (
	"github.com/markcaudill/gomailinabox"
	"testing"
)

func TestParseDNSRecordId(t *testing.T) {
	// Good input
	s := "testdomain.example.com_A_1.1.1.1"
	got, err := parseDNSRecordId(s)
	// err should be nil
	if err != nil {
		t.Errorf("parseDNSRecordId(\"%s\") = %+v, %+v", s, got, err)
	}
	// got.Domain should be set properly
	if got.Domain != "testdomain.example.com" {
		t.Errorf("got.Domain = \"%s\"; wanted \"testdomain.example.com\"", got.Domain)
	}
	// got.Type should be set properly
	if got.Type != "A" {
		t.Errorf("got.Type = \"%s\"; wanted \"A\"", got.Type)
	}
	// got.Value should be set properly
	if got.Value != "1.1.1.1" {
		t.Errorf("got.Value = \"%s\"; wanted \"1.1.1.1\"", got.Value)
	}

	// Bad input
	s = "testdomain.example.com_1.1.1.1"
	got, err = parseDNSRecordId(s)
	// err should not be nil
	if err == nil {
		t.Errorf("parseDNSRecordId(\"%s\") = %+v, %+v; wanted nil, error", s, got, err)
	}
}

func TestGenerateDNSRecordId(t *testing.T) {
	// Good input
	r := &gomailinabox.Record{
		Domain: "testdomain.example.com",
		Type:   "A",
		Value:  "1.1.1.1",
	}
	got, err := generateDNSRecordId(r)
	// err should be nil
	if err != nil {
		t.Errorf("generateDNSRecordId(%+v) = \"%s\", %+v", r, got, err)
	}
	// got should be the id string
	if got != "testdomain.example.com_A_1.1.1.1" {
		t.Errorf("generateDNSRecordId(%+v) = \"%s\"; wanted \"testdomain.example.com_A_1.1.1.1\"", r, got)
	}

	//Bad input
	r = &gomailinabox.Record{
		Domain: "testdomain.example.com",
		Type:   "A",
	}
	got, err = generateDNSRecordId(r)
	// err should not be nil
	if err == nil {
		t.Errorf("recordId(%+v) = \"%s\", %+v; wanted \"\", error", r, got, err)
	}
}
