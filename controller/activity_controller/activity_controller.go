package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang-todo-list/exception"
	"golang-todo-list/helper"
	"golang-todo-list/model/web"
	"golang-todo-list/service"
	"net/http"
	"strconv"
)

type ActivityController struct {
	ActivityService *service.ActivityService
}

func NewActivityController(activityService *service.ActivityService) *ActivityController {
	return &ActivityController{ActivityService: activityService}
}

func (activityController *ActivityController) GetAllActivities(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	activities := activityController.ActivityService.GetAllActivity(r.Context())

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Succes",
		Data:    activities,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (activityController *ActivityController) GetActivityById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	activityId := params.ByName("activityId")
	id, err := strconv.Atoi(activityId)
	helper.PanicIfError(err)

	data := activityController.ActivityService.GetActivityByID(r.Context(), id)

	var resposeData interface{}
	type EmptyData struct{}
	if data.ID == 0 {
		resposeData = EmptyData{}
	} else {
		resposeData = data
	}

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resposeData,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (activityController *ActivityController) CreateActivity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//Read json
	decoder := json.NewDecoder(r.Body)
	activityCreateRequest := web.ActivityCreateRequest{}
	err := decoder.Decode(&activityCreateRequest)
	helper.PanicIfError(err)

	data := activityController.ActivityService.SaveActivity(r.Context(), activityCreateRequest)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(w, webResponse)
}

func (activityController *ActivityController) UpdateActivity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	activityId := params.ByName("activityId")
	id, err := strconv.Atoi(activityId)
	helper.PanicIfError(err)

	activityUpdateRequest := web.ActivityUpdateRequest{}
	helper.ReadFromRequestBody(r, &activityUpdateRequest)
	activityUpdateRequest.ID = id

	data := activityController.ActivityService.Update(r.Context(), activityUpdateRequest)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (activityController *ActivityController) DeleteActivity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	activityId := params.ByName("activityId")
	id, err := strconv.Atoi(activityId)
	helper.PanicIfError(err)

	activityController.ActivityService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    exception.EmptyData{},
	}

	helper.WriteToResponseBody(w, webResponse)
}
