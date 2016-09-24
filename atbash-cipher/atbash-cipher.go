package atbash

import "strings"

// Atbash encrypts a string using an ancient cypher.
func Atbash(s string) string {
	c1 := make(chan rune)
	c2 := make(chan rune)

	go filter(c1, s)
	go format(c1, c2)

	msg := ""
	for r := range c2 {
		msg = msg + string(r)
	}

	return msg
}

func filter(chars chan rune, s string) {
	for _, r := range strings.ToLower(s) {
		switch {
		case strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", r):
			chars <- encode(r)
		case strings.ContainsRune("0123456789", r):
			chars <- r
		}
	}
	close(chars)
}

func format(c1, c2 chan rune) {
	i := 0
	for r := range c1 {
		if i == 5 {
			c2 <- ' '
			i = 0
		}
		c2 <- r
		i++
	}
	close(c2)
}

func encode(r rune) rune {
	return 'a' + 'z' - r
}
