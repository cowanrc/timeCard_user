package employees

import (
	"timeCard/datasources/mysql/employees_db"
	"timeCard/logger"
	errors "timeCard/utils"
)

const (
	queryInsertEmployee = "INSERT INTO employees(firstName, lastName, DoB) VALUES(?, ?, ?);"
)

func (employee *Employee) Save() *errors.RestErr {
	stmt, err := employees_db.Client.Prepare(queryInsertEmployee)
	if err != nil {
		logger.Error("error when trying to prepare save employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(employee.FirstName, employee.LastName, employee.DateOfBirth)
	if saveErr != nil {
		logger.Error("error when trying to save employee", saveErr)
		return errors.NewInternalServerError("database error")
	}

	employeeId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert ID after creating New Employee", err)
		return errors.NewInternalServerError("database error")
	}

	employee.ID = employeeId
	return nil

}
