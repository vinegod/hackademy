package brackets

import (
	"strings"

	stack "github.com/golang-collections/collections/stack"
)

const (
	begin = "{[("
	end   = "}])"
)

func Bracket(s string) (bool, error) {

	if s == "" {
		return true, nil
	}

	bracket := stack.New()
	str := []byte(s)

	for _, char := range str {
		i := strings.IndexByte(begin, char)
		if i != -1 {
			bracket.Push(i)
		} else {
			if bracket.Peek() != strings.IndexByte(end, char) {
				return false, nil
			} else {
				bracket.Pop()
			}
		}
	}
	return bracket.Len() == 0, nil
}
