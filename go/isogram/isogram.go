package isogram

import "strings"

//IsIsogram determines if given string is an isogram
//It ignores strings with multiple spaces and hyphens
func IsIsogram(input string) bool {
	var repeated = map[rune]bool{}
	for _, r := range strings.ToLower(input) {
		if r != ' ' && r != '-' {
			continue
		}
		if repeated[r] {
			return false
		}
		repeated[r] = true
	}
	return true
}
