package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jezman/gorion/render"
)

// doorlistCmd represents the doorlist command
var doorsCmd = &cobra.Command{
	Use:     "doors",
	Aliases: []string{"d"},
	Short:   "List all doors with ID",
	Run: func(cmd *cobra.Command, args []string) {
		doors, err := Env.Doors()
		if err != nil {
			fmt.Println(err)
		}

		table = render.Preparing(doors, "ID", "Title")
		fmt.Println(table.Render())
	},
}
