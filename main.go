package main

import (
	_ "github.com/denisenkom/go-mssqldb"

	"github.com/jezman/gorion/cmd"
)

func main() {
	cmd.Execute()
}
