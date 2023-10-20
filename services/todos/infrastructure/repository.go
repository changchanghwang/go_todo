package repository

import model "todo/services/todos/domain"

var a = []model.Todo{}

func Find() []model.Todo {
	return a
}

func Add(todo model.Todo) {
	a = append(a, todo)
}
