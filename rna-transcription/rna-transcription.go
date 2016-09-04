// Package strand provides a function to transcribe RNA.
package strand

const testVersion = 3

func complement(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return 0 // consistent with behaviour of map version
}

// ToRNA returns a transcription of the supplied RNA string.
func ToRNA(in string) string {
	var ret = make([]rune, len(in))
	for i, r := range in {
		ret[i] = complement(r)
	}
	return string(ret)
}
