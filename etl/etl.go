// Package etl provides a tranform function for use in an ETL project.
package etl

import (
	"strings"
)

// Transform converts a scrabble hand represented in one data schema
// to the same hand using a different schema
func Transform(input oldstyle) newstyle {

	out := make(newstyle)

	for value, letters := range input {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = value
		}
	}
	return out
}

// go will ducktype these local input and output types to etl_test types
type oldstyle map[int][]string
type newstyle map[string]int
