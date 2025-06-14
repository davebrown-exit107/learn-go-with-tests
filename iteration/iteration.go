package iteration

func Repeat(value string) string {
	var rValue string
	count := 4
	for i := 0; i < count; i++ {
		rValue += value
	}
	return rValue
}
