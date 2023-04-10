package easyregex

import (
	"regexp"
	"strings"
)

type RegexLinear []RegexNode

type RegexNode interface {
	Gen() (string, error)
}

func New() RegexLinear {
	return RegexLinear{}
}

func (r RegexLinear) Compile() (*regexp.Regexp, error) {
	strs := make([]string, len(r))
	for i := range strs {
		var err error
		strs[i], err = r[i].Gen()
		if err != nil {
			return nil, err
		}
	}
	str := strings.Join(strs, "")
	rex, err := regexp.Compile(str)
	if err != nil {
		return nil, err
	}
	return rex, nil
}
func (r RegexLinear) MustCompile() *regexp.Regexp {
	ret, err := r.Compile()
	if err != nil {
		panic(err)
	}
	return ret
}
