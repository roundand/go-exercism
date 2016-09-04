// Package romannumerals provides a function to convert numbers from Arabic numerals to Roman numerals.
package romannumerals

import (
	"fmt"
)

const (
	testVersion = 2
	MMMM        = 4000
)

type value struct {
	rom string
	val int
}

var values = []value{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts an int number to a Roman numeral string.
func ToRomanNumeral(num int) (string, error) {
	if num >= MMMM {
		return "", fmt.Errorf("number should be less than MMMM but was '%d'\n", num)
	}
	if num > 0 {
		return roman(num, values), nil
	}
	return "", fmt.Errorf("can't convert value '%d'\n", num)
}

func roman(num int, vals []value) string {
	switch {
	case num == 0:
		return ""
	case num >= vals[0].val:
		return vals[0].rom + roman(num-vals[0].val, vals)
	default:
		return roman(num, vals[1:])
	}
}
