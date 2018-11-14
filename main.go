package main

import (
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/jezman/gorion/cmd"
	"github.com/jezman/libgorion"
)

func main() {
	args := os.Args

	if len(args) > 1 && args[1] != "vesrion" && args[1] != "help" {
		dsn := os.Getenv("BOLID_DSN")
		db, err := libgorion.OpenDB(dsn)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		cmd.Env = db
	}

	cmd.Execute()
}
