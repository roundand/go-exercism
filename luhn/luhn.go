// Package luhn provides functions to generate and validate Luhn check digits.
package luhn

import (
	//  "fmt"
	"strconv"
	"unicode"
)

// Valid checks a numeric unicode string for a correct Luhn code.
func Valid(ns string) bool {
	cs, ok := checkSum(ns)
	return ok && cs == 0
}

func checkSum(ns string) (int, bool) {
	// create array of safely single-byte digits
	digits := make([]byte, 0, len(ns))
	for _, r := range ns {
		if unicode.IsDigit(r) {
			digits = append(digits, byte(r))
		}
	}
	if len(digits) == 0 {
		return 0, false
	}

	checksum, second := 0, false
	for i := len(digits) - 1; i >= 0; i-- { // for each digit, in reverse order...
		d := digits[i]
		n, _ := strconv.Atoi(string(d)) // ... convert it to an int, ignoring errors because we've pre-checked the digit.
		if second {                     // If it's a second digit ...
			n = fixUp(n) // ... fix it up.
		}
		checksum += n    // Add digit to checksum ...
		second = !second // ... and flip our every-second-digit flag.
	}
	return (checksum % 10), true
}

func fixUp(n int) int {
	n *= 2
	if n >= 10 {
		n -= 9
	}
	return n
}

// AddCheck adds a Luhn checksum to a numeric string.
func AddCheck(ns string) string {
	rawCs, _ := checkSum(ns + "0")
	correctCs := (10 - rawCs) % 10
	return ns + strconv.Itoa(correctCs)
}
