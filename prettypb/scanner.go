package scanner

import (
	"fmt"
	"strconv"
	"unicode"
)

// Scanner json scanner
type Scanner struct {
	source   string
	runes    []rune
	tokens   []*Token
	start    int
	current  int
	line     int
	keywords map[string]Type

	errFunc ErrFunc
}

// ErrFunc ...
type ErrFunc func(line int, msg string)

// NewScanner a new json scanner
func NewScanner(source string, errFunc ErrFunc) *Scanner {
	return &Scanner{
		source: source,
		runes:  []rune(source),
		tokens: []*Token{},
		line:   1,
		keywords: map[string]Type{
			"false": False,
			"true":  True,
		},
		errFunc: errFunc,
	}
}

// ScanTokens 返回扫描到的token列表
func (scan *Scanner) ScanTokens() []*Token {
	for !scan.isAtEnd() {
		scan.start = scan.current
		scan.scanToken()
	}
	scan.tokens = append(
		scan.tokens,
		&Token{
			TokenType: Eof,
			Lexeme:    "",
			Literal:   nil,
			Line:      scan.line,
		},
	)
	return scan.tokens
}

func (scan *Scanner) PrintTokens() {
	for _, t := range scan.tokens {
		fmt.Printf("%s\n", t.ToString())
	}
}

func (scan *Scanner) isAtEnd() bool {
	return scan.current >= len(scan.runes)
}

func (scan *Scanner) scanToken() {
	c := scan.advance()
	switch c {
	case '{':
		scan.addToken(LeftBrace, nil)
	case '}':
		scan.addToken(RightBrace, nil)
	case '[':
		scan.addToken(OpenBracket, nil)
	case ']':
		scan.addToken(CloseBracket, nil)
	case ':':
		scan.getColonOrEmptyStr()
	case ',':
		scan.addToken(Comma, nil)
	case ' ', '\r', '\t':
	case '\n':
		scan.line++
	case '"':
		scan.getStr()
	default:
		if unicode.IsDigit(c) {
			scan.getNumber()
		// 可能是 关键字 (true / false / null), 也可能是 string
		} else if unicode.IsLetter(c) {
			scan.getIdentifierOrString()
		} else {
			scan.errFunc(scan.line, "Unexpected token!")
		}
	}
}

func (scan *Scanner) advance() rune {
	scan.current++
	return scan.runes[scan.current-1]
}

func (scan *Scanner) addToken(tokenType Type, literal interface{}) {
	text := string(scan.runes[scan.start:scan.current])
	scan.tokens = append(scan.tokens, &Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: scan.line})
}

func (scan *Scanner) match(expected rune) bool {
	if scan.isAtEnd() {
		return false
	}

	if scan.runes[scan.current] != expected {
		return false
	}

	scan.current++
	return true
}

func (scan *Scanner) peek() rune {
	if scan.current >= len(scan.runes) {
		return '\000' // https://stackoverflow.com/questions/38007361/is-there-anyway-to-create-null-terminated-string-in-go
	}
	return scan.runes[scan.current]
}

func (scan *Scanner) peekNext() rune {
	if (scan.current + 1) >= len(scan.runes) {
		return '\000'
	}
	return scan.runes[scan.current+1]
}

func (scan *Scanner) getColonOrEmptyStr() {
	var tk = Colon
	var literal interface{}

	next := scan.peek()
	if next == ' ' || next == '}'{
		tk, literal = String, ""
	}
	scan.addToken(tk, literal)
}

func (scan *Scanner) getStr() {
	for scan.peek() != '"' && !scan.isAtEnd() {
		if scan.peek() == '\n' {
			scan.line++
		}
		scan.advance()
	}
	if scan.isAtEnd() {
		scan.errFunc(scan.line, "Unterminated string")
		return
	}

	scan.advance()

	value := string(scan.runes[scan.start+1 : scan.current-1])
	scan.addToken(String, value)
}

func (scan *Scanner) getNumber() {
	for unicode.IsDigit(scan.peek()) {
		scan.advance()
	}

	if scan.peek() == '.' && unicode.IsDigit(scan.peekNext()) {
		scan.advance()

		for unicode.IsDigit(scan.peek()) {
			scan.advance()
		}
	}

	text := string(scan.runes[scan.start:scan.current])
	number, _ := strconv.ParseFloat(text, 64)
	scan.addToken(Number, number)
}

func (scan *Scanner) getIdentifierOrString() {
	for isAlphaNumeric(scan.peek()) {
		scan.advance()
	}

	text := string(scan.runes[scan.start:scan.current])
	tokenType, ok := scan.keywords[text]
	if !ok { // 有可能是 string
		scan.addToken(String, text)
		return
	}

	var literal interface{}
	switch tokenType {
	case Null:
		literal = "null"
	case True:
		literal = true
	case False:
		literal = false
	}

	scan.addToken(tokenType, literal)
}



func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
