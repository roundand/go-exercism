// Package anagram provides a function to detect anagrams from a list of candidates.
package anagram

import (
	"sort"
	"strings"
)

// Normalise sorts the characters in a string
func Normalise(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

// Detect returns a list of lower-case anagrams of a specified word, filtered from a supplied list of candidates.
func Detect(word string, candidates []string) []string {
	var anagrams []string
	lWord := strings.ToLower(word)
	nWord := Normalise(lWord)

	for _, candidate := range candidates {
		lCandidate := strings.ToLower(candidate)
		switch {
			case len(word) != len(candidate): // anagrams must be same length
				continue
			case lWord == lCandidate: // a word is not an anagram of itself
				continue
			case nWord == Normalise(lCandidate): // we have a match
				anagrams = append(anagrams, lCandidate)
		}
	}
	return anagrams
}
