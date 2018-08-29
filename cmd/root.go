package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jezman/gorion/models"
	"github.com/spf13/cobra"
)

var (
	employee  string
	err       error
	door      uint
	firstDate string
	lastDate  string
	env       models.Datastore
	timeNow   = time.Now().Local()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorion",
	Short: `
 _____            _             
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

func initDB() (db *models.DB) {
	// read env var
	dsn := os.Getenv("BOLID_DSN")
	// init connection to the mssql
	if db, err = models.OpenDB(dsn); err != nil {
		panic(err)
	}

	// set environment
	env = db
	return
}
