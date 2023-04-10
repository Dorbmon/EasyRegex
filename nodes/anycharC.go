package nodes

import "strings"

type AnyCharC struct {
	Chars []rune
}

func NewAnyCharCNode(chars []rune) *AnyCharC {
	return &AnyCharC{
		Chars: chars,
	}
}
func (a *AnyCharC) Gen() (string, error) {
	ret := strings.Builder{}
	ret.Grow(len(a.Chars) + 2)
	ret.WriteRune('[')
	ret.WriteString(string(a.Chars))
	ret.WriteRune(']')
	return ret.String(), nil
}
