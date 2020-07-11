package eval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfixToSuffix(t *testing.T) {
	t.Parallel()
	cases := []struct {
		desc   string
		infix  string
		result float64
		err    error
	}{
		{
			desc:   "test (1+2)",
			infix:  "(1+2)",
			result: 3,
			err:    nil,
		},
		{
			desc:   "test ((1+2) * (10 * 10))",
			infix:  "((1+2) * (10 * 10))",
			result: 300,
			err:    nil,
		},
		{
			desc:   "with float numbers",
			infix:  "((1.0+2.1) * (8.88 * 10))",
			result: 275.28000000000003,
			err:    nil,
		},
		{
			desc:   "high precedence operator before lower",
			infix:  "8 * (9 / 3 + 2) + 100",
			result: 140,
			err:    nil,
		},
		{
			desc:   "high precedence operator after lower",
			infix:  "8 + (9 / 3 + 2) * 100",
			result: 508,
			err:    nil,
		},
		{
			desc:   "expression with parentheses",
			infix:  "(8 + (9 / 3 + 2) * 100)",
			result: 508,
			err:    nil,
		},
		{
			desc:   "invalid float number",
			infix:  "10.",
			result: 0,
			err:    fmt.Errorf("10. invalid digit number"),
		},
		{
			desc:   "invalid expression - no operand after operator",
			infix:  "(1+2) *",
			result: 0,
			err:    errInvalidExpression,
		},
		{
			desc:   "invalid expression - unbalanced parentheses",
			infix:  "(1+2",
			result: 0,
			err:    errInvalidExpression,
		},
	}

	for _, cs := range cases {
		t.Run(cs.desc, func(t *testing.T) {
			s, err := Eval(cs.infix)
			assert.Equal(t, cs.err, err)
			assert.Equal(t, cs.result, s)
		})
	}
}
