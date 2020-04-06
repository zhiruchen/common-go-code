package eval

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/emirpasic/gods/stacks/arraystack"
)

var (
	// the '(' in stack has lowest precedence
	operatorPrecedence = map[string]int{
		")": 3,
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
		"(": 0,
	}

	errInvalidExpression = fmt.Errorf("invalid expression")
)

// Eval eval a expression
// like 1 + 2, (8 * 10) / 6
func Eval(exp string) (float64, error) {
	exp = strings.ReplaceAll(exp, " ", "")
	result := arraystack.New()
	operatorStack := arraystack.New()

	start := 0
	length := len(exp)

	for start < length {
		c := rune(exp[start])
		if unicode.IsDigit(c) {
			newStart, num, err := parseFloat(exp, start)
			if err != nil {
				return 0, err
			}
			start = newStart
			result.Push(num)
			continue
		}

		if c == '(' {
			start++
			operatorStack.Push("(")
			continue
		}

		if c == ')' {
			for {
				operator, ok := operatorStack.Pop()
				if !ok {
					return 0, fmt.Errorf("pop operator error")
				}
				if operator.(string) == "(" {
					break
				}
				v, err := calculate(operator.(string), result)
				if err != nil {
					return 0, err
				}

				result.Push(v)
			}

			start++
			continue
		}

		p, ok := operatorPrecedence[string(c)]
		if !ok {
			return 0, fmt.Errorf("unsupport operator: %c", c)
		}

		currentOp, ok := operatorStack.Peek()
		if !ok {
			start++
			operatorStack.Push(string(c))
			continue
		}

		currentPrec := operatorPrecedence[currentOp.(string)]
		if currentPrec >= p {
			v, err := calculate(currentOp.(string), result)
			if err != nil {
				return 0, err
			}
			operatorStack.Pop()
			result.Push(v)
		}

		start++
		operatorStack.Push(string(c))
	}

	for {
		op, ok := operatorStack.Pop()
		if !ok {
			break
		}

		v, err := calculate(op.(string), result)
		if err != nil {
			return 0, err
		}
		result.Push(v)
	}

	v, ok := result.Pop()
	if !ok {
		return 0, errInvalidExpression
	}

	return v.(float64), nil
}

func parseFloat(exp string, start int) (int, float64, error) {
	var numbers []rune
	for start < len(exp) {
		c := rune(exp[start])
		if !unicode.IsDigit(c) {
			break
		}

		start++
		numbers = append(numbers, c)
	}

	if start >= len(exp) || rune(exp[start]) != '.' {
		n, _ := strconv.ParseFloat(string(numbers), 10)
		return start, n, nil
	}

	c := rune(exp[start])
	if start+1 == len(exp) {
		return 0, 0, fmt.Errorf("%s invalid digit number", string(append(numbers, c)))
	}

	start++
	nextC := rune(exp[start])
	if !unicode.IsDigit(nextC) {
		return 0, 0, fmt.Errorf("%s invalid digit number", string(append(numbers, c, nextC)))
	}
	numbers = append(numbers, c, nextC)

	start++
	for ; start < len(exp) && unicode.IsDigit(rune(exp[start])); start++ {
		c := rune(exp[start])
		numbers = append(numbers, c)
	}

	n, _ := strconv.ParseFloat(string(numbers), 10)
	return start, n, nil
}

func calculate(op string, result *arraystack.Stack) (float64, error) {
	var n1, n2 interface{}
	var ok bool
	n2, ok = result.Pop()
	if !ok {
		return 0, errInvalidExpression
	}

	n1, ok = result.Pop()
	if !ok {
		return 0, errInvalidExpression
	}

	num1, num2 := n1.(float64), n2.(float64)
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("dividor is zero")
		}
		return num1 / num2, nil
	}

	return 0, fmt.Errorf("unsupport operator")
}
