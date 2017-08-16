package nfs

import (
	"testing"
)

func TestNFSMember__accounts(t *testing.T) {
	acts, err := GetMemberAccounts(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	if len(acts) < 0 {
		t.Fatal("no accounts for member")
	}
}
