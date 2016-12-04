// Package twelve provides functions to generate the Twleve Days of Christmas song.
package twelve

import "fmt"

const testVersion = 1

var (
	content = verses{
		{"first", "a Partridge in a Pear Tree"},
		{"second", "two Turtle Doves"},
		{"third", "three French Hens"},
		{"fourth", "four Calling Birds"},
		{"fifth", "five Gold Rings"},
		{"sixth", "six Geese-a-Laying"},
		{"seventh", "seven Swans-a-Swimming"},
		{"eighth", "eight Maids-a-Milking"},
		{"ninth", "nine Ladies Dancing"},
		{"tenth", "ten Lords-a-Leaping"},
		{"eleventh", "eleven Pipers Piping"},
		{"twelfth", "twelve Drummers Drumming"},
	}
)

type verse struct {
	day  string
	gift string
}

type verses []verse

// gifts returns 0-based list of gifts
func (v verses) gifts(n int) string {
	switch n {
	case 0:
		return v[0].gift
	case 1:
		return fmt.Sprintf("%s, and %s", v[1].gift, v[0].gift)
	default:
		return fmt.Sprintf("%s, %s", v[n].gift, v.gifts(n-1))
	}
}

// Verse returns the 1-based verse of the Twelve Days of Christmas.
func Verse(n int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", content[n-1].day, content.gifts(n-1))
}

// Song returns all twelve verses of the Twelve Days of Christmas.
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		fmt.Printf("Song(), i: %v\n", i)
		song += Verse(i) + "\n"
	}
	return song
}
