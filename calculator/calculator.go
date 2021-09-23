package calculator

import (
	"calc-struct/stack"
	"calc-struct/utils"
	"fmt"
	"strconv"
)

func Calculate(expressionRpn []string) (float64, error) {
	var stackOp []string      // Стек операндов

	for i := range expressionRpn {
		s := expressionRpn[i]
		var v float64
		var err error

		if utils.IsFloat(s) {
			stackOp = stack.Push(s, stackOp)
			continue
		}
		//Если это оператор, соответствующее количество операндов перекладывается из стека во временные переменные
		v, stackOp, err = evaluateOperator(s, stackOp)
		if err != nil {
			return 0, err
		}
		stackOp = stack.Push(fmt.Sprintf("%f", v), stackOp)
	}

	if(stack.Count(stackOp) > 1) {
		panic("Input did not fully simplify")
	}

	stRes, _ := stack.Pop(stackOp)
	result, _ := strconv.ParseFloat(stRes, 64)
	return result, nil
}

func evaluateOperator(operator string, stackOp []string) (float64, []string, error) {
	var a string
	var b string
	b, stackOp = stack.Pop(stackOp)
	a, stackOp = stack.Pop(stackOp)
	bf, err := strconv.ParseFloat(b, 64)
	af, err := strconv.ParseFloat(a, 64)

	if err != nil {
		return 0, stackOp, err
	}

	switch operator {
	case "+":
		return af + bf, stackOp, nil
	case "-":
		return af - bf, stackOp, nil
	case "*":
		return af * bf, stackOp, nil
	case "/":
		return af / bf, stackOp, nil
	}

	return 0, stackOp, fmt.Errorf("unexpected operator")
}