package nodes

import "strings"

type BetweenChar struct {
	Start rune
	End   rune
}

func NewBetweenCharNode(start, end rune) *BetweenChar {
	return &BetweenChar{
		Start: start,
		End:   end,
	}
}
func (a *BetweenChar) Gen() (string, error) {
	ret := strings.Builder{}
	ret.Grow(5)
	ret.WriteRune('[')
	ret.WriteRune(a.Start)
	ret.WriteRune('-')
	ret.WriteRune(a.End)
	ret.WriteRune(']')
	return ret.String(), nil
}
