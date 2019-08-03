// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package nfs

import (
	"testing"
)

// My answers/values, change these if you use your api keys
var (
	minBalance     float32 = 0.01
	friendlyName           = "adam"
	minSites               = 0
	balanceWarning float32 = 10.00
)

func TestNFSAccount__balance(t *testing.T) {
	testClient := createTestClient(t)
	bal, err := GetAccountBalance(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if bal < minBalance {
		t.Fatal("error - nfs acct balance is low")
	}
}

func TestNFSAccount__balanceHigh(t *testing.T) {
	testClient := createTestClient(t)
	bal, err := GetAccountBalanceHighest(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if bal < minBalance {
		t.Fatal("error - nfs acct balance is low")
	}
}

func TestNFSAccount__warnings(t *testing.T) {
	testClient := createTestClient(t)
	err := AddBalanceWarning(testClient, balanceWarning)
	if err != nil {
		t.Fatalf("error adding balance warning - %s", err)
	}

	err = RemoveBalanceWarning(testClient, balanceWarning)
	if err != nil {
		t.Fatalf("error removing balance warning - %s", err)
	}
}

func TestNFSAccount__friendlyName(t *testing.T) {
	testClient := createTestClient(t)
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

func TestNFSAccount__status(t *testing.T) {
	testClient := createTestClient(t)
	s, err := GetAccountStatus(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if s != "Ok" {
		t.Fatalf("error - account status != Ok, got %s", s)
	}
}

func TestNFSAccount__sites(t *testing.T) {
	testClient := createTestClient(t)
	sites, err := GetAccountSites(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}

	// TODO(adam): When I have sites, set len(sites) > 1 check
	if len(sites) < minSites {
		t.Fatal("error - not enough sites listed")
	}
}
