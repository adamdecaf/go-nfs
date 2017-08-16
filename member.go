package nfs

import (
	"fmt"
)

// GET
// accounts - https://members.nearlyfreespeech.net/wiki/API/MemberAccounts
// sites - https://members.nearlyfreespeech.net/wiki/API/MemberSites

func GetMemberAccounts(c *Client) ([]string, error) {
	u := fmt.Sprintf("/member/%s/accounts", c.login)
	resp, err := c.readResponse(c.get(u))
	if err != nil {
		return nil, err
	}
	return readResponseArray(resp)
}
