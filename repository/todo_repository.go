package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-todo-list/helper"
	"golang-todo-list/model/domain"
	"strconv"
)

type TodoRepository struct {
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (repository *TodoRepository) GetAllTodo(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error) {
	query := "SELECT * FROM todos order by id"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	helper.PanicIfError(err)

	var todos []domain.Todo
	for rows.Next() {
		var todo domain.Todo
		err := rows.Scan(&todo.ID, &todo.ActivityGroupID, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		helper.PanicIfError(err)

		todos = append(todos, todo)
	}
	return todos, nil
}

func (repository *TodoRepository) GetOneTodo(ctx context.Context, tx *sql.Tx, ID int) (domain.Todo, error) {
	query := "SELECT * FROM todos WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, ID)
	defer rows.Close()
	helper.PanicIfError(err)

	var todo domain.Todo
	if rows.Next() {
		err := rows.Scan(&todo.ID, &todo.ActivityGroupID, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		helper.PanicIfError(err)
	} else {
		errString := "Todo with ID " + strconv.Itoa(ID) + " Not Found"
		return todo, errors.New(errString)
	}

	return todo, nil
}

func (repository *TodoRepository) Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error) {
	query := "INSERT INTO todos(activity_group_id, title, priority) VALUES($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, todo.ActivityGroupID, todo.Title, todo.Priority).Scan(&todo.ID)
	helper.PanicIfError(err)

	return todo, nil
}

func (repository *TodoRepository) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error) {
	query := "UPDATE todos SET title = $1, updated_at = Now() WHERE id = $2 RETURNING *"
	err := tx.QueryRowContext(ctx, query, todo.Title, todo.ID).Scan(&todo.ID, &todo.ActivityGroupID, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
	helper.PanicIfError(err)

	return todo, nil
}

func (repository *TodoRepository) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	query := "DELETE FROM todos WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, todo.ID)
	helper.PanicIfError(err)
}
