package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, n := range []rune(d) {
		if count, ok := h[n]; !ok {
			return nil, errors.New("invalid Nucleotide")
		} else {
			h[n] = count + 1
		}
	}
	return h, nil
}
