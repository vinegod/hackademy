package cipher

import "strings"

const (
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
	alphabetLen = 26
)

type Vigenere struct {
	Key string
}

func (v Vigenere) Encode(s string) string {

	s = PrepapeString(s)
	keyLen := len(v.Key)

	key := byteToIndex(v.Key)
	str := byteToIndex(s)

	for i, char := range str {
		str[i] = (char + key[i%keyLen]) % alphabetLen
	}

	indexToByte(str)
	return string(str)
}

func (v Vigenere) Decode(s string) string {

	s = PrepapeString(s)
	keyLen := len(v.Key)
	key := byteToIndex(v.Key)
	str := byteToIndex(s)

	for i, char := range str {
		str[i] = (alphabetLen + char - key[i%keyLen]) % alphabetLen
	}

	indexToByte(str)
	return string(str)
}

func byteToIndex(s string) []byte {
	b := []byte(s)
	for i, char := range b {
		b[i] = byte(strings.IndexByte(alphabet, char))
	}
	return b
}

func indexToByte(b []byte) {
	for i, char := range b {
		b[i] = alphabet[char]
	}
}
