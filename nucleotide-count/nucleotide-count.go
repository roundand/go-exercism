// Package dna provides functions for counting nucleotides in DNA.
package dna

import "fmt"

const testVersion = 2

// DNA is composed of nucleotides.
type DNA string

// Histogram is a frequency map of nucleotides
type Histogram map[byte]int

// NewHistogram returns a new Histogram
func NewHistogram() Histogram {
	return map[byte]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
}

// Counts returns a histogram of nucleotide frequencies
func (dna DNA) Counts() (Histogram, error) {
	fm := NewHistogram()
	for i := 0; i < len(dna); i++ {
		if !isNucleotide(dna[i]) {
			return nil, fmt.Errorf("dna: strand %q contains non-nucleotide %q", dna, dna[i])
		}
		fm[dna[i]]++
	}
	return fm, nil
}

// Count returns frequency of a nucleotide in a DNA strand
func (dna DNA) Count(b byte) (int, error) {
	freq := 0
	for i := 0; i < len(dna); i++ {
		if !isNucleotide(b) {
			return 0, fmt.Errorf("dna: strand %q contains non-nucleotide %q", dna, dna[i])
		}
		if dna[i] == b {
			freq++
		}
	}
	return freq, nil
}

func isNucleotide(b byte) bool {
	switch b {
	case 'G', 'T', 'A', 'C':
		return true
	default:
		return false
	}
}
