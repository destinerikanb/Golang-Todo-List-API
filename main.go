package main

import (
	"github.com/go-playground/validator/v10"
	"golang-todo-list/app"
	"golang-todo-list/controller/activity_controller"
	controller2 "golang-todo-list/controller/todo_controller"
	"golang-todo-list/helper"
	"golang-todo-list/repository"
	"golang-todo-list/service"
	"net/http"
)

func main() {
	validate := validator.New()
	DB := app.NewDB()
	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityServices(activityRepository, DB, validate)
	activityController := controller.NewActivityController(activityService)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, DB, validate)
	todoController := controller2.NewTodoController(todoService)

	router := app.NewRouter(activityController, todoController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
