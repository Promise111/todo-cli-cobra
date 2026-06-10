package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCompleted bool
var showPending bool
var listCmd = cobra.Command{
	Use:     "list",
	Aliases: []string{"print", "list"},
	Long:    "List all todos",
	Short:   "List all todos",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(showCompleted, showPending)
		// Todos.List(showCompleted, showPending)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "List completed tasks")
	listCmd.Flags().BoolVarP(&showPending, "pending", "p", false, "List pending tasks")
	rootCmd.AddCommand(&listCmd)
}
