package iteration

func Repeat(char string, len int) string {
	acc := ""

	for i := 0; i < len; i++ {
		acc += char
	}

	return acc
}
