// Package strain provides some pseudo-generic functions.
package strain

// Ints is a slice of int.
type Ints []int

// Lists is a slice of []int.
type Lists [][]int

// Strings is a slice of string.
type Strings []string

// Keep returns those members of an Ints list which pass a provided function.
func (list Ints) Keep(f func(int) bool) Ints {
	var r Ints
	for _, i := range list {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

// Discard returns those members of an Ints list which fail a provided function.
func (list Ints) Discard(f func(int) bool) Ints {
	var r Ints
	for _, i := range list {
		if !f(i) {
			r = append(r, i)
		}
	}
	return r
}

// Keep returns those members of an Ints list which pass a provided function.
func (list Lists) Keep(f func([]int) bool) Lists {
	var r Lists
	for _, i := range list {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

// Keep returns those members of an Strings list which pass a provided function.
func (list Strings) Keep(f func(string) bool) Strings {
	var r Strings
	for _, i := range list {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}
