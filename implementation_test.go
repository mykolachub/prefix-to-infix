package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleOne(t *testing.T) {
	res, err := PrefixToInfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(5+((4-2)*3))", res)
	}
}

func TestSimpleTwo(t *testing.T) {
	res, err := PrefixToInfix("^ + 6 5 + 1 1")
	if assert.Nil(t, err) {
		assert.Equal(t, "((6+5)^(1+1))", res)
	}
}

func TestSimpleThree(t *testing.T) {
	res, err := PrefixToInfix("- 9 + 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "(9-(5+8))", res)
	}
}

func TestSimpleFour(t *testing.T) {
	res, err := PrefixToInfix("/ 6 ^ 4 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(6/(4^2))", res)
	}
}

func TestAdvancedOne(t *testing.T) {
	res, err := PrefixToInfix("- 1 * + 5 * - 4 2 3 / 6 ^ 4 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(1-((5+((4-2)*3))*(6/(4^2))))", res)
	}
}

func TestAdvancedTwo(t *testing.T) {
	res, err := PrefixToInfix("+ - 4 2 - - 9 + 5 8 ^ + 6 5 + 1 1")
	if assert.Nil(t, err) {
		assert.Equal(t, "((4-2)+((9-(5+8))-((6+5)^(1+1))))", res)
	}
}

func TestAdvancedThree(t *testing.T) {
	res, err := PrefixToInfix("+ 5 - 3 ^ 6 + 4 - 5 + 6 - 4 ^ 2 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(5+(3-(6^(4+(5-(6+(4-(2^2))))))))", res)
	}
}

func TestAdvancedFour(t *testing.T) {
	res, err := PrefixToInfix("+ 8 / - - 5 2 3 ^ * / 2 3 6 7")
	if assert.Nil(t, err) {
		assert.Equal(t, "(8+(((5-2)-3)/(((2/3)*6)^7)))", res)
	}
}
