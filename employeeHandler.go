package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// type CreateEmployee struct {
// 	Name        string `json:"name"`
// 	DateOfBirth string `json:"dob"`
// 	ID          int    `json:"id"`
// }

type NewEmployee struct {
	Name        string `json:"Name,omitempty"`
	EmployeeID  int    `json:"employeeID,omitempty"`
	DateOfBirth string `json:"DoB, omitempty"`
}

type Employee struct {
	Name        string `json:"Name,omitempty"`
	EmployeeID  int    `json:"employeeID,omitempty"`
	ClockIn     string `json:"clockIn,omitempty"`
	ClockOut    string `json:"clockOut,omitempty"`
	TotalTime   string `json:"totalTime,omitempty"`
	DateOfBirth string `json:"DoB, omitempty"`
}

var TimeCard = make(map[int]*Employee)
var seq = 1

//CreateEmployeeHandler to enter name and DOB and get an employee ID in return
func CreateEmployeeHandler(ctx echo.Context) error {
	var newEmployee NewEmployee
	newEmployee.EmployeeID = seq

	err := ctx.Bind(&newEmployee)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error binding the structure")
	}

	log.Printf("Employee Name is : %s", newEmployee.Name)
	log.Printf("Employee id is: %v", newEmployee.EmployeeID)

	seq++

	m := map[string]string{
		"name":       newEmployee.Name,
		"employeeID": strconv.Itoa(newEmployee.EmployeeID),
	}

	var employee Employee

	employee.Name = newEmployee.Name
	employee.EmployeeID = newEmployee.EmployeeID
	employee.DateOfBirth = newEmployee.DateOfBirth
	TimeCard[newEmployee.EmployeeID] = &employee

	return ctx.JSON(http.StatusAccepted, m)
}

func GetEmployeeHandler(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	log.Printf("Getting timecard information for employee: %s", TimeCard[id].Name)
	return ctx.JSON(http.StatusOK, TimeCard[id])
}

func GetAllEmployeeHandler(ctx echo.Context) error {
	allEmployees := make([]Employee, 0)

	for _, e := range TimeCard {
		allEmployees = append(allEmployees, Employee{e.Name, e.EmployeeID, e.ClockIn, e.ClockOut, e.TotalTime, e.DateOfBirth})
	}

	return ctx.JSON(http.StatusOK, allEmployees)
}

func DeleteEmployeeHandler(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	log.Printf("Removing employee: %d from database.", TimeCard[id].EmployeeID)
	delete(TimeCard, id)
	return ctx.NoContent(http.StatusNoContent)
}

func ClockInHandler(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	log.Printf("Employee name is : %s", TimeCard[id].Name)

	employee := TimeCard[id]

	employeeClockIn(id)

	m := map[string]string{
		"name":       employee.Name,
		"employeeID": strconv.Itoa(id),
		"clockIn":    employee.ClockIn,
	}

	return ctx.JSON(http.StatusAccepted, m)
}

func ClockOutHandler(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	employee := TimeCard[id]

	employeeClockOut(id)

	m := map[string]string{
		"name":       employee.Name,
		"employeeID": strconv.Itoa(id),
		"clockIn":    employee.ClockIn,
		"clockOut":   employee.ClockOut,
	}

	return ctx.JSON(http.StatusAccepted, m)
}
