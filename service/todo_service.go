package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-todo-list/exception"
	"golang-todo-list/helper"
	"golang-todo-list/model/domain"
	"golang-todo-list/model/web"
	"golang-todo-list/repository"
	"log"
)

type TodoService struct {
	TodoRepository *repository.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository *repository.TodoRepository, db *sql.DB, validate *validator.Validate) *TodoService {
	return &TodoService{
		TodoRepository: todoRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *TodoService) GetAllTodo(ctx context.Context) []domain.Todo {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todos, err := service.TodoRepository.GetAllTodo(ctx, tx)
	helper.PanicIfError(err)

	return todos
}

func (service *TodoService) GetTodoById(ctx context.Context, id int) domain.Todo {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.GetOneTodo(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoService) SaveTodo(ctx context.Context, request web.TodoCreateRequest) domain.Todo {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		ActivityGroupID: request.ActivityGroupId,
		Title:           request.Title,
		Priority:        request.Priority,
	}

	todo, err = service.TodoRepository.Create(ctx, tx, todo)
	helper.PanicIfError(err)

	return todo
}

func (service *TodoService) UpdateTodo(ctx context.Context, request web.TodoUpdateRequest) domain.Todo {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.GetOneTodo(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todo.Title = request.Title
	todo, err = service.TodoRepository.Update(ctx, tx, todo)
	helper.PanicIfError(err)

	return todo
}

func (service *TodoService) DeleteTodo(ctx context.Context, id int) {
	log.Println("Masuk service")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.GetOneTodo(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}
