package pangram

import (
	"strings"
	"unicode"
)

//IsPangram determines if a sentence is a pangram.
//A pangram is a sentence using every letter of the alphabet at least once.

const englishAlphabetLen = 26
func IsPangram(input string) bool {
	appeared := map[rune]bool{}
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r){
			continue
		}
		if !appeared[r] {
			appeared[r] = true
		}
	}

	return len(appeared) == englishAlphabetLen
}
