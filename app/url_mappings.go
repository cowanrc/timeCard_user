package app

import (
	"timeCard_user/rest/employee"
)

func mapUrls() {
	r.POST("/employees", employee.HandleCreateEmployee)
	r.GET("/employees", employee.HandleListEmployees)
	r.GET("/employees/:employeeId", employee.HandleListEmployee)
	r.DELETE("/employees/:employeeId", employee.HandleDeleteEmployee)
}
