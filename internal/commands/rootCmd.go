package commands

import (
	"fmt"
	"os"

	"github.com/promise111/todo-cli-cobra/internal"
	"github.com/spf13/cobra"
)

var Todos = internal.Todos{}

var rootCmd = cobra.Command{
	Use:   "todo",
	Short: "Manage daily tasks",
	Long:  "Manage todos from the terminal with add, edit, delete, toggle, and list commands.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Manage todos from the terminal with add, edit, delete, toggle, and list commands.")
	},
}

func Execute() {
	executeError := rootCmd.Execute()
	if executeError != nil {
		fmt.Println("Execute Error: ", executeError)
		os.Exit(1)
	}
}
