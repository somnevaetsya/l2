package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"l2/develop/dev11/app/models"
	"l2/develop/dev11/app/usecase"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	usecase *usecase.Usecase
}

func MakeHandler(usecase_ *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase_}
}

func (handler *Handler) CreateJSON(code int, data interface{}, resp http.ResponseWriter) {
	rawData, err := json.Marshal(data)
	if err != nil {
		handler.CreateJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
	}
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.Header().Set("Content-Length", strconv.Itoa(len(rawData)))
	resp.WriteHeader(code)
	_, _ = resp.Write(rawData)
}

func (handler *Handler) CreateEvent(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.ParseJSON(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	event, err := handler.usecase.CreateEvent(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusServiceUnavailable, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusCreated, event, resp)
	return
}

func (handler *Handler) DeleteEvent(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.ParseJSON(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	err = handler.usecase.DeleteEvent(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusServiceUnavailable, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusOK, nil, resp)
	return
}

func (handler *Handler) UpdateEvent(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.ParseJSON(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	fmt.Println("UPDATE:", inputData)
	event, err := handler.usecase.UpdateEvent(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusServiceUnavailable, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusOK, event, resp)
	return
}

func (handler *Handler) GetDayEvents(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.GetParse(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	inputData.DateTo = inputData.DateFrom
	inputData.DateTo.Add(time.Hour * 24)
	events, err := handler.usecase.GetEvents(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusOK, events, resp)
	return
}

func (handler *Handler) GetWeekEvents(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.GetParse(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	inputData.DateTo = inputData.DateFrom.Add(time.Hour * 24 * 7)
	events, err := handler.usecase.GetEvents(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusOK, events, resp)
	return
}

func (handler *Handler) GetMonthEvents(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		handler.CreateJSON(http.StatusNotFound, map[string]string{"error": errors.New("404: page not found").Error()}, resp)
		return
	}
	var inputData models.Event
	err := inputData.GetParse(req)
	if err != nil {
		handler.CreateJSON(http.StatusBadRequest, map[string]string{"error": err.Error()}, resp)
		return
	}
	inputData.DateTo = inputData.DateFrom
	inputData.DateTo.Add(time.Hour * 24 * 30)
	events, err := handler.usecase.GetEvents(inputData)
	if err != nil {
		handler.CreateJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
		return
	}
	handler.CreateJSON(http.StatusOK, events, resp)
	return
}
