package nfs

import (
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
