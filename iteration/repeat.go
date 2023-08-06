package iteration

import (
	"fmt"
)

// Repeat a given string n times
func Repeat(char string, numRepeat int) (error, string) {
	var repeated string
	if numRepeat <= 0 {
		err := fmt.Errorf("Got: %q. Expected a number greater than 0 for numRepeat.", numRepeat)
		return err, ""
	}
	for i := 0; i < numRepeat; i++ {
		repeated = repeated + char
	}
	return nil, repeated
}
