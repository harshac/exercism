package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode encodes the string with data compression
func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}
	var encoded string
	lastChar := []rune(input)[0]
	count := 0
	for _, i := range input {
		if i == lastChar {
			count++
			continue
		}
		encoded += encode(count, lastChar)
		lastChar = i
		count = 1
	}
	encoded += encode(count, lastChar)
	return encoded
}

//RunLengthDecode decodes the string with RLE
func RunLengthDecode(encoded string) string {
	var countstr string
	var decoded string
	for _, c := range encoded {
		if unicode.IsNumber(c) {
			countstr += string(c)
			continue
		}
		decoded += strings.Repeat(string(c), toInt(countstr))
		countstr = ""
	}
	return decoded
}

func encode(count int, char rune) string {
	if count == 1 {
		return fmt.Sprintf("%c", char)
	}
	return fmt.Sprintf("%d%c", count, char)
}

func toInt(str string) int {
	if str == "" {
		return 1
	}
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return atoi
}
