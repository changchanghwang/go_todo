package todoService

import (
	model "todo/services/todos/domain"
	"todo/services/todos/dto"
	repository "todo/services/todos/infrastructure"
)

func Add(args dto.TodoCreateDto) {
	todo := model.From(args)

	repository.Save(todo)
}

func List() []model.Todo {
	return repository.Find()
}

func Update(id int, args dto.TodoUpdateDto) {
	todo := repository.FindById(id)
	todo.Update(args)
	repository.Save(*todo)
}
