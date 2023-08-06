package iteration

const numRepeat = 5

// Repeat a given string five times
func Repeat(char string) string {
	var repeated string
	for i := 0; i < numRepeat; i++ {
		repeated = repeated + char
	}
	return repeated
}
