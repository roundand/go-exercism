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
	sep := " "

	// if we have a nounPhrase, add that to the rest of the verse
	if nounPhrase != "" {
		return EmbedSep(Verse(subject, relClauses, ""), nounPhrase, sep)
	}

	// if no nounPhrase but we have a subject, add rest of verse to subject
	if subject != "" {
		return joinStringsSep(subject, Verse("", relClauses, ""), sep)
	}

	// if no nounPhrase or subject but we have multiple relClauses, add first to rest
	if len(relClauses) > 1 {
		return joinStringsSep(relClauses[0], Verse("", relClauses[1:], ""), sep)
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
	{"the horse and the hound and the horn", "belonged to"},
	{"the farmer sowing his corn", "kept"},
	{"the rooster that crowed in the morn", "woke"},
	{"the priest all shaven and shorn", "married"},
	{"the man all tattered and torn", "kissed"},
	{"the maiden all forlorn", "milked"},
	{"the cow with the crumpled horn", "tossed"},
	{"the dog", "worried"},
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

	song := genSong(nounPhrases)
	fmt.Printf("song: %v\n", song)
	return song
}

/*
three cases:

#lastline - intro + noun

#penultimate - intro + noun + that + #lastline

#otherline - intro + noun + that + "\n" + rest
*/
func genSong(nounPhrases []NounPhrase) string {

	// if this is the last entry, generate last verse and be done.
	if (len(nounPhrases) == 1) {
		return joinPhrases(intro, nounPhrases[0].noun)
	}

	// generate remainder of song for remaining entries then generate verse for current entry
	return genSong(nounPhrases[1:]) + "\n\n" + genVerse(nounPhrases)
}

/*
three cases:

#lastline - intro + noun

#penultimate - intro + noun + that + #lastline

#otherline - intro + noun + that + "\n" + rest
*/
func genVerse(nounPhrases []NounPhrase) string {
	if len(nounPhrases) == 1 {
		return joinPhrases(intro, nounPhrases[0].noun)
	}

	return joinPhrases(intro, nounPhrases[0].noun) + "\n" + genLines(nounPhrases)
}

func genLines(nounPhrases []NounPhrase) string {
	if len(nounPhrases) == 0 {
		return ""
	}

	if len(nounPhrases) == 1 {
		return nounPhrases[0].noun
	}

	if len(nounPhrases) == 2 {
		return that + " " + nounPhrases[0].phrase + " " + genLines(nounPhrases[1:])
	}

	return that + " " + nounPhrases[0].phrase + " " + nounPhrases[1].noun + "\n" + genLines(nounPhrases[1:])

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
