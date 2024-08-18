package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Id          int64      `gorm:"primary_key" json:"id"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Status      *int32     `gorm:"column:status"`
	DueDate     *time.Time `gorm:"column:due_date;type:date"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

type RequestTask struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      *int32 `json:"status"`
	DueDate     string `json:"dueDate"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

type ResponseTask struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"dueDate"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

func ConvertStatus(status int32) string {
	switch status {
	case 0:
		return "Waiting List"
	case 1:
		return "In Progress"
	default:
		return "Done"
	}
}
