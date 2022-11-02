package employees

type Employee struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty`
	ID          int64  `json:"id,omitempty"`
	DateCreated string `json:"dateCreated"`
}

type Employees []Employee

type TimeCard struct {
	ID        int64  `json:"id,omitempty"`
	ClockIn   string `json:"clockIn,omitempty"`
	ClockOut  string `json:"clockOut,omitempty"`
	TotalTime string `json:"totalTime,omitempty"`
}
