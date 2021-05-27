package prime

// Nth returns the nth prime number
func Nth(nth int) (int, bool) {
	switch nth {
	case 0:
		return 0, false
	case 1:
		return 2, true
	}
	oddNum := 1
	for nth-1 != 0 {
		oddNum = next(oddNum)
		if isPrime(oddNum) {
			nth--
		}
	}
	return oddNum, true
}

func next(num int) int {
	return num + 2
}

func isPrime(num int) bool {
	for i := 3; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
