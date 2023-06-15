package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olivo4ka37/testRestApi"
)

type Authorization interface {
	CreateUser(user testRestApi.User) (int, error)
	GetUser(username, password string) (testRestApi.User, error)
}

type TodoList interface {
	Create(userId int, list testRestApi.TodoList) (int, error)
	GetAll(userId int) ([]testRestApi.TodoList, error)
	GetById(userId, listId int) (testRestApi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input testRestApi.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item testRestApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]testRestApi.TodoItem, error)
	GetById(userId, itemId int) (testRestApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input testRestApi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
