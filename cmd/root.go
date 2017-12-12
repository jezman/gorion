package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jezman/gorion/models"
	"github.com/spf13/cobra"
)

var (
	employee  string
	door      uint
	firstDate string
	lastDate  string
	database  models.Datastore
	timeNow   = time.Now().Local()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorion",
	Short: "Reports view for access control system NVP Bolid 'Orion Pro'",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initDB() *models.DB {
	// read env var
	dsn := os.Getenv("BOLID_DSN")
	// init connection to the mssql
	db, err := models.OpenDB(dsn)
	if err != nil {
		log.Panic(err)
	}

	// set app connection
	database = db
	return db
}
