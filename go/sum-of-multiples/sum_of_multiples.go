package summultiples

//SumMultiples calculates the sum of all multiples of divisors within the limit
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for k, _ := range multiples(limit, divisors...) {
		sum += k
	}
	return sum
}

func multiples(limit int, divisors ...int) map[int]bool {
	var multiples = map[int]bool{}
	for _, divisor := range divisors {
		multiple := divisor
		multiplier := 2
		for multiple > 0 && multiple < limit {
			multiples[multiple] = true
			multiple = divisor * multiplier
			multiplier++
		}
	}
	return multiples
}
