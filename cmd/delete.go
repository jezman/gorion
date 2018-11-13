package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete workers from access control system",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()
		
		if err := env.DeleteWorker(worker); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("worker successfully delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker full name(first mid last)")
}
