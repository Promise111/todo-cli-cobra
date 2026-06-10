package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use: "delete",
	Aliases: []string{"remove"},
	Short: "Delete todo",
	Long: "Delete a todo by index",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid index. Use a number.")
			os.Exit(1)
		}
		Todos.Delete(index)
		fmt.Printf("Task added: %s", args[0])
	},
}

func init () {
	editCmd.AddCommand(&deleteCmd)
}