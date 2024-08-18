package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-list-app/databases"
	"todo-list-app/helpers"
	"todo-list-app/models"

	"github.com/labstack/echo/v4"
)

func CreateTodo(ctx echo.Context, request *models.RequestTask) (*helpers.Responses, error) {
	var dueDate, createdAt time.Time
	layout := "2006-01-02"
	if request.DueDate != "" {
		date, err := time.ParseInLocation(layout, request.DueDate, time.Local)
		if err != nil {
			fmt.Println("Error parsing date", err)
			return nil, errors.New("error parsing date")
		}
		dueDate = date
	}

	createdAt = time.Now().Local()

	// insert task
	requestTask := &models.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		DueDate:     &dueDate,
		CreatedAt:   &createdAt,
	}

	query := databases.DB.Table("tasks").Omit("updated_at").Create(&requestTask)

	if query.Error != nil {
		fmt.Println("Error Create", query.Error)
		return &helpers.Responses{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: query.Error,
		}, nil
	}

	// get detail
	getDetail, _ := getDetailTask(int64(requestTask.Id))

	return &helpers.Responses{
		Code:    http.StatusOK,
		Message: "Success Create Todo List",
		Data:    getDetail,
	}, nil
}

func UpdateTodo(ctx echo.Context, request *models.RequestTask) (*helpers.Responses, error) {
	id := request.Id
	layout := "2006-01-02"
	var dueDate, updatedAt time.Time

	_, err := getDetailTask(id)
	if err != nil {
		return &helpers.Responses{
			Code:    http.StatusNotFound,
			Message: "Todo Not Found",
			Data:    nil,
		}, nil
	}

	if request.DueDate != "" {
		date, err := time.ParseInLocation(layout, request.DueDate, time.Local)
		if err != nil {
			fmt.Println("Error parsing date", err)
			return nil, errors.New("error parsing date")
		}
		dueDate = date.Local()
	}

	updatedAt = time.Now().Local()

	requestTask := &models.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		DueDate:     &dueDate,
		UpdatedAt:   &updatedAt,
	}

	query := databases.DB.Table("tasks").Where("id = ?", id).Updates(&requestTask)
	if query.Error != nil {
		fmt.Println("Error Update:", err.Error())
		return &helpers.Responses{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Todo List",
			Data:    nil,
		}, nil
	}

	// get detail
	getDetail, _ := getDetailTask(id)

	return &helpers.Responses{
		Code:    http.StatusOK,
		Message: "Success Update Todo List",
		Data:    getDetail,
	}, nil
}

func DeleteTodo(ctx echo.Context, request *models.RequestTask) (*helpers.Responses, error) {
	id := request.Id
	var deletedAt time.Time

	_, err := getDetailTask(id)
	if err != nil {
		return &helpers.Responses{
			Code:    http.StatusNotFound,
			Message: "Todo Not Found",
			Data:    nil,
		}, nil
	}

	deletedAt = time.Now().Local()

	requestTask := &models.Task{
		Id:        id,
		DeletedAt: &deletedAt,
	}

	query := databases.DB.Table("tasks").Omit("created_at, updated_at").Where("id = ?", id).Updates(&requestTask)
	if query.Error != nil {
		fmt.Println("Error delete todo:", query.Error)
		return &helpers.Responses{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Todo List",
			Data:    nil,
		}, nil
	}

	return &helpers.Responses{
		Code:    http.StatusOK,
		Message: "Success Delete Todo List",
		Data:    nil,
	}, nil
}

func DetailTodo(ctx echo.Context, request *models.RequestTask) (*helpers.Responses, error) {
	id := request.Id

	getDetail, err := getDetailTask(id)
	if err != nil {
		return &helpers.Responses{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Failed Get Detail Todo List %v", err.Error()),
			Data:    nil,
		}, nil
	}

	return &helpers.Responses{
		Code:    http.StatusOK,
		Message: "Success Get Detail Todo List",
		Data:    getDetail,
	}, nil
}

func ListTodo(ctx echo.Context, request *helpers.RequestList) (*helpers.ResponsesList, error) {
	datas := make([]*models.Task, 0)
	sortBy := "tasks.created_at desc"

	var counter int64
	var paginate *helpers.Paginate

	query := databases.DB.Table("tasks").Where("deleted_at IS NULL")
	queryCount := databases.DB.Table("tasks").Where("deleted_at IS NULL")

	filterFields, filterValues, err := helpers.ConstructFilterQueryList(request, helpers.ListTodoFilter)
	if err != nil {
		return &helpers.ResponsesList{
			Code:     http.StatusNotFound,
			Message:  err.Error(),
			Data:     datas,
			Paginate: paginate,
		}, nil
	}

	searchFields, searchValues := helpers.ConstructSearchQueryList(request, helpers.ListTodoSearch)

	if len(filterFields) > 0 {
		helpers.ConstructWhere(query, filterFields, filterValues)
		helpers.ConstructWhere(queryCount, filterFields, filterValues)
	}

	if len(searchFields) > 0 {
		helpers.ConstructOr(query, searchFields, searchValues)
		helpers.ConstructOr(queryCount, searchFields, searchValues)
	}

	queryCount.Count(&counter)

	paginate, offset := helpers.Pagination(int(request.Page), int(request.Limit), int(counter))

	query.Offset(offset).Limit(int(paginate.PerPage)).Order(sortBy).Find(&datas)

	return &helpers.ResponsesList{
		Code:     http.StatusOK,
		Message:  "Success Get List Todo List",
		Data:     TransformTodoList(datas),
		Paginate: paginate,
	}, nil
}

func getDetailTask(id int64) (*models.ResponseTask, error) {
	var out *models.Task
	layout := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"
	query := databases.DB.Table("tasks").Where("id = ?", id).Where("deleted_at IS NULL").First(&out)

	if query.Error != nil {
		return nil, query.Error
	}

	var createdAt, updatedAt, dueDate string
	if out.CreatedAt != nil {
		createdAt = out.CreatedAt.Format(layout)
	}

	if out.UpdatedAt != nil {
		updatedAt = out.UpdatedAt.Format(layout)
	}

	if out.DueDate != nil {
		dueDate = out.DueDate.Format(layoutDate)
	}

	data := &models.ResponseTask{
		Id:          out.Id,
		Title:       out.Title,
		Description: out.Description,
		Status:      models.ConvertStatus(*out.Status),
		DueDate:     dueDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	return data, nil
}

func TransformTodoList(req []*models.Task) []*models.ResponseTask {
	layout := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"
	out := make([]*models.ResponseTask, 0)
	for _, v := range req {
		statusInt, _ := strconv.ParseInt(fmt.Sprint(*v.Status), 0, 64)

		var createdAt, updatedAt, dueDate string
		if v.CreatedAt != nil {
			createdAt = v.CreatedAt.Format(layout)
		}

		if v.UpdatedAt != nil {
			updatedAt = v.UpdatedAt.Format(layout)
		}

		if v.DueDate != nil {
			dueDate = v.DueDate.Format(layoutDate)
		}

		dt := &models.ResponseTask{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Status:      models.ConvertStatus(int32(statusInt)),
			DueDate:     dueDate,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}

		out = append(out, dt)
	}

	return out
}
