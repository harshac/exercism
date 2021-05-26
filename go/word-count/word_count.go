package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	reg := regexp.MustCompile(`[\w']+`)
	frequency := Frequency{}
	for _, word := range reg.FindAllString(strings.TrimSpace(input), -1) {
		trimWord := strings.Trim(strings.ToLower(word), "'")
		frequency[trimWord] += 1
	}
	return frequency
}
