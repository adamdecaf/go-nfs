package nfs

import (
	"fmt"
)

//SetDNSRecord https://members.nearlyfreespeech.net/wiki/API/DNSAddRR
func SetDNSRecord(c *Client, domain string, params map[string]string) (string, error) {
	request := fmt.Sprintf("/dns/%s/addRR", domain)
	return dnsRequest(c, request, params)
}

//RemoveDNSRecord https://members.nearlyfreespeech.net/wiki/API/DNSRemoveRR
func RemoveDNSRecord(c *Client, domain string, params map[string]string) (string, error) {
	request := fmt.Sprintf("/dns/%s/removeRR", domain)
	return dnsRequest(c, request, params)
}

//GetDNSRecords https://members.nearlyfreespeech.net/wiki/API/
func GetDNSRecords(c *Client, domain string, params map[string]string) (string, error) {
	request := fmt.Sprintf("/dns/%s/listRRs", domain)
	return dnsRequest(c, request, params)
}

func dnsRequest(c *Client, u string, m map[string]string) (string, error) {
	resp, err := c.post(u, m)
	dnsResponse, err := c.readResponse(resp, err)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return dnsResponse, err
}

//TODO: updateSerial endpoint not implemented
