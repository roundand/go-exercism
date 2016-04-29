// clock provides a type which supports minute arithmetic.
package clock

import (
	"fmt"
)

const testVersion = 4

// basic type which supports hours, minutes and arithmetic.
type Clock struct {
	h int
	m int
}

// creates a new Clock instance.
func New(hour int, minute int) Clock {

	var h, m int

	// work out absolute position on timeline in minutes
	a := (hour * 60) + minute
	if a < 0 {
		// convert a negative point on timeline to hours and minutes
		// calculate minute
		m = (60 + (a % 60)) % 60

		// do hours the obvious way ...
		h = 24 + ((a / 60) % 24)

		// ... unless there were some loose minutes
		if m > 0 {
			h = h - 1
		}
	} else {
		h = (a / 60) % 24
		m = a % 60
	}
	c := Clock{h: h, m: m}
	return c
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.h, clock.m)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.h, (clock.m + minutes))
}
