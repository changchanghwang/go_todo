package repository

import (
	model "todo/services/todos/domain"
)

var todoList = []model.Todo{}

func Find() []model.Todo {
	return todoList
}

func FindById(id int) *model.Todo {
	for _, item := range todoList {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func Save(todo model.Todo) {
	isExist := some(todoList, func(item model.Todo) bool {
		return item.Id == todo.Id
	})
	if !isExist {
		todoList = append(todoList, todo)
		return
	}
	for i, v := range todoList {
		if v.Id == todo.Id {
			todoList[i] = todo
		}
	}
}

func Add(todo model.Todo) {
	todoList = append(todoList, todo)
}

func some(arr []model.Todo, condition func(model.Todo) bool) bool {
	for _, item := range arr {
		if condition(item) {
			return true
		}
	}
	return false
}
