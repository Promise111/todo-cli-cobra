package commands

import (
	"github.com/spf13/cobra"
)

var completed bool
var pending bool
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
	toggleCmd.Flags().BoolVarP(&completed, "completed", "c", false, "List completed tasks")
	toggleCmd.Flags().BoolVarP(&pending, "pending", "p", false, "List pending todos")
	rootCmd.AddCommand(&listCmd)
}
