package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/apcera/termtables"
	"github.com/jezman/libgorion"
	"github.com/spf13/cobra"
)

var (
	worker    string
	err       error
	denied    bool
	door      uint
	firstDate string
	lastDate  string
	env       libgorion.Datastore
	timeNow   = time.Now().Local()
	table     *termtables.Table
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "gorion",
	Short: ` _____            _
|  __ \          (_)            
| |  \/ ___  _ __ _  ___  _ __  
| | __ / _ \| '__| |/ _ \| '_ \ 
| |_\ \ (_) | |  | | (_) | | | |
 \____/\___/|_|  |_|\___/|_| |_|
https://github.com/jezman/gorion

Reports view for access control system NVP Bolid 'Orion Pro'

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initDB() (db *libgorion.Database) {
	dsn := os.Getenv("BOLID_DSN")
	if db, err = libgorion.OpenDB(dsn); err != nil {
		panic(err)
	}

	// set environment
	env = db
	return
}
