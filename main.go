package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.File("/explorer", "ui/index.html")
	e.Static("/explorer", "ui")

	//routes
	e.POST("/employee", CreateEmployeeHandler)
	e.POST("/employee/ClockIn", ClockInHandler)

	log.Printf("listening on port 8080")
	e.Logger.Fatal((e.Start(":8080")))
}
