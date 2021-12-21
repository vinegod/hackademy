package cipher

import "strings"

type Shift struct {
	step int
}

func (c Shift) Encode(s string) string {
	s = PrepapeString(s)

	shift := func(step rune) func(r rune) rune {
		return func(r rune) rune {
			switch {
			case r >= 'a' && r <= 'z':
				t := 'a' + (r-'a'+step)%26
				if t < 'a' {
					t += 26
				}
				return t
			}
			return r
		}
	}

	return strings.Map(shift(rune(c.step)), s)
}

func (c Shift) Decode(s string) string {

	shift := func(step rune) func(r rune) rune {
		return func(r rune) rune {
			switch {
			case r >= 'a' && r <= 'z':
				t := 'a' + (r-'a'-step)%26
				if t < 'a' {
					t += 26
				}
				return t
			}
			return r
		}
	}

	return strings.Map(shift(rune(c.step)), s)
}
