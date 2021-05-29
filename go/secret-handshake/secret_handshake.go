package secret

import (
	"math"
)

var code = map[uint]string{
	1:     "wink",
	10:    "double blink",
	100:   "close your eyes",
	1000:  "jump",
	10000: reverseCommand,
}

const reverseCommand = "reverse"

func Handshake(num uint) []string {
	var codes []string
	binary := toReversedBinary(num)
	for i, u := range binary {
		if u == 1 {
			code := code[toBinaryUint(u, i)]
			if code == reverseCommand {
				codes = reverse(codes)
				continue
			}
			codes = append(codes, code)
		}
	}
	return codes
}

func toReversedBinary(num uint) []uint {
	var output []uint
	for num > 0 {
		output = append(output, num%2)
		num = num / 2
	}
	return output
}

func toBinaryUint(digit uint, position int) uint {
	return digit * uint(math.Pow(float64(10), float64(position)))
}

func reverse(codes []string) []string {
	len := len(codes)
	var reversed = make([]string, len)
	for i, code := range codes {
		reversed[len-1-i] = code
	}
	return reversed
}
