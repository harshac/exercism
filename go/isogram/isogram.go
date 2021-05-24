package isogram

import "strings"

//IsIsogram determines if given string is an isogram
//It ignores strings with multiple spaces and hyphens
func IsIsogram(input string) bool {
	var runeMap = make(map[rune]bool)
	for _, r := range []rune(strings.ToLower(input)) {
		if _, ok := runeMap[r]; ok && r != ' ' && r != '-' {
			return false
		}
		runeMap[r] = true
	}
	return true
}
