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
		val := int(d - '0')
		if d == 'X' && i == len(isbn)-1 {
			val = 10
		}
		sum += (10 - i) * val
	}
	return sum%11 == 0
}
