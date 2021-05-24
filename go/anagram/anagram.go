package anagram

import (
	"sort"
	"strings"
)

//Detect identifies all anagrams of the subject
func Detect(subject string, candidates []string) []string {
	var anagrams []string

	sortedSubject := sortedString(strings.ToLower(subject))
	for _, c := range candidates {
		if len(c) != len(subject) || isWordItself(subject, c) {
			continue
		}
		if sortedString(strings.ToLower(c)) == sortedSubject {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func isWordItself(word, candidate string) bool {
	return strings.ToLower(word) == strings.ToLower(candidate)
}

func sortedString(subject string) string {
	subjectArr := strings.Split(subject, "")
	sort.Strings(subjectArr)
	return strings.Join(subjectArr, "")
}
