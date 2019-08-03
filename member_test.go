// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package nfs

import (
	"testing"
)

func TestNFSMember__accounts(t *testing.T) {
	testClient := createTestClient(t)
	acts, err := GetMemberAccounts(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	if len(acts) < 0 {
		t.Fatal("no accounts for member")
	}
}

func TestNFSMember__sites(t *testing.T) {
	testClient := createTestClient(t)
	sites, err := GetMemberSites(testClient)
	if err != nil {
		t.Fatalf("error - %s", err)
	}
	if len(sites) < minSites {
		t.Fatal("not enough sites member")
	}
}
