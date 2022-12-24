package employees

import (
	"timeCard_user/datasources/mysql/employees_db"
	"timeCard_user/logger"
	"timeCard_user/utils/errors"
)

const (
	//Should probably have employees, then ClockIn/ClockOut as seperate Structs
	queryInsertEmployee = "INSERT INTO employees(firstName, lastName, dateCreated) VALUES(?, ?, ?);"
	queryGetEmployees   = "SELECT * FROM employees;"
	queryGetEmployee    = "SELECT id, firstName, lastName, dateCreated FROM employees WHERE id=?;"
	queryDeleteEmployee = "DELETE FROM employees WHERE id=?;"
)

func (employee *Employee) Save() *errors.RestErr {
	stmt, err := employees_db.Client.Prepare(queryInsertEmployee)
	if err != nil {
		logger.Error("error when trying to prepare save employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(employee.FirstName, employee.LastName, employee.DateCreated)
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

func (employees *Employee) GetAll() ([]Employee, *errors.RestErr) {
	stmt, err := employees_db.Client.Prepare(queryGetEmployees)
	if err != nil {
		logger.Error("error when trying to prepare get employees statement", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to search rows for employees", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer res.Close()

	results := make([]Employee, 0)
	for res.Next() {
		var employee Employee
		if err := res.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.DateCreated); err != nil {
			logger.Error("error when trying to scan employee row in employee struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, employee)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No employees exist")
	}

	return results, nil
}

func (employee *Employee) Get() *errors.RestErr {
	stmt, err := employees_db.Client.Prepare(queryGetEmployee)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.ID)
	if getErr := result.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.DateCreated); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *Employee) Delete() *errors.RestErr {
	stmt, err := employees_db.Client.Prepare(queryDeleteEmployee)
	if err != nil {
		logger.Error("error when trying to prepare delete employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(employee.ID); err != nil {
		logger.Error("error when trying to get user from database", err)
		return errors.NewNotFoundError("Employee does not exist")
	}

	return nil
}
