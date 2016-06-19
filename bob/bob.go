// package bob responds to conversational inputs likes a grumpy teenager.
package bob // package name must match the package name in bob_test.go

import(
  "regexp"
  "strings"
)

const testVersion = 2 // same as targetTestVersion

// bob's sole function is to respond to Alice's inputs
func Hey(in string) string {
  hasLetters, _ := regexp.MatchString(`\pL`, in) // unicode letter character class
  hasLetters, _ := regexp.MatchString(`\pL`, in) // unicode letter character class
  hasDigits, _ := regexp.MatchString(`\pN`, in) // unicode number character class
  isQuestion, _ := regexp.MatchString(`\?\s*$`, in)
  // He answers 'Whoa, chill out!' if you yell at him.
  if (hasLetters && in == strings.ToUpper(in)) {
    return "Whoa, chill out!"
  }

  // Bob answers 'Sure.' if you ask him a question.
  if (isQuestion) {
    return "Sure."
  }

  // He says 'Fine. Be that way!' if you address him without actually saying anything.
  if !(hasLetters || hasDigits) {
    return "Fine. Be that way!"
  }

  // He answers 'Whatever.' to anything else.
  return "Whatever."
}
