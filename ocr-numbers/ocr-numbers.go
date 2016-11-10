// Package ocr provides functions to scan visual images of digits.
package ocr

import (
	"strings"
)

const (
	d0 = " _ | ||_|   "
	d1 = "     |  |   "
	d2 = " _  _||_    "
	d3 = " _  _| _|   "
	d4 = "   |_|  |   "
	d5 = " _ |_  _|   "
	d6 = " _ |_ |_|   "
	d7 = " _   |  |   "
	d8 = " _ |_||_|   "
	d9 = " _ |_| _|   "
)

// Recognize scans a multi-row string image and returns an array of recognised digits, or an error.
func Recognize(image string) []string {
	ret := []string{}
	ss := strings.Split(image, "\n")[1:] // split by row, discarding initial \n

	// for each line of digits (4 rows of characters)
	for x := 0; x < len(ss); x += 4 {
		line := ""

		// for each digit (3 chars in each row)
		for y := 0; y < len(ss[x+1]); y += 3 {

			// for each row of characters in this line of digits
			digit := ""
			for z := x; z < x+4; z++ {
				digit += ss[z][y : y+3]
			}

			line += recognizeDigit(digit)
		}

		ret = append(ret, line)
	}

	return ret
}

// recognizeDigit - optimised for speed.
func recognizeDigit(d string) string {
	switch d {
	case d0:
		return "0"
	case d1:
		return "1"
	case d2:
		return "2"
	case d3:
		return "3"
	case d4:
		return "4"
	case d5:
		return "5"
	case d6:
		return "6"
	case d7:
		return "7"
	case d8:
		return "8"
	case d9:
		return "9"
	}
	return "?"
}
