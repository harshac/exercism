package protein

import (
	"errors"
)

const (
	stop          = "STOP"
	methionine    = "Methionine"
	phenylalanine = "Phenylalanine"
	leucine       = "Leucine"
	serine        = "Serine"
	tyrosine      = "Tyrosine"
	cysteine      = "Cysteine"
	tryptophan    = "Tryptophan"
)

var codonToProtein = map[string]string{
	"AUG": methionine,
	"UUU": phenylalanine,
	"UUC": phenylalanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UGC": cysteine,
	"UGU": cysteine,
	"UGG": tryptophan,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

var (
	//ErrStop is the error when stop sequence is found
  ErrStop        = errors.New("stop")
	//ErrInvalidBase is the error an invalid sequence is found
  ErrInvalidBase = errors.New("invalid codon")
)

//FromRNA gives the list of proteins found in RNA
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	runes := []rune(rna)
	for i := 0; i < len(runes); i = i + 3 {
		codon := string(runes[i : i+3])
		protein, err := FromCodon(codon)
		if err != nil && err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

//FromCodon returns the name of protein from the codon
func FromCodon(c string) (string, error) {
	protein, ok := codonToProtein[c]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == stop {
		return "", ErrStop
	}
	return protein, nil
}
