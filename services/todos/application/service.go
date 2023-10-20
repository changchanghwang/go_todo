package application

import (
	model "todo/services/todos/domain"
	"todo/services/todos/dto"
	repository "todo/services/todos/infrastructure"
)

func Add(args dto.TodoCreateDto) {
	todo := model.From(args)

	repository.Add(todo)
}

func List() []model.Todo {
	return repository.Find()
}
