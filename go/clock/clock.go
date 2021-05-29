package clock

import "fmt"

const (
	minsInAnHour = 60
	hoursInADay  = 24
)

//Clock represents time without date
type Clock struct {
	hours int
	mins  int
}

//String formats the time of the Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.mins)
}

//Add adds given mins to the Clock
func (c Clock) Add(mins int) Clock {
	return New(c.hours, c.mins+mins)
}

//Subtract deducts given mins from the Clock
func (c Clock) Subtract(mins int) Clock {
	return New(c.hours, c.mins-mins)
}

//New creates a new Clock without date
func New(hours int, mins int) Clock {
	if mins < 0 {
		hours += (mins / minsInAnHour) - 1
		mins = minsInAnHour + (mins % minsInAnHour)
	}
	if hours < 0 {
		hours = hoursInADay + (hours % hoursInADay)
	}
	if mins >= minsInAnHour {
		hours += mins / minsInAnHour
		mins %= minsInAnHour
	}
	if hours >= hoursInADay {
		hours %= hoursInADay
	}
	return Clock{hours: hours, mins: mins}
}
