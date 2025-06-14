package iteration

import "strings"

func Repeat(value string, count int) string {
	var rValue strings.Builder
	if count <= 0 {
		return ""
	}
	for i := 0; i < count; i++ {
		rValue.WriteString(value)
	}
	return rValue.String()
}
