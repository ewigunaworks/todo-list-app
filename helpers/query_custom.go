package helpers

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

var (
	ListTodoFilter = map[string]string{"CreatedAt": "created_at", "Title": "title", "DueDate": "due_date"}
	ListTodoSearch = map[string]string{"CreatedAt": "created_at", "Status": "status", "Title": "title", "Description": "description", "DueDate": "due_date"}
)

func ConstructFilterQueryList(in *RequestList, filterMap map[string]string) ([]string, []interface{}, error) {
	filterFields := make([]string, 0)
	filterValues := make([]interface{}, 0)

	filterField := in.FilterField
	filterValue := in.FilterValue

	if in.IsFilterFieldEmpty() && !in.IsFilterValueEmpty() {
		return filterFields, filterValues, errors.New("filter field cannot be empty")
	}

	if !in.IsFilterFieldEmpty() && in.IsFilterValueEmpty() {
		return filterFields, filterValues, errors.New("filter value cannot be empty")
	}

	if filterField != "" && filterValue != "" {
		if filterField == "Status" {
			if strings.Contains(strings.ToLower(filterValue), "waiting") {
				filterValue = "0"
			}

			if strings.Contains(strings.ToLower(filterValue), "progress") {
				filterValue = "1"
			}

			if strings.Contains(strings.ToLower(filterValue), "done") {
				filterValue = "2"
			}
		}

		if _, ok := filterMap[filterField]; ok {
			where := filterMap[filterField]
			filterFields = append(filterFields, fmt.Sprint(where+" = ?"))
			filterValues = append(filterValues, filterValue)
		} else {
			return filterFields, filterValues, errors.New("filter field invalid")
		}
	}

	return filterFields, filterValues, nil
}

func ConstructSearchQueryList(in *RequestList, searchMap map[string]string) ([]string, []interface{}) {
	searchFields := make([]string, 0)
	searchValues := make([]interface{}, 0)

	if in.Search != "" {
		for _, searchField := range searchMap {
			searchFields = append(searchFields, fmt.Sprint(searchField+" LIKE ?"))
			searchValues = append(searchValues, fmt.Sprint("%"+in.Search+"%"))
		}
	}

	return searchFields, searchValues
}

func ConstructWhere(db *gorm.DB, fields []string, values []interface{}) {
	db.Where(strings.Join(fields, " AND "), values...)
}

func ConstructWhereLike(db *gorm.DB, fields []string, values []interface{}) {
	// var conditions []string
	if len(fields) > 0 {
		for i := 0; i < len(fields); i++ {
			// conditions = append(conditions, fields[i] + " LIKE '%"+values...
		}
	}
	db.Where(strings.Join(fields, " AND "), values...)
}

func ConstructOr(db *gorm.DB, fields []string, values []interface{}) {
	db.Where(strings.Join(fields, " OR "), values...)
}
