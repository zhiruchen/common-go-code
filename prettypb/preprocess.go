package scanner

import (
	"fmt"
	"regexp"
	"strings"
)

// PreProcessPbMessage ...
func PreProcessPbMessage(text string) string {
	re := regexp.MustCompile(`<\*>\(0[xX][0-9a-fA-F]+\)`)
	newText := re.ReplaceAllLiteralString(text, "")
	fmt.Println(newText)

	re2 := regexp.MustCompile(`<nil>`)
	newText1 := re2.ReplaceAllLiteralString(newText, "null")

	re3 := regexp.MustCompile(`\w+:`)
	newText2 := re3.ReplaceAllStringFunc(newText1, func(s string) string {
		return `"` + s[:len(s)-1] + `":`
	})

	newText2 = strings.Replace(newText2, `\"`, `"`, -1)
	fmt.Println(newText2)

	return newText2
}
