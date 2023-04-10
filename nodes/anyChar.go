package nodes

type AnyChar struct{}

func NewAnyCharNode(start, end rune) *AnyChar {
	return &AnyChar{}
}
func (a *AnyChar) Gen() (string, error) {
	return ".", nil
}
