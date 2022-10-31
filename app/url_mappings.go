package app

import (
	"timeCard/rest/employee"
)

func mapUrls() {
	r.POST("/employees", employee.HandleCreateEmployee)
}
