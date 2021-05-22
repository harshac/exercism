package hamming

import "errors"

//Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("lengths of polynucleotides do not match")
	}
	var dist int
	for i, r := range ar {
		if r != br[i] {
			dist++
		}
	}
	return dist, nil
}
