package main

import (
	// "fmt"
	"log"

	"github.com/promise111/todo-cli-cobra/internal"
	"github.com/promise111/todo-cli-cobra/internal/commands"
)

func main() {
	log.SetPrefix("Todo CLI with cobra: ")
	log.SetFlags(0)
	
	// New Storage
	storage := internal.NewStorage[internal.Todos]("todo.json")
	storage.Load(&commands.Todos)
	log.Println("todos length is ", len(commands.Todos))
	commands.Execute()
	storage.Save(commands.Todos)
}
