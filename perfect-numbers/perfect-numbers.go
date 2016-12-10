// Package perfect provides a function to classify natural numbers.
package perfect

import "errors"

const testVersion = 1

// Classification as a perfect number.
type Classification int

const (
	// ClassificationAbundant - sum of dividends > than n
	ClassificationAbundant Classification = iota

	// ClassificationDeficient - sum of dividends < n
	ClassificationDeficient

	// ClassificationPerfect - sum of dividends == n
	ClassificationPerfect
)

// ErrOnlyPositive - can't test 0 or negative numbers.
var ErrOnlyPositive = errors.New("perfect: cannot classify zero or negative numbers")

// Classify determines if a positive integer is deficient, perfect or abundant.
func Classify(n uint64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}

	var sum, i uint64
	for i = 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
