package main

import "timeCard_user/app"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"

func main() {
	app.StartApplication()
}

// 	log.Printf("Creating your time card application")
// 	e := echo.New()
// e.Pre(middleware.RemoveTrailingSlash())

// e.File("/explorer", "ui/index.html")
// e.Static("/explorer", "ui")

// 	//routes
// 	e.POST("/employees",employeeHandler.)
// 	e.GET("/employees", GetAllEmployeeHandler)
// 	e.GET("/employees/:id", GetEmployeeHandler)
// 	e.POST("/employees/ClockIn/:id", ClockInHandler)
// 	e.POST("/employees/ClockOut/:id", ClockOutHandler)
// 	e.DELETE("/employees/:id", DeleteEmployeeHandler)

// 	log.Printf("listening on port 8080")
// 	e.Logger.Fatal((e.Start(":8080")))

// }
