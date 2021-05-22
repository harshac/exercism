//Package proverb creates a proverb based on given words
package proverb

import "fmt"

//Proverb generates a proverb based on the given words
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var proverb []string
	for i, word := range rhyme {
		if i < len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[i+1]))
		}
	}
	return append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
