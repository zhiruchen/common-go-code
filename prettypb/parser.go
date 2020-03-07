package scanner

import (
	"bytes"
	"fmt"
)
const blankSpace = "    "

type JSONObject map[string]interface{}

func (obj JSONObject) Represent() string {
	if obj == nil {
		return "{}"
	}

	var count = len(obj)
	buf := bytes.NewBufferString("{")
	for k, v := range obj {
		buf.WriteString("\n" + blankSpace)
		buf.WriteString(k + ": ")

		switch v.(type) {
		case string:
			buf.WriteString(v.(string))
		case float64, bool:
			buf.WriteString(fmt.Sprintf("%v", v))
		case nil:
			buf.WriteString("null")
		case JSONObject:
			buf.WriteString("\n")
			vv := v.(JSONObject)
			buf.WriteString(vv.Represent())
		case JSONArray:
			vv := v.(JSONArray)
			buf.WriteString(vv.Represent())
		}

		if count > 1 {
			buf.WriteString(",")
		}
		count--
	}

	buf.WriteString("\n" + blankSpace + "}")
	return buf.String()
}

type Parser struct {
	tokens  []*Token
	current int
	errFunc func(msg ...string)
}

type pair struct {
	key string
	val interface{}
}

type JSONArray []interface{}

func (array JSONArray) Represent() string {
	if array == nil {
		return "[]"
	}

	var count = len(array)
	buf := bytes.NewBufferString("[")

	for _, v := range array {
		buf.WriteString("\n" + blankSpace)

		switch v.(type) {
		case string, float64, bool:
			buf.WriteString(fmt.Sprintf("%v", v))
		case nil:
			buf.WriteString("null")
		case JSONObject:
			buf.WriteString("\n")
			buf.WriteString(blankSpace)
			vv := v.(JSONObject)
			buf.WriteString(vv.Represent())
		case JSONArray:
			vv := v.(JSONArray)
			buf.WriteString(vv.Represent())
		}

		if count > 1 {
			buf.WriteString(",")
		}

		count--
	}

	buf.WriteString("\n" + "]")
	return fmt.Sprintf(buf.String())
}

func NewParser(tokens []*Token, errFunc func(msg ...string)) *Parser {
	return &Parser{tokens: tokens, errFunc: errFunc}
}

func (p *Parser) Parse() Object {

	if p.check(LeftBrace) {
		return p.object()
	}

	if p.check(OpenBracket) {
		return p.array()
	}

	p.errFunc(p.peek().ToString(), "unexpected token")
	return nil
}

func (p *Parser) object() JSONObject {
	p.consume(LeftBrace, "expect `{`")
	obj := make(JSONObject)

	if !p.check(RightBrace) {
		p.members(obj)
	}

	p.consume(RightBrace, fmt.Sprintf("expect `}` after %v", p.peek().Lexeme))

	return obj
}

func (p *Parser) members(obj JSONObject) {
	pair := p.pair()
	obj[pair.key] = pair.val

	for p.check(Comma) {
		p.consume(Comma, fmt.Sprintf("expect `,` after %v", pair.val))

		pair := p.pair()
		obj[pair.key] = pair.val
	}
}

func (p *Parser) pair() *pair {
	key := p.consume(String, "expect key")
	p.consume(Colon, fmt.Sprintf("expect `:` after key: %+v", key))

	pr := &pair{key: key.Literal.(string)}
	pr.val = p.getValue()

	return pr
}

func (p *Parser) getValue() interface{} {
	if p.check(String) {
		val := p.consume(String, "expect string")
		return val.Literal.(string)
	}

	if p.check(Number) {
		val := p.consume(Number, "expect number")
		return val.Literal.(float64)
	}

	if p.check(True) {
		val := p.consume(True, "expect true")
		return val.Literal.(bool)
	}

	if p.check(False) {
		val := p.consume(False, "expect false")
		return val.Literal.(bool)
	}

	if p.check(Null) {
		p.consume(Null, "expect null")
		return nil
	}

	if p.check(LeftBrace) {
		return p.object()
	}

	if p.check(OpenBracket) {
		return p.array()
	}

	p.errFunc(p.peek().ToString(), "unsupported value type")
	return nil
}

func (p *Parser) array() JSONArray {
	p.consume(OpenBracket, "expect `[`")
	array := p.elements()
	p.consume(CloseBracket, "expect `]`")

	return array
}

func (p *Parser) elements() JSONArray {
	val := p.getValue()
	array := JSONArray{val}

	for p.check(Comma) {
		p.consume(Comma, fmt.Sprintf("expect `,` after %v", val))

		array = append(array, p.getValue())
	}

	return array
}

func (p *Parser) consume(t Type, msg string) *Token {
	if p.check(t) {
		return p.advance()
	}

	p.errFunc(msg)
	return nil
}

func (p *Parser) match(tokenTypes ...Type) bool {
	for _, t := range tokenTypes {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == t
}

func (p *Parser) peek() *Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() *Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == Eof
}

func (p *Parser) previous() *Token {
	return p.tokens[p.current-1]
}
