// Package matrix extension to support the detection of saddle points in matrices.
package matrix

// Pair is an array - not a slice - of two integers.
type Pair [2]int

// Saddle returns a list of any and all saddle points in a matrix.
func (m *Matrix) Saddle() []Pair {
	points := []Pair{}

	// cache lowest values for each column
	loCol := make([]int, len((*m)[0]))
	for c := range (*m)[0] {
		loCol[c] = lowest(m.Col(c))
	}

	// for every value in every row ...
	for r, row := range *m {

		// ... if it's a saddle point, note it.
		hiRow := highest(row) // cache highest value in row
		for c, val := range row {
			if val == hiRow && val == loCol[c] {
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
