// Package beer provides a function to generate the first 99 verses of the beer song.
package beer

import (
	"fmt"
)

const testVersion = 1

const vN = "%d bottles of beer on the wall, %[1]d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n"
const v2 = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
const v1 = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
const v0 = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"

// Verse generates individual verses of the beer song.
func Verse(v int) (string, error) {
	if v < 0 || v > 99 {
		return "", fmt.Errorf("No such verse: %d\n", v)
	}
	switch v {
	case 2:
		return v2, nil
	case 1:
		return v1, nil
	case 0:
		return v0, nil
	default:
		return fmt.Sprintf(vN, v, v-1), nil
	}
}

// Verses generate a range of verses of the beer song.
func Verses(from, to int) (string, error) {
	if from < to {
		return "", fmt.Errorf("From and To values out of order - from: %d, to: %d.\n", from, to)
	}
	var r = ""
	for v := from; v >= to; v-- {
		s, err := Verse(v)
		if err != nil {
			return "", err
		}
		if r == "" {
			r = s
		} else {
			r += "\n" + s
		}
	}
	return r + "\n", nil
}

// Song generates the entire beers song
func Song() string {
	r, _ := Verses(99, 0)
	return r
}
