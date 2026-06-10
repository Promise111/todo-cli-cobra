package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = cobra.Command{
	Use: "edit",
	Aliases: []string{"update"},
	Short: "Edit a todo",
	Long: "Edit a todo by index and title",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		splitArgs := strings.SplitN(args[0], ":", 2)
		if len(splitArgs) != 2 {
			fmt.Println("Error: Invalid input format. Use 'add:title'")
			os.Exit(1)
		}
		index, err := strconv.Atoi(splitArgs[0])
		if err != nil {
			fmt.Println("Error: Invalid index. Use a number.")
			os.Exit(1)
		}
		if splitArgs[1] == "" {
			fmt.Println("Error: Title is required.")
			os.Exit(1)
		}
		Todos.Edit(index, splitArgs[1])
		fmt.Printf("Task edited: %s", args[0])
	},
}

func init () {
	rootCmd.AddCommand(&editCmd)
}