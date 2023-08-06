package iteration

import "fmt"

// Repeat a given character five times
func Repeat(char string) string {
	return fmt.Sprintf("%s%s%s%s%s", char, char, char, char, char)
}
