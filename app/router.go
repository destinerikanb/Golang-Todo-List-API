package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-todo-list/controller/activity_controller"
	controller2 "golang-todo-list/controller/todo_controller"
	"golang-todo-list/exception"
)

func NewRouter(activityController *controller.ActivityController, todoController *controller2.TodoController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/activities", activityController.GetAllActivities)
	router.GET("/api/v1/activities/:activityId", activityController.GetActivityById)
	router.POST("/api/v1/activities", activityController.CreateActivity)
	router.PATCH("/api/v1/activities/:activityId", activityController.UpdateActivity)
	router.DELETE("/api/v1/activities/:activityId", activityController.DeleteActivity)

	router.GET("/api/v1/todos", todoController.GetAllTodos)
	router.GET("/api/v1/todos/:todoId", todoController.GetTodoById)
	router.POST("/api/v1/todos", todoController.CreateTodo)
	router.PUT("/api/v1/todos/:todoId", todoController.UpdateTodo)
	router.DELETE("/api/v1/todos/:todoId", todoController.DeleteTodo)

	router.PanicHandler = exception.ErrorHandler

	return router
}
