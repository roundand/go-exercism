// Package foodchain generates an old lady who ate a fly
package foodchain

import (
//  "strings"
//  "fmt"
)

const testVersion = 2

var vv = []struct{text, repeats string} {
	{`I know an old lady who swallowed a fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,
``},

	{`I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.`,
`She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a bird.
How absurd to swallow a bird!`,
`She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!`,
`She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a dog.
What a hog, to swallow a dog!`,
`She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!`,
`She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!`,
`She swallowed the cow to catch the goat.
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`},

	{`I know an old lady who swallowed a horse.
She's dead, of course!`,``},
}


// returns 1-based verse of song
func Verse(v int) string {
  if vv[v - 1].repeats == `` {
    return vv[v - 1].text
  }
  return vv[v - 1].text + "\n" + vv[v - 1].repeats
}

// returns 1-based verses of song
func Verses(i, j int) string {
  r := ""
  for x := i; x <= j; x++ {
    if (r != "") {
      r += "\n\n"
    }
    r += Verse(x)
  }
  return  r
}

func Song() string {
  return Verses(1, 8)
}
