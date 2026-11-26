package attendance

import "time"

type Attendance struct {
	ID          int64     `json:"id"`
	EmployeeID  int64     `json:"employee_id"`
	CheckedInAt time.Time `json:"checked_in_at"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
}
