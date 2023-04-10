package easyregex

import "github.com/Dorbmon/EasyRegex/nodes"

func (r RegexLinear) AnyCharC(runes ...rune) RegexLinear {
	r = append(r, nodes.NewAnyCharCNode(runes))
	return r
}
