// Package cryptosquare exports a function to encrypt plaintext using the square code method.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode returns the square code encrypted version of a plaintext parameter.
func Encode(pt string) string {

	// normalised text is stripped of punctuation and forced to lower-case
	nt := ""
	for _, rn := range pt {
		if unicode.IsLetter(rn) || unicode.IsDigit(rn) {
			nt += strings.ToLower(string(rn))
		}
	}

	// calculate row and column lengths
	r, c := calcSides(len(nt))

	// create normalised rectangle
	var nr []string
	for i := 0; i < r; i++ {
		start := i * c
		// last row could be short
		end := int(math.Min(float64(start+c), float64(len(nt))))
		nr = append(nr, nt[start:end])
	}

	// create encrypted rectangle
	// reading normalised rectangle column by column
	// and writing to encrypted rectangle row by row
	var er []string = make([]string, c)
	for ci := 0; ci < c; ci++ {
		for ri := 0; ri < r; ri++ {
			if len(nr[ri]) > ci {
				er[ci] += nr[ri][ci : ci+1]
			}
		}
	}

	// return
	return strings.Join(er, " ")
}

// calcSides calculates the squarest rectangle that will contain the message.
// "The size of the rectangle (`r x c`) should be decided by the length of the message,
// such that `c >= r` and `c - r <= 1`, where `c` is the number of columns
// and `r` is the number of rows."
func calcSides(l int) (r, c int) {
	r, c = int(math.Sqrt(float64(l))), 0
	if (r * r) == l {
		c = r
	} else if r*(r+1) > l {
		c = r + 1
	} else {
		r, c = r+1, r+1
	}
	return r, c
}
