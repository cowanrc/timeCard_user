package main

import (
	"log"
	"time"
)

func employeeClockIn(id int) {
	log.Printf("ClockIN function being called")
	employee := TimeCard[id]
	log.Printf("Employee exists: %s", employee.EmployeeName)
	employee.ClockIn = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	log.Printf("Employee clocked in at: %s", employee.ClockIn)
	return
}

func employeeClockOut(id int) {
	employee := TimeCard[id]
	employee.ClockOut = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	return
}

func employeeExists(ID int) bool {
	_, ok := TimeCard[ID]
	return ok
}

// func employeeTotalTime(name string) {
// 	currentEmployee := TimeCard[name]
// 	totalTime := currentEmployee.ClockOut.Sub(currentEmployee.ClockIn)
// 	currentEmployee.TotalTime = totalTime.Round(time.Second).String()
// }
