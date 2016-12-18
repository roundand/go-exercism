// Package trinary provides a function to parse trinary numbers.
package trinary

import "fmt"

const testVersion = 1

// ParseTrinary converts a trinary string to an int64 number.
func ParseTrinary(tri string) (int64, error) {
	var ret int64
	for _, r := range tri {
		if r < '0' || r > '2' {
			return 0, fmt.Errorf("bad digit: %q", r)
		}
		ret *= 3
		ret += int64(r - '0')

		// check for overflow
		if ret < 0 {
			return 0, fmt.Errorf("overflow: %d", ret)
		}
	}
	return ret, nil
}
