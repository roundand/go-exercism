// Package diffsquares provides functions for sum of squares and square of sums
package diffsquares

import "flag"

// SquareOfSums returns the square of the sum of the first n integers
func SquareOfSums(n int) int {
	s := sumOfSeries(n)
	return s * s
}

var forloop *bool

func init() {
  forloop = flag.Bool("forloop", false, "a bool")
  flag.Parse()
}

// sumOfSeries calculates the sum of a series
func sumOfSeries(n int) int {

  if ! *forloop {
    // the clever way -
    // on my box: BenchmarkSquareOfSums-4	2000000000	         1.30 ns/op
  	return n * (n + 1) / 2
  }

	// the obvious way -
  // on my box: BenchmarkSquareOfSums-4	30000000	        56.4 ns/op
  s := 0
  for i := 1; i <= n; i++ {
    s += i
  }
  return s
}

// SumOfSquares returns the sum of the squares of the first n integers
func SumOfSquares(n int) int {

	// use a simple for loop, the go way
	s := 0
	for i := 1; i <= n; i++ {
		s += i * i
	}
	return s
}

// Difference returns difference between SquareOfSums(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
