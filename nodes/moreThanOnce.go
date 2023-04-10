package nodes

type MoreThanOnce struct {
	Char rune
}

func NewMoreThanOnceNode(char rune) *MoreThanOnce {
	return &MoreThanOnce{
		Char: char,
	}
}
func (a *MoreThanOnce) Gen() (string, error) {
	return string(a.Char) + "+", nil
}
