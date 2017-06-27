package rdb_test

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/osuripple/rdb"
	"gopkg.in/testfixtures.v2"
)

var fixtures *testfixtures.Context

func TestMain(m *testing.M) {
	// open mysql database connection
	err := rdb.OpenDatabase("mysql", getenvDef("DSN", "root@/test?multiStatements=true"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rdb.DB = rdb.DB.Debug()

	// only for testing purposes, this creates the tables we need (possibly different than
	// the ones created using bdzr)
	rdb.DB.AutoMigrate(
		&rdb.User{},
	)

	fixtures, err = testfixtures.NewFolder(rdb.DB.DB(), &testfixtures.MySQL{}, "fixtures")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := m.Run()

	rdb.DB.Exec("DROP DATABASE test")
	rdb.DB.Exec("CREATE DATABASE test")

	os.Exit(res)
}

func getenvDef(s string, def string) string {
	x := os.Getenv(s)
	if x == "" {
		return def
	}
	return x
}

func prepare() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
