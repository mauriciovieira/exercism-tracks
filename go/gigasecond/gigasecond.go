// Package gigasecond implements a function to add a giga second to a given time
package gigasecond

import "time"

// AddGigasecond adds a giga second to a time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
