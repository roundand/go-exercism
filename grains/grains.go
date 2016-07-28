// Package grains provides functions to calculate
// the number of grains required to reqard the King's loyal servant
package grains

import (
	"flag"
	"fmt"
)

// allow switching optimisation strategy
var (
	fastSquare *bool
  fastTotal *bool
)

func init() {
	fastSquare = flag.Bool("fastSquare", false, "a bool")
  fastTotal = flag.Bool("fastTotal", false, "a bool")
	flag.Parse()
}

// Square calculates the number of grains on any individual chessboard square from 1 to 64.
func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, fmt.Errorf("cell '%n' is off the board", s)
	}

  // use optimised method?
  if *fastSquare {
    return uint64(1) << uint(s - 1), nil
  }

	if s == 1 {
		return 1, nil
	}

	ss, _ := Square(s - 1)
	return 2 * ss, nil
}

// Total calculates number of grains required for entire chessboard
func Total() uint64 {
  if *fastTotal {
    x := (uint64(1) << 63)
    y := (uint64(1) << 63) - 1
    return x + y
  }
	return addSquares(64)
}

// addSquares adds grains from each square on the board
func addSquares(s int) uint64 {
	if s == 1 {
		return 1
	}

	ss, _ := Square(s)
	return ss + addSquares(s-1)
}
