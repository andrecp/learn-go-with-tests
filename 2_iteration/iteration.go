package iteration

import "strings"

// Repeat repeats a character, duh!
func Repeat(character string, numberOfTimes int) string {
	return strings.Repeat(character, numberOfTimes)
}
