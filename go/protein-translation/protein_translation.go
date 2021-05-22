package protein

import (
	"errors"
)

const (
	Stop          = "STOP"
	Methionine    = "Methionine"
	Phenylalanine = "Phenylalanine"
	Leucine       = "Leucine"
	Serine        = "Serine"
	Tyrosine      = "Tyrosine"
	Cysteine      = "Cysteine"
	Tryptophan    = "Tryptophan"
)

var codonToProtein = map[string]string{
	"AUG": Methionine,
	"UUU": Phenylalanine,
	"UUC": Phenylalanine,
	"UUA": Leucine,
	"UUG": Leucine,
	"UCU": Serine,
	"UCC": Serine,
	"UCA": Serine,
	"UCG": Serine,
	"UAU": Tyrosine,
	"UAC": Tyrosine,
	"UGC": Cysteine,
	"UGU": Cysteine,
	"UGG": Tryptophan,
	"UAA": Stop,
	"UAG": Stop,
	"UGA": Stop,
}

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid codon")
)

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

func FromCodon(c string) (string, error) {
	protein, ok := codonToProtein[c]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == Stop {
		return "", ErrStop
	}
	return protein, nil
}
