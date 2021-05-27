package summultiples

//SumMultiples calculates the sum of all multiples of divisors within the limit
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor > 0 && i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
