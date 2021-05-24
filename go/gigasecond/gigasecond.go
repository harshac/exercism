package gigasecond

import "time"

const gigasecond = time.Second * 1e9

//AddGigasecond adds a gigasecond to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
