package routes

import (
	"os"
	"todo-list-app/controllers"

	"github.com/labstack/echo/v4"
)

func Router() {
	routing := echo.New()

	routeV1 := routing.Group("/api/v1/todo")

	routeV1.POST("/create", controllers.CreateTodo)
	routeV1.PUT("/update/:id", controllers.UpdateTodo)
	routeV1.DELETE("/delete/:id", controllers.DeleteTodo)
	routeV1.GET("/detail/:id", controllers.DetailTodo)
	routeV1.GET("/list", controllers.ListTodo)

	routing.Logger.Fatal(routing.Start(":" + os.Getenv("APP_PORT")))
}
