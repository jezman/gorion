package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/scyllabd/termtables"
	"github.com/jezman/libgorion"
	"github.com/spf13/cobra"
)

// VERSION application
const VERSION = "0.1.0"

var (
	worker    string
	err       error
	denied    bool
	door      uint
	firstDate string
	lastDate  string
	Env       libgorion.Datastore
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
