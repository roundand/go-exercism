// Package octal provides a function to parse octal strings.
package octal

import (
	"fmt"
)

const testVersion = 1

// ParseOctal parses an octal string and returns an integer or an error.
func ParseOctal(inp string) (int64, error) {
	len := len(inp)
	var ret int64
	for i, o := range inp {
		switch o {
		case '0', '1', '2', '3', '4', '5', '6', '7':
			// left shift bit pattern '0001' three binary places for each octal place
			pow := 1 << (uint(len-i-1) * 3)
			ret += int64(o-'0') * int64(pow)
		default:
			return 0, fmt.Errorf("Non-octal character %q found in %q\n", o, inp)
		}
	}
	return ret, nil
}
