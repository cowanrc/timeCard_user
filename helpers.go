package main

import (
	"log"
	"time"
)

func employeeClockIn(id int) {
	log.Printf("ClockIn function being called")
	employee := TimeCard[id]
	log.Printf("Employee exists: %s", employee.Name)
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

func employeeTotalTime(id int) {
	employee := TimeCard[id]
	clockIn, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", employee.ClockIn)
	clockOut, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", employee.ClockOut)
	totalTime := clockOut.Sub(clockIn)
	employee.TotalTime = totalTime.Round(time.Second)

	log.Printf("Employee: %s worked for a total of: %v", employee.Name, employee.TotalTime)
}
