package helpers

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type Responses struct {
	Status  string      `json:"Status"`
	Code    int         `json:"StatusCode"`
	Message interface{} `json:"Message"`
	Data    interface{} `json:"Data"`
}

type ResponsesList struct {
	Status   string      `json:"Status"`
	Code     int         `json:"StatusCode"`
	Message  string      `json:"Message"`
	Data     interface{} `json:"Data"`
	Paginate interface{} `json:"Paginate,omitempty"`
}

func (response *ResponsesList) Success(code int, message string, data interface{}) {
	response.Code = code
	response.Status = "success"
	response.Message = message
	response.Data = data
}

func (response *ResponsesList) Error(code int, message string, data interface{}) {
	response.Code = code
	response.Status = "failed"
	response.Message = message
	response.Data = data
}

func (response *Responses) Success(code int, message string, data interface{}) {
	response.Status = "success"
	response.Code = code
	response.Message = message
	response.Data = data
}

func (response *Responses) Error(code int, message string, data interface{}) {
	response.Status = "failed"
	response.Code = code
	response.Message = message
	response.Data = data
}

func ListResponses(c echo.Context, res *ResponsesList) error {
	var responseStruct = new(ResponsesList)
	if res.Code == 200 || res.Code == 201 || res.Code == 202 {
		responseStruct.Success(res.Code, res.Message, res.Data)
	} else {
		HandleError(res.Message, res.Data)
		if res.Data == nil {
			responseStruct.Error(res.Code, res.Message, nil)
		} else if fmt.Sprintf("%v", reflect.TypeOf(res.Data).Kind()) == "ptr" {
			responseStruct.Error(res.Code, res.Message, fmt.Sprintf("%v", res.Data))
		} else {
			responseStruct.Error(res.Code, res.Message, res.Data)
		}
	}
	responseStruct.Paginate = res.Paginate

	result := strings.Split(res.Message, "desc = ")
	if len(result) > 1 {
		res.Message = result[1]
	}
	return c.JSON(res.Code, responseStruct)
}

func HandleResponse(ctx echo.Context, code int, message string, data interface{}) (err error) {
	var responseStruct = new(Responses)

	if code == 200 || code == 201 || code == 202 {
		responseStruct.Success(code, message, data)
	} else {
		HandleError(message, data)
		if data == nil {
			responseStruct.Error(code, message, nil)
		} else if fmt.Sprintf("%v", reflect.TypeOf(data).Kind()) == "ptr" {
			responseStruct.Error(code, message, fmt.Sprintf("%v", data))
		} else {
			responseStruct.Error(code, message, data)
		}
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.Response().WriteHeader(code)
	return ctx.JSON(code, responseStruct)
}

func HandleError(message string, err interface{}) {
	fmt.Println()
	log.Println("========== Start Error Message ==========")
	log.Println("Message => " + message + ".")

	if err != nil {
		log.Println("Error => ", err)
	}

	log.Println("========== End Of Error Message ==========")
	fmt.Println()
}
