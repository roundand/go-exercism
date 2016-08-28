// Package palindrome provides a function and return type for calculating palindromic products.
package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product int // palindromic, of course
	// list of all possible two-factor factorizations of Product, within
	// given limits, in order
	Factorizations [][2]int
}

// Products calculates minimum and maximum palindromic products of the specified minimum and maximum factors.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	// validations
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(p) {
				if len(pmin.Factorizations) == 0 || p < pmin.Product { // new minimum product
					pmin = Product{p, [][2]int{{i, j}}}
				} else if p == pmin.Product { // new factorisation for minimum product
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}
				if len(pmax.Factorizations) == 0 || p > pmax.Product { // new maximum product
					pmax = Product{p, [][2]int{{i, j}}}
				} else if p == pmax.Product { // new factorisation for maximum product
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
			}
		}
	}

	if len(pmin.Factorizations) == 0 {
		return Product{}, Product{}, errors.New("No palindromes")
	}

	return pmin, pmax, nil
}

func isPalindrome(p int) bool {
	s := strconv.Itoa(p)
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
