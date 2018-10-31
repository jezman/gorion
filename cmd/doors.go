package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
	"github.com/spf13/cobra"
)

// doorlistCmd represents the doorlist command
var doorsCmd = &cobra.Command{
	Use:     "doors",
	Aliases: []string{"d"},
	Short:   "List all doors with ID",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		doors, err := env.Doors()
		if err != nil {
			fmt.Println(err)
		}

		table = render.Preparing(doors, "ID", "Title")
		fmt.Println(table.Render())
	},
}
