package todoService

import (
	model "todo/services/todos/domain"
	"todo/services/todos/dto"
	repository "todo/services/todos/infrastructure"
)

func Add(args dto.TodoCreateDto) {
	todo := model.From(args)

	done := make(chan struct{})
	go repository.SaveAsync(todo, done)
	<-done
}

func List() []model.Todo {
	return repository.Find()
}

func Update(id int, args dto.TodoUpdateDto) {
	todo := repository.FindById(id)
	todo.Update(args)

	done := make(chan struct{})
	go repository.SaveAsync(*todo, done)
	<-done
}

func Delete(id int) {
	done := make(chan struct{})
	go repository.DeleteAsync(id, done)
	<-done
}
