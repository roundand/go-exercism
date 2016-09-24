// Package phonenumber provides a function to clean up US phone numbers.
package phonenumber

import "fmt"

// Number cleans up a US-style phone number.
func Number(s string) (string, error) {
	d := getDigits(s)
	switch {
	case len(d) < 10:
		return "", fmt.Errorf("too few digits in %q (%q)", s, d)
	case len(d) > 11:
		return "", fmt.Errorf("too many digits in %q (%q)", s, d)
	case len(d) == 10:
		return string(d), nil
	case d[0] == '1':
		return string(d[1:]), nil
	}
	return "", fmt.Errorf("11 digit number not starting with '1' %q", s)
}

// AreaCode returns a US-style area code
func AreaCode(s string) (string, error) {
	d, e := Number(s)
	if e != nil {
		return d, e
	}
	return fmt.Sprintf("%s", string(d[0:3])), nil
}

// Format returns a US-style formatted phone number.
func Format(s string) (string, error) {
	d, e := Number(s)
	if e != nil {
		return d, e
	}
	return fmt.Sprintf("(%s) %s-%s", string(d[0:3]), string(d[3:6]), string(d[6:])), nil
}

func getDigits(s string) string {
	d := []rune{}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			d = append(d, r)
		}
	}
	return string(d)
}
