package reverse

//Reverse reverses the given the string
func Reverse(input string) string {
	var reverse string
	runes := []rune(input)
	length := len(runes)
	for i := 0; i < length; i++ {
		reverse += string(runes[length-1-i])
	}
	return reverse
}
