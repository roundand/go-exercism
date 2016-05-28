// gigasecond provides a function which adds one gs to a date.
package gigasecond

// we'll need the time library
import (
	"time"
)

// Constant declaration.
const testVersion = 4

// API function to add one gigasecond to a time value using standard go type
func AddGigasecond(t time.Time) time.Time {
	gs, _ := time.ParseDuration("1000000000s")
	return t.Add(gs)
}
