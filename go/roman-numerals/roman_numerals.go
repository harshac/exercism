package romannumerals

import (
	"errors"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var arabicToRoman = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("invalid input")
	}
	var roman string
	return eval(number, roman), nil
}

func eval(number int, roman string) string {
	if number == 0 {
		return roman
	}
	if number/1000 > 0 {
		roman = roman + "M"
		return eval(number-1000, roman)
	}
	if number/500 > 0 {
		roman = roman + "D"
		return eval(number-500, roman)
	}
	if number/100 > 0 {
		roman = roman + "C"
		return eval(number-100, roman)
	}
	if number/50 > 0 {
		roman = roman + "L"
		return eval(number-50, roman)
	}
	if number/10 > 0 {
		roman = roman + "X"
		return eval(number-10, roman)
	}
	if number/5 > 0 {
		roman = roman + "V"
		return eval(number-5, roman)
	}
	roman = roman + "I"
	if strings.Contains(roman, "IIII"){
		roman = strings.Replace(roman, "IIII", "IV", 1)
	}
	return eval(number-1, roman)
}
