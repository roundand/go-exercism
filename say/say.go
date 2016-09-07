// Package say provides a function to convert integers to prose strings.
package say

import (
	"fmt"
)

type unit struct {
	name  string
	value uint64
}

// useful vars
var (
	n20 = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	n100  = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	units = []unit{
		{"thousand", 1e3},
		{"million", 1e6},
		{"billion", 1e9},
		{"trillion", 1e12},
		{"quadrillion", 1e15},
		{"quintillion", 1e18},
	}
)

// Say converts an integer to an English string.
func Say(n uint64) string {
	if n < 1000 {
		return sayLow(n)
	}
	for _, u := range units {
		if n < u.value*1000 || u.value == 1e18 {
			return sayHigh(n, u)
		}
	}

	panic(fmt.Sprintf("Can't match number '%d'.\n", n))
	return ""
}

// sayLow converts irregular numbers, lower than 1000, to strings.
func sayLow(n uint64) string {
	if n < 20 {
		return n20[n]
	}
	if n < 100 {
		r := n % 10
		if r == 0 {
			return n100[n/10]
		}
		return fmt.Sprintf("%s-%s", n100[n/10], Say(r))
	}
	c := n / 100
	r := n % 100
	if r == 0 {
		return fmt.Sprintf("%s hundred", Say(c))
	}
	return fmt.Sprintf("%s hundred %s", Say(c), Say(r))
}

// sayHigh converts regular numbers, the part larger than 1000, to strings.
func sayHigh(n uint64, u unit) string {
	c := n / u.value
	r := n % u.value
	if r == 0 {
		return fmt.Sprintf("%s %s", Say(c), u.name)
	}
	return fmt.Sprintf("%s %s %s", Say(c), u.name, Say(r))
}
