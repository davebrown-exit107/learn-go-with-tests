package iteration

import "strings"

func Repeat(value string) string {
	var rValue strings.Builder
	count := 4
	for i := 0; i < count; i++ {
		rValue.WriteString(value)
	}
	return rValue.String()
}
