package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
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

		query := "SELECT GIndex, Name FROM AcessPoint ORDER BY GIndex"
		doors, err := database.Doors(query)
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("ID", "Door")

		for _, d := range doors {
			table.AddRow(d.ID, d.Name)
		}

		fmt.Println(table.Render())
	},
}
