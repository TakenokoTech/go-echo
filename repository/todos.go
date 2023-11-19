package repository

import (
	"go-echo/domain/todos"
	"go-echo/utils"
	"gorm.io/gorm"
)

type todosRepository struct {
	db *gorm.DB
}

var _ todos.Repository = todosRepository{}

func NewTodosRepository(db *gorm.DB) todos.Repository {
	return todosRepository{
		db: db,
	}
}

func (t todosRepository) SelectAll() (todos []todos.Todo) {
	utils.
		GetTransaction(t.db, nil).
		Find(&todos)
	return
}

func (t todosRepository) Insert(todo todos.Todo) {
	utils.
		GetTransaction(t.db, nil).
		Select("user_id", "task", "status").
		Create(&todo)
}
