// Package prime returns a function which generates the prime factors of a number.
package prime

const testVersion = 2

// Factors returns a slice of int64 prime factors of an int64 number.
func Factors(n int64) []int64 {
	f := []int64{}
	for i := int64(2); i <= n; {
		x := n / i // integer division, of course
		if x*i != n {
			i++ // not a clean division, so try next divisor
		} else { // clean division, so ...
			f = append(f, i) // note the prime factor,
			n = x            // and analyse remaining factors
		}
	}
	return f
}
