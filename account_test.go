package nfs

import (
	"testing"
)

// My answers/values, change these if you use your api keys
var (
	minBalance float32 = 0.01
	friendlyName = "adam"
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

func TestNFSAccount__balanceHigh(t *testing.T) {
	bal, err := GetAccountBalanceHighest(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if bal < minBalance {
		t.Fatal("error - nfs acct balance is low")
	}
}

func TestNFSAccount__friendlyName(t *testing.T) {
	fn, err := GetFriendlyName(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	if fn != friendlyName {
		t.Fatalf("error - nfs acct friendlyName is wrong: %s", fn)
	}

	err = SetFriendlyName(testClient, "other")
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	SetFriendlyName(testClient, friendlyName)
}
