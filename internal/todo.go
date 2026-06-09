package internal

import (
	"time"
	"github.com/promise111/todo-cli-cobra/internal/utils"
)

type Todo struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"CompletedAt"`
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, newTodo)
}

func (todos *Todos) Edit(index int, title string) error {
	err := utils.ValidateTodos(len(*todos), index)
	if err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) Delete(index int) error {
	err := utils.ValidateTodos(len(*todos), index)
	if err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}
