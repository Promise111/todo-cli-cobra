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

	// initiate cobra on startup
	var todos = internal.Todos{}
	// New Storage
	storage := internal.NewStorage[internal.Todos]("todo.json")
	storage.Load(&todos)
	// todos.Add("Buy bread")
	// todos.Add("Bake some meat pies")
	// todos.Edit(1,"Sky dive later this morning")
	// todos.Toggle(1)
	todos.List()
	// fmt.Printf("%+v \n\n", todos)
	storage.Save(todos)
	commands.Execute()
}
