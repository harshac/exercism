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
	for i, u := range toReversedBinary(num) {
		if u == 1 {
			code := code[key(i)]
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

func key(position int) uint {
	return uint(math.Pow(float64(10), float64(position)))
}

func reverse(codes []string) []string {
	var reversed []string
	for _, code := range codes {
		reversed = append([]string{code}, reversed...)
	}
	return reversed
}
