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
	nt := normalise(pt)
	r, c := calcRowsCols(len(nt))

	// create encrypted rectangle
	// reading normalised rectangle row by row
	// and writing to encrypted rectangle col by col
	var enc []string = make([]string, c) // encrypted output
	var cur int = 0                      // current linear position
	for row := 0; row < r; row++ {       // which row of the (virtual) rectangle are we reading?
		for col := 0; col < c && cur < len(nt); col++ { // which column are we reading?
			enc[col] += string(nt[cur]) // write current character, using column as row
			cur++
		}
	}

	return strings.Join(enc, " ")
}

// normalised text is stripped of punctuation and forced to lower-case
func normalise(pt string) string {
	nt := ""
	for _, rn := range pt {
		if unicode.IsLetter(rn) || unicode.IsDigit(rn) {
			nt += string(rn)
		}
	}
	return strings.ToLower(nt)
}

// calcRowsCols calculates the squarest rectangle that will contain the message.
// "The size of the rectangle (`r x c`) should be decided by the length of the message,
// such that `c >= r` and `c - r <= 1`, where `c` is the number of columns
// and `r` is the number of rows."
func calcRowsCols(l int) (r, c int) {
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
