package service

import "github.com/hurucik/TODO_app/pckg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repose *repository.Repository) *Service {
	return &Service{}
}
