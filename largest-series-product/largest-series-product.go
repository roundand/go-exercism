// Package lsproduct provides a function to calculate the largest series product/
package lsproduct

import (
  "fmt"
	"strconv"
)

const testVersion = 3

func LargestSeriesProduct(numseq string, span int) (int, error) {

  // sanity tests
  if span > len(numseq) {
    return 0, fmt.Errorf("span %d is greater than length of sequence %q", span, len(numseq))
  }
  if span < 0 {
    return 0, fmt.Errorf("span %d is less than zero", span)
  }

	// convert to array of int
	ints := make([]int, len(numseq))
	for i, r := range numseq {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, err
		}
		ints[i] = n
	}

  // find max product
  max := 0
  for i := 0; i <= (len(ints) - span); i++ {
    x := 1
    for j := i; j < (i + span); j++ {
      x *= ints[j]
    }
    if x > max {
      max = x
    }
  }

  return max, nil
}
