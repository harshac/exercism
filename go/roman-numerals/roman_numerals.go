package romannumerals

import "errors"

var arabicToRoman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", errors.New("invalid input")
	}
	var roman string
	for _, base := range []int{1000, 100, 10, 1} {
		quotient := num / base
		num = num % base
		switch {
		case quotient == 9:
			roman += arabicToRoman[base]
			roman += arabicToRoman[base*10]
		case quotient >= 5:
			roman += arabicToRoman[base*5]
			roman += appendMultiple(quotient-5, base)
		case quotient == 4:
			roman += arabicToRoman[base]
			roman += arabicToRoman[base*5]
		default:
			roman += appendMultiple(quotient, base)
		}
	}
	return roman, nil
}

func appendMultiple(times int, key int) string {
	var roman string
	for i := 0; i < times; i++ {
		roman += arabicToRoman[key]
	}
	return roman
}
