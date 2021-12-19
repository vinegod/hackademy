package cipher

import "strings"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func NewCaesar() Cipher {
	return Shift{step: 3}
}

func NewShift(i int) Cipher {
	if i >= 26 || i <= -26 || i == 0 {
		return nil
	}
	return Shift{step: i}
}

func NewVigenere(s string) Cipher {
	if strings.ContainsAny(s, " ,-1234567890"+strings.ToUpper(alphabet)) {
		return nil
	}
	if len(s) <= 1 {
		return nil
	}
	if strings.Count(s, string(s[0])) == len(s) {
		return nil
	}
	return Vigenere{Key: s}
}

func PrepapeString(s string) string {
	for _, c := range []string{" ", ",", "@", "-", ".", "!", "#", "?", "1", "2", "3"} {
		s = strings.ReplaceAll(s, c, "")
	}
	return strings.ToLower(s)
}
