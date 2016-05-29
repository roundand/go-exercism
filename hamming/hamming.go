// hamming provides a function which calculates the Hamming difference between two DNA strands, represented by equal length strings
package hamming

import (
	"fmt"
)

const testVersion = 4

// returns Hamming distance between two strings as an integer
func Distance(a, b string) (int, error) {

	// if strings are empty or of different lengths, return error
	if (len(a) != len(b)) {
		return -1, fmt.Errorf("Expected arguments of equal length but got len(a) '%v' and len(b) '%v'", len(a), len(b))
	}

	dist := 0

	// for bonus points, we'll cast both strings to rune arrays
	for i, e := range ([]rune(a)) {
		if (e != ([]rune(b))[i]) {
			dist++
		}
	}

	// go there
	return dist, nil
}
