// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package nfs

import (
	"fmt"
)

// https://members.nearlyfreespeech.net/wiki/API/MemberAccounts
func GetMemberAccounts(c *Client) ([]string, error) {
	u := fmt.Sprintf("/member/%s/accounts", c.login)
	resp, err := c.readResponse(c.get(u))
	if err != nil {
		return nil, err
	}
	return readResponseArray(resp)
}

// https://members.nearlyfreespeech.net/wiki/API/MemberSites
func GetMemberSites(c *Client) ([]string, error) {
	u := fmt.Sprintf("/member/%s/sites", c.login)
	resp, err := c.readResponse(c.get(u))
	if err != nil {
		return nil, err
	}
	return readResponseArray(resp)
}
