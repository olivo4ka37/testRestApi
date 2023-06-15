package service

import (
	"github.com/olivo4ka37/testRestApi"
	"github.com/olivo4ka37/testRestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user testRestApi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list testRestApi.TodoList) (int, error)
	GetAll(userId int) ([]testRestApi.TodoList, error)
	GetById(userId, listId int) (testRestApi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input testRestApi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item testRestApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]testRestApi.TodoItem, error)
	GetById(userId, itemId int) (testRestApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input testRestApi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
