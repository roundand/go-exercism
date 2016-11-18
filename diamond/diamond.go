// Package diamond provides a function to print text diamonds for specified letters.
package diamond

import (
	"fmt"
	"strings"
)

const testVersion = 1

// Gen returns the text diamond corresponding to the specified letter.
func Gen(letter byte) (string, error) {
	if !(letter >= 'A' && letter <= 'Z') {
		return "", fmt.Errorf("letter: %q outside range 'A' - 'Z'\n", letter)
	}

	ret := ""
	for c := letter; c >= 'A'; c-- {
		if c == letter { // if we hace the initial letter
			ret = line(c, letter-'A') // create centre line
		} else { // otherwise
			l := line(c, letter-'A') // generate next line
			ret = l + ret + l        // wrap existing lines with preceding and following lines
		}
	}
	// diagnostics
	fmt.Printf("Diamond for %q:\n%s\n", letter, strings.Replace(ret, " ", "·", -1))
	return ret, nil
}

func line(c byte, width byte) string {
	s := spread(c)
	padding := strings.Repeat(" ", int(width+'A'-c))
	return fmt.Sprintf("%s%s%s\n", padding, s, padding)
}

func spread(c byte) string {
	if c == 'A' {
		return "A"
	}
	// internal padding for 'B' and on goes 1, 3, 5...
	padding := 1 + int(c-'B')*2
	return fmt.Sprintf("%c%s%c", c, strings.Repeat(" ", padding), c)
}
