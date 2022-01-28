package main

import (
	"time"
)

type PaginationInput struct {
	Count int
	Page  int
}

type TodoSearchInput struct {
	ContentCode *string
	ContentName *string
	StartDate   *time.Time
	EndDate     *time.Time
	Status      *string
}

type TodoInput struct {
	ContentCode *string
	ContentName *string
	Description *string
	StartDate   *time.Time
	EndDate     *time.Time
	Status      *string
	CreatedBy   *string
	UpdatedBy   *string
}

type PaginationInfo struct {
	TotalCount  int
	CurrentPage int
	TotalPages  int
}

type TodoPagination struct {
	PaginationInfo PaginationInfo
	Todos          []Todo
}

type Todo struct {
	Id          int
	ContentCode string
	ContentName string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Status      string
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
}

type TodoService interface {
	GetTodo(int) (Todo, error)
	TodoPages(PaginationInput) (TodoPagination, error)
	TodoSearch(TodoSearchInput, PaginationInput) (TodoPagination, error)
	CreateTodo(TodoInput) (Todo, error)
	UpdateTodo(int, TodoInput) (Todo, error)
}
