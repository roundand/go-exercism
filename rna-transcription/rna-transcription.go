// Package strand provides a function to transcribe RNA.
package strand

const testVersion = 3

var complement = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns a transcription of the supplied RNA string.
func ToRNA(in string) string {
	var ret []rune
	for _, r := range in {
		ret = append(ret, complement[r])
	}
	return string(ret)
}
