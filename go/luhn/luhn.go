package luhn

import (
	"strconv"
	"strings"
)

//Valid determines whether a number is valid per the Luhn formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	applyLuhn := len(input)%2 == 0
	var sum int
	for _, d := range input {
		digit, err := strconv.Atoi(string(d))
		if err != nil {
			return false
		}
		if applyLuhn {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		applyLuhn = !applyLuhn
	}
	return sum%10 == 0
}
