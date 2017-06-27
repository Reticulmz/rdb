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

func TestGetUserByUsername(t *testing.T) {
	prepare()
	user := rdb.GetUserByUsername("FokaBot")
	if user == nil {
		t.Fatal()
	}
	if user.ID != 999 {
		t.Fatal("FokaBot does not return 999 but", user.ID)
	}
}

func TestGetUserByUsernameSafe(t *testing.T) {
	prepare()
	user := rdb.GetUserByUsernameSafe("[_potato_]")
	if user == nil {
		t.Fatal()
	}
	if user.ID != 1003 {
		t.Fatal("[_potato_] does not return 1003 but", user.ID)
	}
}

func TestGetUserByEmail(t *testing.T) {
	prepare()
	user := rdb.GetUserByEmail("nicememe@google.com")
	if user == nil {
		t.Fatal()
	}
	if user.ID != 1002 {
		t.Fatal("nicememe@google.com does not return 1002 but", user.ID)
	}
}

func TestGetFakeUser(t *testing.T) {
	usr := rdb.GetUser(-42)
	if usr != nil {
		t.Fatal("-42 returned", usr)
	}
}
