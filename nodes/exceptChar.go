package nodes

import "strings"

type ExceptChar struct {
	Chars []rune
}

func NewExceptCharNode(chars []rune) *ExceptChar {
	return &ExceptChar{
		Chars: chars,
	}
}
func (a *ExceptChar) Gen() (string, error) {
	ret := strings.Builder{}
	ret.Grow(len(a.Chars) + 3)
	ret.WriteRune('[')
	ret.WriteRune('^')
	ret.WriteString(string(a.Chars))
	ret.WriteRune(']')
	return ret.String(), nil
}
