package eval

type tokenKind int

const (
	number tokenKind = 1
	leftParen
	rightParen
)

type token struct {
	kind tokenKind
}

type expScanner struct {
	source  string
	start   int
	current int
	tokens  []token
}
