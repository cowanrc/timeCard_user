package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Printf("Creating your time card application")
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.File("/explorer", "ui/index.html")
	e.Static("/explorer", "ui")

	//routes
	e.POST("/employees", CreateEmployeeHandler)
	e.GET("/employees/:id", GetEmployeeHandler)
	e.POST("/employees/ClockIn/:id", ClockInHandler)
	e.POST("employees/ClockOut/:id", ClockOutHandler)

	log.Printf("listening on port 8080")
	e.Logger.Fatal((e.Start(":8080")))

}
