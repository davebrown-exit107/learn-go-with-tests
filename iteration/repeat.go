package iteration

// Repeat a given character five times
func Repeat(char string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return repeated
}
