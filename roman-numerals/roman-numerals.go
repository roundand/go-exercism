// Package romannumerals provides a function to convert numbers from Arabic numerals to Roman numerals.
package romannumerals

import (
  "fmt"
)

const (
	testVersion = 2
  MMMM        = 4000
	M           = 1000
  CM          = 900
	D           = 500
  CD          = 400
	C           = 100
  XC          = 90
	L           = 50
  XL          = 40
	X           = 10
  IX          = 9
	V           = 5
  IV          = 4
	I           = 1
)

// ToRomanNumeral converts an int number to a Roman numeral string.
func ToRomanNumeral(num int) (string, error) {
  if num >= MMMM {
    return "", fmt.Errorf("number should be less than MMMM but was '%d'\n", num)
  }
	if num >= M {
    r, _ := ToRomanNumeral(num - M)
		return "M" + r, nil
	}
  if num >= CM {
    r, _ := ToRomanNumeral(num - CM)
		return "CM" + r, nil
	}
	if num >= D {
    r, _ := ToRomanNumeral(num - D)
		return "D" + r, nil
	}
  if num >= CD {
    r, _ := ToRomanNumeral(num - CD)
		return "CD" + r, nil
	}
	if num >= C {
    r, _ := ToRomanNumeral(num - C)
		return "C" + r, nil
	}
  if num >= XC {
    r, _ := ToRomanNumeral(num - XC)
		return "XC" + r, nil
	}
	if num >= L {
    r, _ := ToRomanNumeral(num - L)
		return "L" + r, nil
	}
  if num >= XL {
    r, _ := ToRomanNumeral(num - XL)
		return "XL" + r, nil
	}
	if num >= X {
    r, _ := ToRomanNumeral(num - X)
		return "X" + r, nil
	}
  if num >= IX {
		return "IX", nil
	}
	if num >= V {
    r, _ := ToRomanNumeral(num - V)
		return "V" + r, nil
	}
  if num >= IV {
		return "IV", nil
	}
	if num > I {
    r, _ := ToRomanNumeral(num - I)
		return "I" + r, nil
	}
  if num == I {
		return "I", nil
	}
	return "", fmt.Errorf("can't convert value '%d'\n", num)
}
