// Package igpay provides a function to convert English to Pig Latin.
package igpay

import "strings"

type key int
type word struct {
	k key
	v string
}

const (
	c0 key = iota
	c1
	c2
	c3
)

// PigLatin converts an English word or sentence to Pig Latin.
func PigLatin(e string) string {
	// Using channels for a flow-based solution.
	tokens := make(chan string)
	words := make(chan word)
	pig := make(chan string)

	go lex(tokens, e)
	go parse(words, tokens)
	go encode(pig, words)

	// assemble output and return it
	var ret []string
	for w := range pig {
		ret = append(ret, w)
	}
	return strings.Join(ret, " ")
}

func lex(tokens chan string, e string) {
	defer close(tokens)
	for _, t := range strings.Split(e, " ") {
		tokens <- t
	}
}

func parse(words chan word, tokens chan string) {
	defer close(words)
	for t := range tokens {
		var k key
		switch {
		case string(t[0:2]) == "ch":
			k = 2
		case string(t[0:2]) == "qu":
			k = 2
		case string(t[0:2]) == "sq":
			k = 3
		case string(t[0:2]) == "sc" && isVowel(t[2]):
			k = 2
		case string(t[0:2]) == "sc":
			k = 3
		case string(t[0:2]) == "th" && isVowel(t[2]):
			k = 2
		case string(t[0:2]) == "th":
			k = 3
		case string(t[0:1]) == "y" && !isVowel(t[1]):
			k = 0
		case string(t[0:1]) == "x" && !isVowel(t[1]):
			k = 0
		case isVowel(t[0]):
			k = 0
		default:
			k = 1
		}
		words <- word{k, t}
	}
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func encode(pig chan string, words chan word) {
	defer close(pig)
	for w := range words {
		switch w.k {
		case c3:
			pig <- code3(w.v)
		case c2:
			pig <- code2(w.v)
		case c1:
			pig <- code1(w.v)
		default:
			pig <- code0(w.v)
		}
	}
}

func code0(w string) string {
	return w + "ay"
}

func code1(w string) string {
	return w[1:] + w[0:1] + "ay"
}

func code2(w string) string {
	return w[2:] + w[0:2] + "ay"
}

func code3(w string) string {
	return w[3:] + w[0:3] + "ay"
}
