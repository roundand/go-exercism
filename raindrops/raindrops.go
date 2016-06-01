// raindrops converts an integer to raindrop-speak
package raindrops

import (
	"fmt"
)

const testVersion = 2

// do expensive assignment outside function for performance reasons
// surprisingly (for mr, anyway), var is consistently faster than const
var pling, plang, plong = "Pling", "Plang", "Plong"

func Convert(n int) (drops string) {
	// this takes slightly longer
	// var pling, plang, plong = "Pling", "Plang", "Plong"

	if n%3 == 0 {
		drops = pling
	}
	if n%5 == 0 {
		drops += plang
	}
	if n%7 == 0 {
		drops += plong
	}
	if drops == "" {
		return fmt.Sprintf("%v", n)
	}
	return drops
}
