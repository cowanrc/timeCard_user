package employee

import (
	"net/http"
	"timeCard/services"

	"github.com/gin-gonic/gin"
)

func HandleDeleteEmployee(c *gin.Context) {
	employeeId, idErr := getEmployeeId(c.Param("employeeId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.EmployeeService.DeleteEmployee(employeeId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
