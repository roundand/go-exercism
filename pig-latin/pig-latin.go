// Package igpay provides a function to convert English to Pig Latin.
package igpay

import "strings"

// PigLatin converts an English word or sentence to Pig Latin.
func PigLatin(e string) string {
	// Using channels for a flow-based solution.
	in := make(chan string)
	out := make(chan string)

	go words(in, e)
	go rules(out, in)

	// assemble output and return it
	var ret []string
	for w := range out {
		ret = append(ret, w)
	}
	return strings.Join(ret, " ")
}

func words(in chan string, s string) {
	defer close(in)
	for _, w := range strings.Split(s, " ") {
		in <- w
	}
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func rules(out, in chan string) {
	defer close(out)
	for w := range in {
		switch {
		case string(w[0:2]) == "ch":
			out <- code2(w)
		case string(w[0:2]) == "qu":
			out <- code2(w)
		case string(w[0:2]) == "sq":
			out <- code3(w)
		case string(w[0:2]) == "sc" && isVowel(w[2]):
			out <- code2(w)
		case string(w[0:2]) == "sc":
			out <- code3(w)
		case string(w[0:2]) == "th" && isVowel(w[2]):
			out <- code2(w)
		case string(w[0:2]) == "th":
			out <- code3(w)
		case string(w[0:1]) == "y" && !isVowel(w[1]):
			out <- code0(w)
		case string(w[0:1]) == "x" && !isVowel(w[1]):
			out <- code0(w)
		case isVowel(w[0]):
			out <- code0(w)
		default:
			out <- code1(w)
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
