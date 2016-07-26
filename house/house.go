// Package house provides functions for constructing recursive nursery rhymes.
package house

/*

an ad-hoc Song grammar -

SONG   			::= VERSES
VERSES 			::= VERSE [BLANKLINE VERSES]
VERSE  			::= SUBJECT [NEWLINE RELPHRASES] STOP
SUBJECT     ::= INTRO phrase
RELPHRASES 	::= RELPHRASE [NEWLINE RELPHRASES]
RELPHRASE 	::= REF link phrase

INTRO				::= "This is"
REF				 	::= "that"
STOP				::= "."
BLANKLINE   ::= "\n\n"
NEWLINE			::= "\n"

*/


type Topic struct {
	phrase string
	link string
}

var topics = []Topic{
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
	{"the house that Jack built", ""},
}

var (
	intro = "This is"
	ref  = "that"
	stop = "."
	blankline = "\n\n"
	newline = "\n"
)

/*
Song generates all verses of a recursively structured song.

SONG   			::= VERSES
*/
func Song() string {

	song := genVerses(topics)
	return song
}

/*
genVerses recursively generated verses from topics, but in reverse order

VERSES 			::= VERSE [BLANKLINE VERSES]
*/
func genVerses(topics []Topic) string {

	// if this is the last entry, generate last verse and be done.
	if len(topics) == 1 {
		return genVerse(topics)
	}

	// generate remainder of song for remaining entries then generate verse for current entry
	return genVerses(topics[1:]) + blankline + genVerse(topics)
}

/*
genVerse generates a single verse from topics

VERSE  			::= SUBJECT [NEWLINE RELPHRASES] STOP
SUBJECT     ::= INTRO phrase
*/
func genVerse(topics []Topic) string {
	if len(topics) == 1 {
		return compose(intro, genRelPhrases(topics)) + stop
	}

	return compose(intro, topics[0].phrase) + newline + genRelPhrases(topics) + stop
}

/*
genRelPhrases generates a sequence of relative phrases from topics

RELPHRASES 	::= RELPHRASE [NEWLINE RELPHRASES]
RELPHRASE 	::= REF link phrase
*/
func genRelPhrases(topics []Topic) string {
	if len(topics) == 1 {
		return topics[0].phrase
	}

	if len(topics) == 2 {
		return compose(ref, topics[0].link, topics[1].phrase)
	}

	return compose(ref, topics[0].link, topics[1].phrase) + newline + genRelPhrases(topics[1:])
}

// compose strings using a space seperator, recursively of course
func compose(items ...string) string {
	if len(items) == 1 {
		return items[0]
	}
	return items[0] + " " + compose(items[1:]...)
}

//
// non-Song tests
//

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
