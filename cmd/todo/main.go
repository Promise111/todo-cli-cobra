package main

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	"github.com/promise111/todo-cli-cobra/internal"
	"github.com/promise111/todo-cli-cobra/internal/commands"
)

func main() {
	log.SetPrefix("Todo CLI with cobra: ")
	log.SetFlags(0)
	
	// New Storage
	storage := internal.NewStorage[internal.Todos]("todo.json")
	loadErr := storage.Load(&commands.Todos)
	if loadErr != nil {
		fmt.Println("Load Error: ", loadErr)
		os.Exit(1)
	}
	log.Println("todos length is ", len(commands.Todos))
	commands.Execute()
	saveErr := storage.Save(commands.Todos)
	if saveErr != nil {
		fmt.Println("Save Error: ", saveErr)
		os.Exit(1)
	}
}
