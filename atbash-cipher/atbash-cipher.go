package atbash

import "bytes"

// Atbash encrypts a string using an ancient cypher.
func Atbash(s string) string {
	c1 := make(chan rune)
	c2 := make(chan rune)

	// pipeline using close / range control pattern
	go filter(c1, s)
	go format(c1, c2)

	var bb bytes.Buffer
	for r := range c2 {
		bb.WriteRune(r)
	}
	return bb.String()
}

func filter(chars chan rune, s string) {
	for _, r := range s { //} strings.ToLower(s) {
		switch {
		case r >= 'A' && r <= 'Z':
			chars <- encode(r + 'a' - 'A')
		case r >= 'a' && r <= 'z':
			chars <- encode(r)
		case r >= '0' && r <= '9':
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
