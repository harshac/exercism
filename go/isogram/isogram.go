package isogram

import "strings"

func IsIsogram(input string) bool {
	runes := []rune(strings.ToLower(input))
	for i, r := range runes {
		if exists(r, runes[i+1:]) {
			return false
		}
	}
	return true
}

func exists(ir rune, runes []rune) bool {
	for _, r := range runes {
		if r == ir && !(r == '-' || r == ' ') {
			return true
		}
	}
	return false
}
