// Package wordcount provides a function to count words in a string.
package wordcount

import (
  "sort"
  "strings"
  "unicode"
)

const testVersion = 2

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {

  // normalise by forcing into lower case and stripping punctuation
  var np []rune //non-punctuation marks
  for _, r := range(phrase) {
    if !(unicode.IsPunct(r) || unicode.IsSymbol(r)) {
      np = append(np, unicode.ToLower(r))
    }
  }
  phrase = string(np)

  // create frequency map of words
  words := strings.Split(phrase, " ")
  sort.Strings(words)
  frequency := make(Frequency, 0)
  for _, word := range(words) {
    if word != "" {
      frequency[word]++
    }
  }

  return frequency
}
