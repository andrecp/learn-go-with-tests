package iteration

// Repeat repeats a character, duh!
func Repeat(character string, numberOfTimes int) string {
	var repeated string
	for i := 0; i < numberOfTimes; i++ {
		repeated += character
	}
	return repeated
}
