package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	onlyLettersRegex := regexp.MustCompile("[^a-zA-Z']+")
	processedString := onlyLettersRegex.ReplaceAllString(s, " ")

	space := regexp.MustCompile(`\s+`)
	wws := space.ReplaceAllString(processedString, " ")

	split := strings.Split(wws, " ")
	var abbr string
	for _, word := range split {
		abbr += word[0:1]
	}
	return strings.ToUpper(abbr)
}
