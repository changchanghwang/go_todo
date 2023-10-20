package infrastructure

import "todo/services/todos/domain"

var a = []domain.Todo{}

func Find() []domain.Todo {
	return a
}

func Add(todo domain.Todo) {
	a = append(a, todo)
}
