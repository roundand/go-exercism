// Package matrix supports the detection of saddle points in matrices.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

const testVersion = 1

// Matrix is a rectangular array of integers.
type Matrix [][]int

// New parses input into a rectangular slice of integer slices, and returns it
func New(s string) (*Matrix, error) {
	m := Matrix{}
	rows := strings.Split(s, "\n")
	if len(rows) == 0 {
		return nil, fmt.Errorf("zero rows found in %q\n", s) // need non-zero rows
	}

	// loop through rows, then items in rows
	for _, row := range rows {
		items := strings.Split(strings.Trim(row, " "), " ")

		if len(items) == 0 {
			return nil, fmt.Errorf("zero items found in %q\n", row)
		}

		// check that pseudo-array is rectangular
		if len(m) > 0 && len(items) != len(m[0]) {
			return nil, fmt.Errorf("irregular row length found in %q\n", s)
		}

		// parse items into a slice of integers, and add it to the matrix
		ints := []int{}
		for _, item := range items {
			i, err := strconv.Atoi(item)
			if err != nil {
				return nil, fmt.Errorf("unparseable value found in %q\n", item)
			}
			ints = append(ints, i)
		}
		m = append(m, ints)
	}

	return &m, nil
}

// Col returns one column of a matrix
func (m *Matrix) Col(i int) []int {
	col := make([]int, len(*m))
	for x := range *m {
		col[x] = (*m)[x][i]
	}
	return col
}

// Pair is an array - not a slice - of two integers.
type Pair [2]int

// Saddle returns a list of all if any saddle points in a matrix.
func (m *Matrix) Saddle() []Pair {
	points := []Pair{}

	// for every row in matrix ...
	for r, row := range *m {

		// ... for each element in row ...
		for c, val := range row {

			// ... if it's a saddle point, note it.
			if val == highest(row) && val == lowest(m.Col(c)) {
				points = append(points, Pair{r, c})
			}
		}
	}

	return points
}

// func highest returns highest value in an integer slice
func highest(seq []int) int {
	return most(seq, func(x, y int) bool { return x > y })
}

// func lowest returns lowest value in an integer slice
func lowest(seq []int) int {
	return most(seq, func(x, y int) bool { return x < y })
}

// most returns most more() value in an integer slice
func most(seq []int, more func(x, y int) bool) int {
	m := seq[0]
	for _, v := range seq {
		if more(v, m) {
			m = v
		}
	}
	return m
}
