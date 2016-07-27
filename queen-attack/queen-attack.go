// Package queenattack provides a function to determine whether
// two queens can attack each in the next move
package queenattack

import (
	"fmt"
	"strings"
)

const (
	file = "abcdefgh"
	rank = "12345678"
)

// CanQueenAttack determines whether two pieces can attack on the same
// vertical, horizontal or diagonal
func CanQueenAttack(w string, b string) (canAttack bool, err error) {

	// check for obvious errors
	if len(w) != 2 || len(b) != 2 {
		return false, fmt.Errorf("invalid location - w: '%v', b: '%v'.", w, b)
	}
	if !(onBoard(w) && onBoard(b)) {
		return false, fmt.Errorf("off board - w: '%v', b: '%v'.", w, b)
	}
	if w == b {
		return false, fmt.Errorf("same square - w: '%v', b: '%v'.", w, b)
	}

	// if same file or rank, queen can attack
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}

	// if same diagonal, queen can attack
	if forwardDiagonal(w) == forwardDiagonal(b) || backwardDiagonal(w) == backwardDiagonal(b) {
		return true, nil
	}

	// otherwise queen can't attack
	return false, nil
}

// check file and rank both on board
func onBoard(loc string) bool {
	return strings.Contains(file, loc[0:1]) && strings.Contains(rank, loc[1:2])
}

// calculate forward diagonal from file and rank
func forwardDiagonal(loc string) int {
	f, r := fileRank(loc)
	return f - r
}

// calculate backward diagonal from file and rank
func backwardDiagonal(loc string) int {
	f, r := fileRank(loc)
	return f + r
}

// parse location into integer offsets
func fileRank(loc string) (f, r int) {
	return strings.Index(file, loc[0:1]), strings.Index(rank, loc[1:2])
}
