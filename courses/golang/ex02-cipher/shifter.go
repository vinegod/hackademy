package cipher

type Shift struct {
	step int
}

func (c Shift) Encode(s string) string {
	str := []byte(PrepapeString(s))

	for i, char := range str {
		str[i] = char + byte(c.step)

		if str[i] > 'z' {
			str[i] -= 26
		} else if str[i] < 'a' {
			str[i] += 26
		}
	}
	return string(str)
}

func (c Shift) Decode(s string) string {

	str := []byte(s)

	for i, char := range str {
		str[i] = char - byte(c.step)
		if str[i] > 'z' {
			str[i] -= 26
		} else if str[i] < 'a' {
			str[i] += 26
		}
	}
	return string(str)
}
