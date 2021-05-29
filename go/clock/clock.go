package clock

import "fmt"

const (
	minsInAnHour = 60
	hoursInADay  = 24
)

type clock struct {
	hours int
	mins  int
}

//String formats the time of the clock
func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.mins)
}

//Add adds given mins to the clock
func (c clock) Add(mins int) clock {
	return New(c.hours, c.mins+mins)
}

//Subtract deducts given mins from the clock
func (c clock) Subtract(mins int) clock {
	return New(c.hours, c.mins-mins)
}

//New creates a new clock without date
func New(hours int, mins int) clock {
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
	return clock{hours: hours, mins: mins}
}
