package internal

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/promise111/todo-cli-cobra/internal/utils"
	"github.com/liamg/tml"
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

func (todos *Todos) Toggle(index int) error {
	t := *todos
	err := utils.ValidateTodos(len(*todos), index)
	if err != nil {
		return err
	}

	if !t[index].Completed {
		now := time.Now()
		t[index].Completed = true
		t[index].CompletedAt = &now
	} else {
		t[index].Completed = false
		t[index].CompletedAt = nil
	}

	return nil
}

func (todos *Todos) List() {

	tbl := table.New(os.Stdout)

	tbl.AddHeaders("s/n", "title", "completed", "createdAt", "completedAt")
	tbl.SetRowLines(false)

	for i, todo := range *todos {
		sn := strconv.Itoa(i)
		title := todo.Title
		completed := "❌"
		completedAt := ""
		if todo.Completed == true {
			completed = "✅"
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}
		createdAt := todo.CreatedAt.Format(time.RFC1123)
		tbl.AddRow(sn, tml.Sprintf("<cyan>%v</cyan>",title), completed, createdAt, completedAt)
	}

	tbl.Render()

}
