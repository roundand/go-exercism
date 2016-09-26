// Package igpay provides a function to convert English to Pig Latin.
package igpay

import "strings"

// PigLatin converts an English word or sentence to Pig Latin.
func PigLatin(e string) string {
	// Using channels for a flow-based solution.
	in := make(chan string)
	pl0 := make(chan string)
	pl1 := make(chan string)
	pl2 := make(chan string)
	pl3 := make(chan string)
	out := make(chan string)

	go words(in, e)
	go rules(pl0, pl1, pl2, pl3, in)
	go code0(out, pl0)
	go code1(out, pl1)
	go code2(out, pl2)
	go code3(out, pl3)

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

func rules(pl0, pl1, pl2, pl3, in chan string) {
	defer close(pl0)
	defer close(pl1)
	defer close(pl2)
	defer close(pl3)
	for w := range in {
		switch {
		case string(w[0:2]) == "ch":
			pl2 <- w
		case string(w[0:2]) == "qu":
			pl2 <- w
		case string(w[0:2]) == "sq":
			pl3 <- w
		case string(w[0:2]) == "sc" && isVowel(w[2]):
			pl2 <- w
		case string(w[0:2]) == "sc":
			pl3 <- w
		case string(w[0:2]) == "th" && isVowel(w[2]):
			pl2 <- w
		case string(w[0:2]) == "th":
			pl3 <- w
		case isVowel(w[0]):
			pl0 <- w
		default:
			pl1 <- w
		}
	}
}

func code0(out, pl chan string) {
	defer close(out)
	for w := range pl {
		out <- w + "ay"
	}
}

func code1(out, pl chan string) {
	//	defer close(out)
	for w := range pl {
		out <- w[1:] + w[0:1] + "ay"
	}
}

func code2(out, pl chan string) {
	//	defer close(out)
	for w := range pl {
		out <- w[2:] + w[0:2] + "ay"
	}
}

func code3(out, pl chan string) {
	//	defer close(out)
	for w := range pl {
		out <- w[3:] + w[0:3] + "ay"
	}
}
