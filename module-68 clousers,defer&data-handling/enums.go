package main

import "fmt"

type Day int

const (
	Monday    Day = iota // 0
	Sunday               // 1
	Tuesday              // 2
	Wednesday            // 3
	Friday
	Saturday
	Thursday
)

func dayStatus(day Day) string {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		return "office is open"
	case Thursday:
		return "half day"
	case Saturday, Friday:
		return "off day"
	default:
		return "invalid day"
	}
}

type OfficeStatus string

const (
	Open    OfficeStatus = "open"
	Closed  OfficeStatus = "closed"
	HalfDay OfficeStatus = "half_day"
)

func enum_func() {
	fmt.Println(dayStatus(Friday))
}
