package domain

import (
	"math/rand"
	"todo/services/todos/dto"
)

type Todo struct {
	Id          int
	Title       string
	Description string
}

func From(args dto.TodoCreateDto) Todo {
	return Todo{
		Id:          rand.Int(),
		Title:       args.Title,
		Description: args.Description,
	}
}
