package isbn

import "strings"

//IsValidISBN validates book identification numbers.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var sum int
	for i, d := range isbn {
		sum += (10 - i) * digit(d, i)
	}
	return sum%11 == 0
}

func digit(d rune, index int) int {
	if d == 'X' && index == 9 {
		return 10
	}
	return int(d - '0')
}
