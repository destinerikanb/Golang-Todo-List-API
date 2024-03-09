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
)

type ActivityService struct {
	ActivityRepository *repository.ActivityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewActivityServices(activityRepository *repository.ActivityRepository, db *sql.DB, validate *validator.Validate) *ActivityService {
	return &ActivityService{
		ActivityRepository: activityRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *ActivityService) GetAllActivity(ctx context.Context) []domain.Activity {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activities, err := service.ActivityRepository.GetAllActivity(ctx, tx)
	helper.PanicIfError(err)

	return activities
}

func (service *ActivityService) GetActivityByID(ctx context.Context, id int) domain.Activity {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, err := service.ActivityRepository.GetOneActivity(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return activity
}

func (service *ActivityService) SaveActivity(ctx context.Context, request web.ActivityCreateRequest) domain.Activity {
	//validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	//DB transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//call repository
	activity := domain.Activity{
		Email: request.Email,
		Title: request.Title,
	}

	activity, _ = service.ActivityRepository.Create(ctx, tx, activity)

	return activity
}

func (service *ActivityService) Update(ctx context.Context, request web.ActivityUpdateRequest) domain.Activity {
	//validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activity, err := service.ActivityRepository.GetOneActivity(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	activity.Title = request.Title
	activity, err = service.ActivityRepository.Update(ctx, tx, activity)

	return activity
}

func (service *ActivityService) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, err := service.ActivityRepository.GetOneActivity(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ActivityRepository.Delete(ctx, tx, activity)
}
