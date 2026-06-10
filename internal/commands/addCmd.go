package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use: "add",
	Aliases: []string{"enter"},
	Short: "Add new todo",
	Long: "Add new todo to existing slice of todos",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		Todos.Add(args[0])
		fmt.Printf("Task added: %s", args[0])
	},
}

func init () {
	rootCmd.AddCommand(&addCmd)
}