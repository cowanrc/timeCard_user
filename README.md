# TimeCard

## Start:

To start the time card application, run this command in your terminal in the directory

`
go run .
`

This starts the application locally and allows access to the Swagger UI at :8080/explorer/

## REST API Calls
1. GET /employees returns all employees with name, employeeID, and DoB. If the employee has clocked in and/or clocked out, the data will be returned as well
2. POST /employees takes a payload that contains the employees name and DoB and generates an employeeID for that employee
3. GET /employees/{id} returns all information associated to that specific employee. If the employee has clocked in and clocked out, the response payload returns totalTime
4. POST /employees/ClockIn/{id} takes the employeeID as the parameter and returns the time the employee clocked in UTC
5. POST /employees/ClockOut/{id} takes the employeeID and returns the clock in and clock out time
6. DELETE /employees/{id} deletes the employee's information

## Tests

To run tests, run this command in your terminal in the directory

`
go test -v
`