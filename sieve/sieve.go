// Package sieve proides a function that generates primes using the Sieve of Eratosthenes.
package sieve

// Sieve returns all primes up to the specified limit.
func Sieve(limit int) []int {
	var results []int

	// create sieve of non-primes
	sieve := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !sieve[i] { // if this value is not blocked by the sieve ...
			results = append(results, i)       // ... add it to the result set ...
			for j := 2 * i; j < limit; j += i { // and block all multiples of this value
				sieve[j] = true
			}
		}
	}
	return results
}
