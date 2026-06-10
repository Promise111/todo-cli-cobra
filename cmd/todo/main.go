package main

import (
	// "fmt"
	"log"

	"github.com/promise111/todo-cli-cobra/internal"
	"github.com/promise111/todo-cli-cobra/internal/commands"
)

func main() {
	log.SetPrefix("Todo CLI with cobra: ")
	log.SetFlags(1)

	// New Storage
	storage := internal.NewStorage[internal.Todos]("todo.json")
	storage.Load(&commands.Todos)
	storage.Save(commands.Todos)
	commands.Execute()
}
