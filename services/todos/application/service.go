package application

import (
	"todo/services/todos/domain"
	"todo/services/todos/infrastructure"
)

func Add(title string, description string) {
	todo := domain.Todo{
		Title:       title,
		Description: description,
	}

	infrastructure.Add(todo)
}

func List() []domain.Todo {
	return infrastructure.Find()
}
