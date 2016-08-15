package dao

import (
	"testing"
	"os"
	"github.com/gocraft/dbr"
	_ "github.com/go-sql-driver/mysql"
)

var (
	sess *dbr.Session
	messageDao  Message
)

func TestMain(m *testing.M) {
	conn, err := dbr.Open("mysql", "root:@tcp(localhost:3306)/mboard_test", nil)
	if err != nil {
		panic(err)
	}
	sess = conn.NewSession(nil)
	messageDao = NewMessage(sess)

	os.Exit(m.Run())
}

func truncateTables(){
	tables := []string{
		"message",
	}
	for _, table := range tables {
		_, err := sess.Exec("TRUNCATE TABLE " + table)
		if err != nil{
			panic(err)
		}
	}
}