// Package foodchain generates an old lady who ate a fly
package foodchain

const testVersion = 2

type verse struct{ item, list string }

var lyrics = []verse{
	{`I know an old lady who swallowed a fly.`,
		`I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.`,
		`She swallowed the spider to catch the fly.`},

	{`I know an old lady who swallowed a bird.
How absurd to swallow a bird!`,
		`She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.`},

	{`I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!`,
		`She swallowed the cat to catch the bird.`},

	{`I know an old lady who swallowed a dog.
What a hog, to swallow a dog!`,
		`She swallowed the dog to catch the cat.`},

	{`I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!`,
		`She swallowed the goat to catch the dog.`},

	{`I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!`,
		`She swallowed the cow to catch the goat.`},

	{`I know an old lady who swallowed a horse.
She's dead, of course!`, ``},
}

// returns 1-based verse of song
func Verse(v int) string {
	// if it's the final verse, return it alone
	if v == len(lyrics) {
		return lyrics[v-1].item
	}

	// otherwise return the item plus the list plus the list of all previous stanzas
	r := lyrics[v-1].item
	for i := v - 1; i >= 0; i-- {
		r += "\n" + lyrics[i].list
	}
	return r
}

// returns 1-based verses of song
func Verses(i, j int) string {
	r := ""
	for x := i; x <= j; x++ {
		if r != "" {
			r += "\n\n"
		}
		r += Verse(x)
	}
	return r
}

// returns whole song
func Song() string {
	return Verses(1, 8)
}
