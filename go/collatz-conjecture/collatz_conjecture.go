package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("invalid input")
	}
	return nextStep(input, 0), nil
}

func nextStep(input, steps int) int {
	if input == 1 {
		return steps
	}
	if input%2 == 0 {
		return nextStep(input/2, steps+1)
	}
	return nextStep(3*input+1, steps+1)
}
