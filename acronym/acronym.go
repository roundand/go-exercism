// Package acronym provides a function to create acronyms from phrases.
package acronym

import "strings"

const testVersion = 1

const (
	ignoreSpace  = iota // at start, or after space - goes to ignoreLower on non-blank
	ignoreLower         // ignore next letter unless it's an upper - goes to ignoreSpace on space, or ignoreLetter on upper
	ignoreLetter        // ignore next letter - goes to ignoreSpace on space, or ignoreLower on lower
)

// abbreviate creates an upper case acronym from a list of words seperated by spaces.
func abbreviate(words string) string {
	out := []rune{}      // output
	state := ignoreSpace // should we capture the next letter?
	for _, r := range words {
		switch {
		case r == ' ' || r == '-':
			// regardless of previous state, we're in space - catch next letter of any case
			state = ignoreSpace
		case state == ignoreSpace:
			// First letter following spaces - capture it, and ignore the next whether uppercase or lower
			out = append(out, r)
			state = ignoreLetter
		case state == ignoreLetter && strings.ToUpper(string(r)) != string(r):
			// wasn't uppercase - we'll catch the next uppercase, if any
			state = ignoreLower
		case state == ignoreLower && strings.ToLower(string(r)) != string(r):
			// capture a camelCase upper
			out = append(out, r)
			state = ignoreLetter
		}
	}
	return strings.ToUpper(string(out))
}
