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

func SaveAsync(todo model.Todo, done chan struct{}) {
	isExist := some(todoList, func(item model.Todo) bool {
		return item.Id == todo.Id
	})
	if !isExist {
		todoList = append(todoList, todo)
		done <- struct{}{}
		return
	}
	for i, v := range todoList {
		if v.Id == todo.Id {
			todoList[i] = todo
		}
	}
	done <- struct{}{}
}

func DeleteAsync(id int, done chan struct{}) {
	for i, v := range todoList {
		if v.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}
	done <- struct{}{}
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
