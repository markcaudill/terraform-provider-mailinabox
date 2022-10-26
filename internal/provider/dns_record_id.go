package provider

import (
	"fmt"
	"strings"

	"github.com/markcaudill/gomailinabox"
)

func parseDNSRecordId(s string) (*gomailinabox.Record, error) {
	parts := strings.Split(s, "_")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Unable to parse \"%s\" into gomailinabox.Record.", s)
	}
	return &gomailinabox.Record{Domain: parts[0], Type: parts[1], Value: parts[2]}, nil
}

func generateDNSRecordId(r *gomailinabox.Record) (string, error) {
	parts := []string{r.Domain, r.Type, r.Value}
	// Error if any parts are nil
	for _, p := range parts {
		if p == "" {
			return "", fmt.Errorf("Unable to create id with gomailinabox.Record that has empty values: %+v.", r)
		}
	}
	return strings.Join(parts, "_"), nil
}
