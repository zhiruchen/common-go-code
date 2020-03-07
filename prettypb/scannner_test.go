package scanner

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestRegexReplace(t *testing.T) {
	text := `<*>(0xc115903ec0){}`

	// 0xc115903ec0
	// 0xc023982b80
	// 0xc0e983b620
	re := regexp.MustCompile(`<\*>\(0[xX][0-9a-fA-F]+\)`)
	results := re.FindAllString(text, -1)
	fmt.Println(results)

	newText := re.ReplaceAllLiteralString(text, "")
	fmt.Println(newText)

	re2 := regexp.MustCompile(`<nil>`)
	newText1 := re2.ReplaceAllLiteralString(newText, "null")
	fmt.Println(newText1)

	re3 := regexp.MustCompile(`\w+:`)
	vars := re3.FindAllString(newText1, -1)
	fmt.Println(vars)

	newText2 := re3.ReplaceAllStringFunc(newText1, func(s string) string {
		return `"` + s[:len(s)-1] + `":`
	})

	newText2 = strings.Replace(newText2, `\"`, `"`, -1)
	fmt.Println(newText2)
}

func TestNewScanner(t *testing.T) {
	text := `<*>(0xc115903ec0){}`
	fmt.Println(text)

	newText := PreProcessPbMessage(text)
	fmt.Println(newText)

	scanner := NewScanner(newText, func(line int, msg string) {
		fmt.Println("line: ", line, "message: ", msg)
		panic(msg)
	})

	tokens := scanner.ScanTokens()
	for _, t:= range tokens {
		fmt.Println(t.ToString())
	}
}

func TestRune(t *testing.T) {
	rs := []rune("-2.922127")
	for _, r := range rs {
		fmt.Println("r:", r)
	}
}