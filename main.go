package main

import (
	"todo-list-app/databases"
	"todo-list-app/models"
	"todo-list-app/routes"
)

func main() {
	// init DB
	db := databases.InitDB()
	db.AutoMigrate(&models.Task{})

	// routing
	routes.Router()
}
