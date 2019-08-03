// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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

	testClient = NewClientForAccount(accountId, apiKey, login)
}
