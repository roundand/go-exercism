// Package prime provides a function to calculate the Nth prime.
package prime

// Nth calculates the nth prime.
func Nth(n int) (int, bool) {
	// validate n
	if n < 1 {
		return 0, false
	}
	// use Sieve of Eratosthenes
	p := []int{2}
	for x := 3; len(p) < n; x++ {
		isPrime := true
		for _, y := range p {
			if x%y == 0 {
				isPrime = false
				break // x ain't prime, move along now...
			}
		}
		if isPrime {
			p = append(p, x)
		}
	}
	return p[len(p)-1], true
}
