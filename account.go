// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package nfs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//GetAccountBalance https://members.nearlyfreespeech.net/wiki/API/AccountBalance
func GetAccountBalance(c *Client) (float32, error) {
	return balanceRequest(c, fmt.Sprintf("/account/%s/balance", c.accountId))
}

//GetAccountBalanceHighest https://members.nearlyfreespeech.net/wiki/API/AccountBalanceHigh
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

//AddBalanceWarning https://members.nearlyfreespeech.net/wiki/API/AccountAddWarning
func AddBalanceWarning(c *Client, bal float32) error {
	u := fmt.Sprintf("/account/%s/addWarning", c.accountId)
	params := make(map[string]string)
	params["balance"] = fmt.Sprintf("%.2f", bal)

	return c.checkErrors(c.post(u, params))
}

//RemoveBalanceWarning https://members.nearlyfreespeech.net/wiki/API/AccountRemoveWarning
func RemoveBalanceWarning(c *Client, bal float32) error {
	u := fmt.Sprintf("/account/%s/removeWarning", c.accountId)
	params := make(map[string]string)
	params["balance"] = fmt.Sprintf("%.2f", bal)

	return c.checkErrors(c.post(u, params))
}

//GetFriendlyName https://members.nearlyfreespeech.net/wiki/API/AccountFriendlyName
func GetFriendlyName(c *Client) (string, error) {
	u := fmt.Sprintf("/account/%s/friendlyName", c.accountId)
	return c.readResponse(c.get(u))
}
func SetFriendlyName(c *Client, name string) error {
	u := fmt.Sprintf("/account/%s/friendlyName", c.accountId)
	return c.checkErrors(c.put(u, name))
}

// https://members.nearlyfreespeech.net/wiki/API/AccountStatus
type accountStatus struct {
	Status string `json:"status"`
}

func GetAccountStatus(c *Client) (string, error) {
	u := fmt.Sprintf("/account/%s/status", c.accountId)
	s, err := c.readResponse(c.get(u))
	if err != nil {
		return "", err
	}

	d := json.NewDecoder(strings.NewReader(s))
	var st accountStatus
	err = d.Decode(&st)
	if err != nil {
		return "", err
	}
	return st.Status, nil
}

// https://members.nearlyfreespeech.net/wiki/API/AccountSites
func GetAccountSites(c *Client) ([]string, error) {
	u := fmt.Sprintf("/account/%s/sites", c.accountId)
	resp, err := c.readResponse(c.get(u))
	if err != nil {
		return nil, err
	}
	return readResponseArray(resp)
}
