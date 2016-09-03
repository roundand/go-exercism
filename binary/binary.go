// Package binary provides a function to generate the numeric value of a binary.
package binary

import (
	"fmt"
)

const testVersion = 1

// ParseBinary parses a string of '1' and '0' values to generate the int result, or an error.
func ParseBinary(binary string) (int, error) {

	var result = 0
	l := len(binary) - 1
	for i, c := range binary {
		switch c {
		case '1':
			result += (1 << uint(l - i))
		case '0':
			continue
		default:
			return 0, fmt.Errorf("Non-binary character '%s' encountered.", string(c))
		}
	}

	return result, nil
}
