package nfs

import (
	"os"
)

var (
	testClient *Client
)

func init() {
	accountId := os.Getenv("NFS_ACCOUNT_ID")
	apiKey := os.Getenv("NFS_API_KEY")
	login := os.Getenv("NFS_LOGIN")

	if (accountId == "" || apiKey == "") || login == "" {
		panic("error - no NFS_* env variables set")
	}

	testClient = NewClient(accountId, apiKey, login)
}
