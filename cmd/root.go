package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jezman/gorion/models"
	"github.com/spf13/cobra"
)

var database models.Datastore

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorion",
	Short: "Reports viewing for access control system NVP Bolid 'Orion Pro'",
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
	config := new(models.Config)
	// init connection to the mssql
	db, err := models.OpenDB(config.Read("config.json"))
	if err != nil {
		log.Panic(err)
	}

	// set app connection
	database = db
	return db
}
