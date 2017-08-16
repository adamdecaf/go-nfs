package nfs

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// https://members.nearlyfreespeech.net/wiki/API/AccountBalance
func GetAccountBalance(c *Client) (float32, error) {
	return balanceRequest(c, fmt.Sprintf("/account/%s/balance", c.accountId))
}

// https://members.nearlyfreespeech.net/wiki/API/AccountBalanceHigh
func GetAccountBalanceHighest(c *Client) (float32, error) {
	return balanceRequest(c, fmt.Sprintf("/account/%s/balanceHigh", c.accountId))
}

func balanceRequest(c *Client, u string) (float32, error) {
	resp, err := c.get(u)
	balance, err := c.readResponse(resp, err)
	if err != nil {
		return 0.0, err
	}

	f, err := strconv.ParseFloat(balance, 32)
	if err != nil {
		return 0.0, err
	}

	return float32(f), err
}

// https://members.nearlyfreespeech.net/wiki/API/AccountFriendlyName
func GetFriendlyName(c *Client) (string, error) {
	u := fmt.Sprintf("/account/%s/friendlyName", c.accountId)
	return c.readResponse(c.get(u))
}
func SetFriendlyName(c *Client, name string) error {
	u := fmt.Sprintf("/account/%s/friendlyName", c.accountId)
	_, err := c.readResponse(c.put(u, name))
	return err
}

// https://members.nearlyfreespeech.net/wiki/API/AccountStatus
func GetAccountStatus(c *Client) (string, error) {
	u := fmt.Sprintf("/account/%s/status", c.accountId)
	return c.readResponse(c.get(u))
}

// https://members.nearlyfreespeech.net/wiki/API/AccountSites
func GetAccountSites(c *Client) ([]string, error) {
	u := fmt.Sprintf("/account/%s/sites", c.accountId)
	resp, err := c.readResponse(c.get(u))
	if err != nil {
		return nil, err
	}

	// Parse json array
	var sites = make([]string, 0)
	err = json.Unmarshal([]byte(resp), &sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}
