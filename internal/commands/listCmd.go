package commands

import (
	"github.com/spf13/cobra"
)

var listCmd = cobra.Command{
	Use: "list",
	Aliases: []string{"print", "list"},
	Long: "List all todos",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string){
		Todos.List()
	},
}

func init () {
	rootCmd.AddCommand(&listCmd)
}
