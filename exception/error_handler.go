package exception

import (
	"github.com/go-playground/validator/v10"
	"golang-todo-list/helper"
	"golang-todo-list/model/web"
	"net/http"
)

type EmptyData struct{}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if validationError(w, r, err) {
		return
	}
	internalServerError(w, r, err)

}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Status:  "Bad request",
			Message: "Title cannot be null",
			Data:    EmptyData{},
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	//konversi error ke dalam struck NotFoundError
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Status:  "Not Found",
			Message: exception.Error,
			Data:    EmptyData{},
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status:  "Internal Server Error",
		Message: err,
		Data:    EmptyData{},
	}

	helper.WriteToResponseBody(w, webResponse)

}
