package employee

import (
	"net/http"
	"strconv"
	"timeCard/services"
	"timeCard/utils/errors"

	"github.com/gin-gonic/gin"
)

func getEmployeeId(employeeIdParam string) (int64, *errors.RestErr) {
	employeeId, empErr := strconv.ParseInt(employeeIdParam, 10, 64)

	if empErr != nil {
		return 0, errors.NewBadRequestError("user id should be an integer")
	}

	return employeeId, nil
}

func HandleListEmployees(c *gin.Context) {
	employees, getErr := services.EmployeeService.GetAllEmployees()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func HandleListEmployee(c *gin.Context) {
	employeeId, idErr := getEmployeeId(c.Param("employeeId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	employee, getErr := services.EmployeeService.GetEmployee(employeeId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, employee)

}
