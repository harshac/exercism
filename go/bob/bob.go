package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)
	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	if isUpper(trimmedRemark) && isQuestion(trimmedRemark) {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion(trimmedRemark) {
		return "Sure."
	}

	if isUpper(trimmedRemark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isUpper(s string) bool {
	if !containsLetter(s) {
		return false
	}
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func containsLetter(s string) bool {
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
