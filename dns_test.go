package nfs

import (
	"fmt"
	"testing"
)

// My answers/values, change these if you use your api keys
var (
	domain = "yourdomain.com"
)

func TestDNS_ListRecords(t *testing.T) {
	p := make(map[string]string)
	records, err := GetDNSRecords(testClient, domain, p)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	fmt.Print(records)
}

func TestDNS_AddRecord(t *testing.T) {
	p := map[string]string{
		"name": "terraformtest",
		"type": "TXT",
		"data": "TXT Record For Addition",
	}
	records, err := SetDNSRecord(testClient, domain, p)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	fmt.Print(records)
}

func TestDNS_RemoveRecord(t *testing.T) {
	p := map[string]string{
		"name": "terraformtest",
		"type": "TXT",
		"data": "TXT Record For Removal",
	}
	records, err := RemoveDNSRecord(testClient, domain, p)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	fmt.Print(records)
}
