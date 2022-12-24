package services

import (
	"fmt"
	"timeCard_user/domain/employees"
	"timeCard_user/utils/date_utils"
	"timeCard_user/utils/errors"
)

var (
	EmployeeService employeesServiceInterface = &employeesService{}
)

type employeesService struct{}

type employeesServiceInterface interface {
	CreateEmployee(employees.Employee) (*employees.Employee, *errors.RestErr)
	GetAllEmployees() (employees.Employees, *errors.RestErr)
	GetEmployee(int64) (*employees.Employee, *errors.RestErr)
	DeleteEmployee(int64) *errors.RestErr
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

func (s *employeesService) DeleteEmployee(employeeId int64) *errors.RestErr {
	employee := &employees.Employee{ID: employeeId}
	_, err := EmployeeService.GetEmployee(employee.ID)

	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("Employee ID %d does not exist", employeeId))
	}

	return employee.Delete()
}
