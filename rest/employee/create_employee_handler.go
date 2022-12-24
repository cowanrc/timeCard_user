package employee

import (
	"net/http"

	"timeCard_user/domain/employees"
	"timeCard_user/services"
	"timeCard_user/utils/errors"

	"github.com/gin-gonic/gin"
)

// CreateEmployeeHandler to enter name and DOB and get an employee ID in return
func HandleCreateEmployee(c *gin.Context) {
	var employee employees.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.EmployeeService.CreateEmployee(employee)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
