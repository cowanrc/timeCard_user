package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateEmployee struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dob"`
	ID          int    `json:"id"`
}

type Employee struct {
	EmployeeName string    `json:"employeeName,omitempty"`
	EmployeeID   int       `json:"employeeID,omitempty"`
	ClockIn      string    `json:"clockIn,omitempty"`
	ClockOut     string    `json:"clockOut,omitempty"`
	TotalTime    time.Time `json:"totalTime,omitempty"`
}

var TimeCard = make(map[int]*Employee)
var seq = 1

var employee Employee

//CreateEmployeeHandler to enter name and DOB and get an employee ID in return
func CreateEmployeeHandler(ctx echo.Context) error {
	var newEmployee Employee
	newEmployee.EmployeeID = seq

	err := ctx.Bind(&newEmployee)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error binding the structure")
	}

	log.Printf("Employee Name is : %s", employee.EmployeeName)
	log.Printf("Employee id is: %v", employee.EmployeeID)
	TimeCard[newEmployee.EmployeeID] = &employee

	for k, v := range TimeCard {
		log.Printf("Values and Keys %v %v", k, v)
	}
	seq++

	m := map[string]string{
		"name":       newEmployee.EmployeeName,
		"employeeID": strconv.Itoa(newEmployee.EmployeeID),
	}

	return ctx.JSON(http.StatusAccepted, m)
}

func ClockInHandler(ctx echo.Context) error {
	id := employee.EmployeeID
	log.Printf("Employee ID is: %v", id)
	for k, v := range TimeCard {
		log.Printf("Values and Keys %v %v", k, v)
	}

	err := ctx.Bind(&employee)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error binding the structure")
	}

	if !employeeExists(id) {
		return echo.NewHTTPError(http.StatusBadRequest, "Employee ID doesn't exist so you cannot clock in")
	}

	employeeClockIn(id)

	m := map[string]string{
		"name":       employee.EmployeeName,
		"employeeID": strconv.Itoa(employee.EmployeeID),
		"clockIn":    employee.ClockIn,
	}

	// if employee.ClockIn != "" && employee.ClockOut == "" {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Cannot Clock in again without clocking out first")
	// }

	return ctx.JSON(http.StatusAccepted, m)

}
