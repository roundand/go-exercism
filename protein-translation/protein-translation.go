// Package protein provides functions for protein translation.
package protein

const testVersion = 1

var codon = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon takes a codon and returns the name of the protein that will be generated.
func FromCodon(s string) string {
	return codon[s]
}

// FromRNA takes a string of codons and returns an array of proteins.
func FromRNA(s string) []string {
	r := []string{}
	for i := 0; len(s) >= i+3; i += 3 {
		p := FromCodon(s[i : i+3])
		if p == "STOP" {
			return r
		}
		r = append(r, p)
	}
	return r
}
