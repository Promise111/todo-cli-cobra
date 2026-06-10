package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var toggleCmd = cobra.Command{
	Use:     "toggle",
	Aliases: []string{"toggle"},
	Short:   "Toggle a todo",
	Long:    "Toggle a todo by index",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid index. Use a number.")
			os.Exit(1)
		}
		toggErr := Todos.Toggle(index)
		if toggErr != nil {
			fmt.Println("Toggle Error: ", toggErr)
			os.Exit(1)
		}
		fmt.Printf("Task toggled: %s", args[0])
	},
}

func init() {
	rootCmd.AddCommand(&toggleCmd)
}
