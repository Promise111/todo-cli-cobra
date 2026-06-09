package utils

import (
	"errors"
	"github.com/promise111/todo-cli-cobra/internal"
)

func ValidateTodos (todos *internal.Todos,index int) error {
	if index < 0 || index > len(*todos) {
		return errors.New("Error: Invalid index")
	}
	return nil
}
