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
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Todos.List(showCompleted, showPending)
		if !showCompleted && !showPending {
			fmt.Println("Listing all todos")
		} else if showCompleted {
			fmt.Println("Listing completed todos")
		} else if showPending {
			fmt.Println("Listing pending todos")
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "List completed tasks")
	listCmd.Flags().BoolVarP(&showPending, "pending", "p", false, "List pending tasks")
	rootCmd.AddCommand(&listCmd)
}
