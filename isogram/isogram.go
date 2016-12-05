// Package isogram provides a function to determine if a string is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram determines whether or not a string is an isogram.
func IsIsogram(s string) bool {
	letters := map[rune]bool{}
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			if letters[r] {
				return false
			}
			letters[r] = true
		}
	}
	return true
}
