package main

type Todo struct {
}

type TodoPagination struct {
}

type PaginationInput struct {
}

type TodoSearchInput struct {
}

type TodoInput struct {
}

type TodoService interface {
	GetTodo(int) (Todo, error)
	TodoPages(PaginationInput) (TodoPagination, error)
	TodoSearch(TodoSearchInput, PaginationInput) (TodoPagination, error)
	CreateTodo(TodoInput) (Todo, error)
	UpdateTodo(int, TodoInput) (Todo, error)
}
