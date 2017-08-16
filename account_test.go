package nfs

import (
	"testing"
)

// My answers/values, change these if you use your api keys
var (
	minBalance float32 = 0.01
)

func TestNFSAccount__balance(t *testing.T) {
	bal, err := GetAccountBalance(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if bal < minBalance {
		t.Fatal("error - nfs acct balance is low")
	}
}
