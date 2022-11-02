package services

import (
	"timeCard/domain/employees"
	"timeCard/utils/date_utils"
	"timeCard/utils/errors"
)

var (
	EmployeeService employeesServiceInterface = &employeesService{}
)

type employeesService struct{}

type employeesServiceInterface interface {
	CreateEmployee(employees.Employee) (*employees.Employee, *errors.RestErr)
	GetAllEmployees() (employees.Employees, *errors.RestErr)
	GetEmployee(int64) (*employees.Employee, *errors.RestErr)
}

func (s *employeesService) CreateEmployee(employee employees.Employee) (*employees.Employee, *errors.RestErr) {

	employee.DateCreated = date_utils.GetNowDBFormat()

	if err := employee.Save(); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *employeesService) GetAllEmployees() (employees.Employees, *errors.RestErr) {
	dao := &employees.Employee{}
	return dao.GetAll()
}

func (s *employeesService) GetEmployee(employeeId int64) (*employees.Employee, *errors.RestErr) {
	result := &employees.Employee{ID: employeeId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
