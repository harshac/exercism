package clock

import "fmt"

const (
	minsInAnHour = 60
	hoursInADay  = 24
)

//Clock represents minutes
type Clock struct {
	mins int
}

//String formats the time of the Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.mins/minsInAnHour, c.mins%minsInAnHour)
}

//Add adds given mins to the Clock
func (c Clock) Add(mins int) Clock {
	return New(0, c.mins+mins)
}

//Subtract deducts given mins from the Clock
func (c Clock) Subtract(mins int) Clock {
	return New(0, c.mins-mins)
}

//New creates a new Clock without date
func New(hours int, mins int) Clock {
	m := mins + hours*minsInAnHour
	m %= hoursInADay * minsInAnHour
	if m < 0 {
		m += hoursInADay * minsInAnHour
	}
	return Clock{mins: m}
}
