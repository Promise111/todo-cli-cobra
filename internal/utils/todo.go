package utils

import (
	"errors"
)

func ValidateTodos(todosLength int, index int) error {
	if index < 0 || index >= todosLength {
		return errors.New("Error: Invalid index")
	}
	return nil
}

func FormatTMLColor(color string) string {
	return "<" + color + ">%s</" + color + ">"
}
