package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-todo-list/exception"
	"golang-todo-list/helper"
	"golang-todo-list/model/web"
	"golang-todo-list/service"
	"log"
	"net/http"
	"strconv"
)

type TodoController struct {
	TodoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{TodoService: todoService}
}

func (todoController *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todos := todoController.TodoService.GetAllTodo(r.Context())

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (todoController *TodoController) GetTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	data := todoController.TodoService.GetTodoById(r.Context(), id)

	var responseData interface{}
	if data.ID == 0 {
		responseData = exception.EmptyData{}
	} else {
		responseData = data
	}

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    responseData,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (todoController *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//Read json
	todoCreateRequest := web.TodoCreateRequest{}
	helper.ReadFromRequestBody(r, &todoCreateRequest)
	log.Println(todoCreateRequest)

	data := todoController.TodoService.SaveTodo(r.Context(), todoCreateRequest)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(w, webResponse)
}

func (todoController *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoUpdateRequest := web.TodoUpdateRequest{}
	helper.ReadFromRequestBody(r, &todoUpdateRequest)
	todoUpdateRequest.ID = id
	helper.PanicIfError(err)

	data := todoController.TodoService.UpdateTodo(r.Context(), todoUpdateRequest)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (todoController *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Masuk controller")
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoController.TodoService.DeleteTodo(r.Context(), id)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    exception.EmptyData{},
	}

	helper.WriteToResponseBody(w, webResponse)

}
