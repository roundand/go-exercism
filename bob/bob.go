// package bob responds to conversational inputs likes a grumpy teenager.
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2 // same as targetTestVersion

/////////
//
// this version runs "go test - bench ." in 1.637 seconds on my laptop
//
/////////

// bob's sole function is to respond to Alice's inputs
func Hey(in string) string {

	// He answers 'Whoa, chill out!' if you yell at him.
	if in != strings.ToLower(in) && in == strings.ToUpper(in) {
		return "Whoa, chill out!"
	}

	// Bob answers 'Sure.' if you ask him a question.
	if isQuestion, _ := regexp.MatchString(`\?\s*$`, in); isQuestion {
		return "Sure."
	}

	// He says 'Fine. Be that way!' if you address him without actually saying anything.
	if hasWord, _ := regexp.MatchString(`\w`, in); !hasWord {
		return "Fine. Be that way!"
	}

	// He answers 'Whatever.' to anything else.
	return "Whatever."
}
