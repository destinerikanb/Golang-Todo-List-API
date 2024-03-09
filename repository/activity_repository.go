package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-todo-list/helper"
	"golang-todo-list/model/domain"
	"strconv"
)

type ActivityRepository struct {
}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{}
}

func (repository *ActivityRepository) GetAllActivity(ctx context.Context, tx *sql.Tx) ([]domain.Activity, error) {
	query := "SELECT * FROM activities order by id"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var activities []domain.Activity
	for rows.Next() {
		actvity := domain.Activity{}

		err := rows.Scan(&actvity.ID, &actvity.Email, &actvity.Title, &actvity.CreatedAt, &actvity.UpdatedAt, &actvity.DeletedAt)
		helper.PanicIfError(err)

		activities = append(activities, actvity)
	}

	return activities, nil
}

func (repository *ActivityRepository) GetOneActivity(ctx context.Context, tx *sql.Tx, ID int) (domain.Activity, error) {
	query := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, ID)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var activity domain.Activity
	if rows.Next() {
		//variable to handle null value in time data type (non nil pointer)
		var (
			deletedAt sql.NullTime
		)

		err := rows.Scan(&activity.ID, &activity.Email, &activity.Title, &activity.CreatedAt, &activity.UpdatedAt, &deletedAt)
		helper.PanicIfError(err)

		//assign value
		if deletedAt.Valid {
			activity.DeletedAt = &deletedAt.Time
		} else {
			activity.DeletedAt = nil
		}

	} else {
		errStr := "Activity with ID " + strconv.Itoa(ID) + " Not Found"
		return activity, errors.New(errStr)
	}

	return activity, nil
}

func (repository *ActivityRepository) Create(ctx context.Context, tx *sql.Tx, activity domain.Activity) (domain.Activity, error) {
	query := "INSERT INTO activities(email, title) VALUES($1, $2) RETURNING id, created_at, updated_at, deleted_at"
	err := tx.QueryRowContext(ctx, query, activity.Email, activity.Title).Scan(&activity.ID, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
	if err != nil {
		panic(err)
	}
	return activity, nil
}

func (repository *ActivityRepository) Update(ctx context.Context, tx *sql.Tx, activity domain.Activity) (domain.Activity, error) {
	query := "UPDATE activities SET title = $1, updated_at = Now() WHERE id = $2 RETURNING *"
	err := tx.QueryRowContext(ctx, query, activity.Title, activity.ID).Scan(&activity.ID, &activity.Email, &activity.Title, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
	if err != nil {
		panic(err)
	}
	return activity, nil
}

func (repository *ActivityRepository) Delete(ctx context.Context, tx *sql.Tx, activity domain.Activity) {
	query := "DELETE FROM activities WHERE id = $1"

	_, err := tx.ExecContext(ctx, query, activity.ID)
	helper.PanicIfError(err)
}
