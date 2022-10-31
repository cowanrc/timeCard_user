package services

import (
	"timeCard/domain/employees"
	errors "timeCard/utils"
)

var (
	EmployeeService employeesServiceInterface = &employeesService{}
)

type employeesService struct{}

type employeesServiceInterface interface {
	CreateEmployee(employees.Employee) (*employees.Employee, *errors.RestErr)
}

func (s *employeesService) CreateEmployee(employee employees.Employee) (*employees.Employee, *errors.RestErr) {

	if err := employee.Save(); err != nil {
		return nil, err
	}

	return &employee, nil
}
