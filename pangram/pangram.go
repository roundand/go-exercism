// Package pangram provides a function to check if a sentance is a pangram.
package pangram

import "strings"

const testVersion = 1

// IsPangram tests if an input string is an ascii pangram.
func IsPangram(s string) bool {
	found := map[rune]bool{}
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' { // ascii character
			found[r] = true
		}
	}
	return len(found) == (('z' - 'a') + 1) // did we catch them all?
}
