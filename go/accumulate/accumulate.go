package accumulate

//Accumulate applies given operation and collects results in slice of string
func Accumulate(words []string, operation func(string) string) []string {
	var output []string
	for _, word := range words {
		output = append(output, operation(word))
	}
	return output
}
