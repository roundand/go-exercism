// package bob responds to conversational inputs likes a grumpy teenager.
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2 // same as targetTestVersion

/////////
//
// this version runs "go test - bench ." in avg 1.22 seconds on my laptop
//
/////////

// bob's sole function is to respond to Alice's inputs
func Hey(in string) string {

	var isQuestion = regexp.MustCompile(`\?\s*$`)
	var hasWord = regexp.MustCompile(`\w`)

	// He answers 'Whoa, chill out!' if you yell at him.
	if in != strings.ToLower(in) && in == strings.ToUpper(in) {
		return "Whoa, chill out!"
	}

	// Bob answers 'Sure.' if you ask him a question.
	if isQuestion.MatchString(in) {
		return "Sure."
	}

	// He says 'Fine. Be that way!' if you address him without actually saying anything.
	if !(hasWord.MatchString(in)) {
		return "Fine. Be that way!"
	}

	// He answers 'Whatever.' to anything else.
	return "Whatever."
}
