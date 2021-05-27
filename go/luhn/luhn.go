package luhn

import (
	"strconv"
	"strings"
)

//Valid determines whether a number is valid per the Luhn formula
func Valid(input string) bool {
	trimmedInput := strings.ReplaceAll(input, " ", "")
	if len(trimmedInput) <= 1 {
		return false
	}
	digits := strings.Split(trimmedInput, "")
	length := len(digits)
	var applyLuhn bool
	if length%2 == 0 {
		applyLuhn = true
	}
	var sum int
	for _, d := range digits {
		digit, err := strconv.Atoi(d)
		if err != nil {
			return false
		}
		toBeAdded := digit

		if applyLuhn {
			double := digit * 2
			toBeAdded = double
			if double > 9 {
				toBeAdded = double - 9
			}
			applyLuhn = false
		} else {
			applyLuhn = true
		}
		sum += toBeAdded
	}
	return sum%10 == 0
}
