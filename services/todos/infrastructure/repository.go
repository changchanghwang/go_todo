package repository

import (
	"todo/lib/datasource"
	model "todo/services/todos/domain"

	"gorm.io/gorm/clause"
)

var todoList = []model.Todo{}

func Find() []model.Todo {
	results := []model.Todo{}
	datasource.DB.Model(&model.Todo{}).Find(&results)
	return results
}

func FindById(id int) model.Todo {
	result := model.Todo{}
	datasource.DB.Model(&model.Todo{Id: id}).First(&result)
	return result
}

func Save(todo model.Todo) {
	datasource.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&todo)
}

func Delete(id int) {
	datasource.DB.Delete(&model.Todo{}, id)
}
