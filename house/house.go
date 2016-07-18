// Package house provides functions for constructing recursive nursery rhymes.
package house

import (
	"fmt"
)

// Embed adds a noun phrase to a relative phrase and returns the result.
func Embed(relPhrase, nounPhrase string) string {
	return joinPhrases(relPhrase, nounPhrase)
}

// EmbedSep adds a noun phrase to a relative phrase using the provided seperator and returns the result.
func EmbedSep(relPhrase, nounPhrase, sep string) string {
	return joinStringsSep(relPhrase, nounPhrase, sep)
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure, as one long line.
func Verse(subject string, relClauses []string, nounPhrase string) string {
	return verseSep(subject, relClauses, nounPhrase, " ")
}

// verseSep generates a verse of a song with relative clauses that have
// a recursive structure, using a parameterised line seperator.
func verseSep(subject string, relClauses []string, nounPhrase, sep string) string {

	// if we have a nounPhrase, add that to the rest of the verse
	if nounPhrase != "" {
		return EmbedSep(verseSep(subject, relClauses, "", sep), nounPhrase, sep)
	}

	// if no nounPhrase but we have a subject, add rest of verse to subject
	if subject != "" {
		return joinStringsSep(subject, verseSep("", relClauses, "", sep), sep)
	}

	// if no nounPhrase or subject but we have multiple relClauses, add first to rest
	if len(relClauses) > 1 {
		return joinStringsSep(relClauses[0], verseSep("", relClauses[1:], "", sep), sep)
	}

	// if just one relClauses, return it
	if len(relClauses) > 0 {
		return relClauses[0]
	}

	// edge case - if we start off with no relClauses, return empty string
	return ""
}

type NounPhrase struct {
	noun string
	phrase string
}

/*
This is the horse and the hound and the horn
that belonged to the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
*/
var nounPhrases = []NounPhrase {
	{"the cat", "killed"},
	{"the rat", "ate"},
	{"the malt", "lay in"},
	{"the house that Jack built.", ""},
}

var (
	intro = "This is"
	that = "that"
)

// Song generates all verses of a recursively structured song.
func Song() string {

	var phrases []string = make([]string, len(nounPhrases))
	for i, np := range(nounPhrases) {
		phrases[i] = np.noun + " " + that + " " + np.phrase
	}

	song := genSong(nounPhrases, phrases)
	fmt.Printf("song: %v\n", song)
	return song
}

func genSong(nounPhrases []NounPhrase, phrases []string) string {

	// if this is the last entry, generate last verse and be done.
	if (len(nounPhrases) == 1) {
		return verseSep(joinPhrases(intro, nounPhrases[0].noun), phrases[0:0], "", "\n")
	}

	// generate remainder of song for remaining entries then generate verse for current entry
	return genSong(nounPhrases[1:], phrases[1:]) + "\n\n" + verseSep(intro, phrases, "", "\n")
}

// joinPhrases joins two strings, putting a space between them
// if both are non-empty.
func joinPhrases(p1, p2 string) string {
	return joinStringsSep(p1, p2, " ")
}

// joinStringsSep joins two strings, putting a seperator between them
// if both are non-empty.
func joinStringsSep(p1, p2, sep string) string {
	if p1 != "" && p2 != "" {
		return p1 + sep + p2
	}
	if p1 != "" {
		return p1
	}
	return p2
}
