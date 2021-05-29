package allyourbase

import (
	"errors"
	"math"
)

//ConvertToBase converts given digits from input to output base
func ConvertToBase(inputBase int, digits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if len(digits) == 0 {
		return []int{0}, nil
	}
	decimal, err := toDecimal(digits, inputBase)
	if err != nil {
		return nil, err
	}
	return fromDecimal(decimal, outputBase), nil
}

func toDecimal(digits []int, inputBase int) (int, error) {
	var decimal int
	for i, digit := range digits {
		if digit < 0 || digit >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal += digit * int(math.Pow(float64(inputBase), float64(len(digits)-1-i)))
	}
	return decimal, nil
}

func fromDecimal(num int, outputBase int) []int {
	var output []int
	if num == 0 {
		return []int{0}
	}
	for num > 0 {
		output = append([]int{num % outputBase}, output...)
		num = num / outputBase
	}
	return output
}
