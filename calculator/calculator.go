package calculator

import (
	stackR "calc-struct/stack"
	"calc-struct/utils"
	"strconv"
)

func Calculate(str string) (uint8, error) {

	stack := ""
	var result uint8

	for i := 0; i < len(str); i++{
		c := str[i]

		if utils.Is_digit(c) {
			stackR.Push(string(c), stack)
			continue
		}

		var a int
		var b int

		switch c {
		case '+':
			a, b, stack = getOperand(stack)
			stackR.Push(string(a + b), stack)
		case '-':
			a, b, stack = getOperand(stack)
			stackR.Push(string(a - b), stack)
		case '*':
			a, b, stack = getOperand(stack)
			stackR.Push(string(a * b), stack)
		case '/':
			a, b, stack = getOperand(stack)
			stackR.Push(string(a / b), stack)
		}
	}

	if(stackR.Count(stack) > 1) {
		panic("Input did not fully simplify")
	}

	result, stack = stackR.Pop(stack)
	return result, nil
}

func getOperand(stack string) (int, int, string) {
	var a uint8
	var b uint8
	b, stack = stackR.Pop(stack)
	a, stack = stackR.Pop(stack)
	var intb, _ = strconv.Atoi(string(b))
	var inta, _ = strconv.Atoi(string(a))

	return inta, intb, stack
}