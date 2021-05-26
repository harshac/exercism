package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	frequency := Frequency{}
	reg := regexp.MustCompile(`\w+('\w)*`)
	lower := strings.ToLower(input)
	for _, word := range reg.FindAllString(lower, -1) {
		frequency[word] += 1
	}
	return frequency
}
