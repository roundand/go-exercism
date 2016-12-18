// Package matrix implements row, col and constructor functions for rectangular integer matrices.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

const testVersion = 1

// Matrix is a rectangular array of integers.
type Matrix [][]int

// New parses input into a rectangular integer array and returns it
func New(s string) (Matrix, error) {
	m := Matrix{}
	rows := strings.Split(s, "\n")
	if len(rows) == 0 {
		return m, fmt.Errorf("zero rows found in %q\n", s) // need non-zero rows
	}

	// loop through rows and columns
	for _, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), " ")
		if len(cols) == 0 {
			return m, fmt.Errorf("zero cols found in %q\n", row) // need non-zero cols
		}
		ints := []int{}
		for _, cell := range cols {
			i, err := strconv.Atoi(cell)
			if err != nil {
				return m, fmt.Errorf("unparseable value found in %q\n", cell)
			}
			ints = append(ints, i)
		}
		m = append(m, ints)
	}

	// check that array is rectangular
	x := len(m[0])
	for y := range m {
		if len(m[y]) != x {
			return m, fmt.Errorf("irregular row length found in %q\n", s)
		}
	}
	return m, nil
}

// Row returns one row of a matrix
func (m Matrix) Row(i int) []int {
	r := make([]int, len(m[i]))
	copy(r, m[i])
	return r
}

// Rows returns all rows of a matrix
func (m Matrix) Rows() [][]int {
	res := make([][]int, len(m))
	for i := range m {
		res[i] = m.Row(i)
	}
	return res
}

// Col returns one column of a matrix
func (m Matrix) Col(i int) []int {
	col := make([]int, len(m))
	for x := range m {
		col[x] = m[x][i]
	}
	return col
}

// Cols returns all columns of a matrix
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := range m[0] {
		cols[i] = m.Col(i)
	}
	return cols
}

// Set sets one cell of a matrix
func (m Matrix) Set(row, col, val int) bool {
	if row >= len(m) || row < 0 || col >= len(m[0]) || col < 0 {
		return false
	}
	m[row][col] = val
	return true
}
