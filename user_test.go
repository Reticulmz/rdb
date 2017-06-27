package rdb_test

import (
	"testing"

	"github.com/osuripple/rdb"
)

func TestGetUser(t *testing.T) {
	prepare()
	user := rdb.GetUser(999)
	if user == nil {
		t.Fatal()
	}
	if user.Username != "FokaBot" {
		t.Fatal("999 does not return fokabot but", user.Username)
	}
}
