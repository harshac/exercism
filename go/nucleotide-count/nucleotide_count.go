package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. 
type DNA string

//Counts the number of each nucleotide in given DNA
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, n := range []rune(d) {
		count, ok := h[n]
    if !ok {
			return nil, errors.New("invalid Nucleotide")
		}
		h[n] = count + 1
	}
	return h, nil
}
