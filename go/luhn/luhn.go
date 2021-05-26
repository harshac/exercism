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
	numbers, err := luhn(trimmedInput)
	if err != nil {
		return false
	}

	var sum int
	for _, d := range numbers {
		sum += d
	}
	return sum%10 == 0
}

func luhn(input string) ([]int, error) {
	digits := strings.Split(input, "")
	length := len(digits)
	luhnDigits := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(digits[i])
		if err != nil {
			return nil, err
		}
		luhnDigits[i] = digit

		if (length-i)%2 == 0 {
			double := digit * 2
			luhnDigits[i] = double
			if double > 9 {
				luhnDigits[i] = double - 9
			}
		}
	}
	return luhnDigits, nil
}
