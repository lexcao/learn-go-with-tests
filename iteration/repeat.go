package iteration

func Repeat(char string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
