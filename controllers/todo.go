package controllers

import (
	"net/http"
	"strconv"
	"todo-list-app/helpers"
	"todo-list-app/models"
	"todo-list-app/service"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	request := new(models.RequestTask)
	if err := c.Bind(request); err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	response, err := service.CreateTodo(c, request)
	if err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if response.Code != http.StatusOK {
		return helpers.HandleResponse(c, response.Code, response.Message.(string), nil)
	}

	return helpers.HandleResponse(c, response.Code, response.Message.(string), response.Data)
}

func UpdateTodo(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "invalid parameter", nil)
	}

	request := new(models.RequestTask)
	request.Id = int64(id)
	if err := c.Bind(request); err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	response, err := service.UpdateTodo(c, request)
	if err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if response.Code != http.StatusOK {
		return helpers.HandleResponse(c, response.Code, response.Message.(string), nil)
	}

	return helpers.HandleResponse(c, response.Code, response.Message.(string), response.Data)

}

func DeleteTodo(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "invalid parameter", nil)
	}

	request := new(models.RequestTask)
	request.Id = int64(id)
	if err := c.Bind(request); err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	response, err := service.DeleteTodo(c, request)
	if err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if response.Code != http.StatusOK {
		return helpers.HandleResponse(c, response.Code, response.Message.(string), nil)
	}

	return helpers.HandleResponse(c, response.Code, response.Message.(string), response.Data)

}

func DetailTodo(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "invalid parameter", nil)
	}

	request := new(models.RequestTask)
	request.Id = int64(id)
	if err := c.Bind(request); err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	response, err := service.DetailTodo(c, request)
	if err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if response.Code != http.StatusOK {
		return helpers.HandleResponse(c, response.Code, response.Message.(string), nil)
	}

	return helpers.HandleResponse(c, response.Code, response.Message.(string), response.Data)

}

func ListTodo(c echo.Context) error {
	page := c.QueryParam("Page")
	limit := c.QueryParam("Limit")
	sort := c.QueryParam("Sort")
	sortBy := c.QueryParam("SortBy")
	filterField := c.QueryParam("FilterField")
	filterValue := c.QueryParam("FilterValue")
	keyword := c.QueryParam("Search")
	newPage, _ := strconv.Atoi(page)
	newLimit, _ := strconv.Atoi(limit)

	request := &helpers.RequestList{
		Page:        int32(newPage),
		Limit:       int32(newLimit),
		Sort:        sort,
		SortBy:      sortBy,
		FilterField: filterField,
		FilterValue: filterValue,
		Search:      keyword,
	}

	response, err := service.ListTodo(c, request)
	if err != nil {
		return helpers.HandleResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if response.Code != http.StatusOK {
		return helpers.ListResponses(c, response)
	}

	return helpers.ListResponses(c, response)
}
