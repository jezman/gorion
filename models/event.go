package models

import (
	"time"
)

type Event struct {
	Employee   Employee
	FirstTime  time.Time
	LastTime   time.Time
	Company    Company
	Door       Door
	Event      string
	WorkedTime time.Duration
}
