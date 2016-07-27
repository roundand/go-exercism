// Package Slice provides functions to report series of characters from strings
package slice

// function All returns an array of substrings of specified length from a string
func All(n int, s string) []string {
	c := 1 + len(s) - n      // how many substrings will we generate?
	r := make([]string, c)   // allocate our result slice
	for x := 0; x < c; x++ { // populate result slice
		r[x] = s[x : x+n]
	}
	return r
}

// function UnsafeFirst returns firsts substring of specified length from a string
func UnsafeFirst(n int, s string) string {
	return s[:n] // fingers in ears, eyes closed, "la la la la la ..."
}

// function First returns first substring of specified length from a string, plus an error indicator
func First(n int, s string) (series string, ok bool) {
	if n > len(s) {
		return "", false // not possible
	}
	return s[:n], true // that's more like it
}
