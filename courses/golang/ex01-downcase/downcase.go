package downcase

func Downcase(s string) (string, error) {
	str := []byte(s)
	for i, c := range str {
		if c >= 'A' && c <= 'Z' {
			str[i] = byte(c + 32)
		}
	}
	return string(str), nil
}
