package converter

import (
	"calc-struct/utils"
	"fmt"
)

func prior(c uint8) int {
	switch c {
	case '(':
		return 1
	case '+', '-':
		return 2
	case '*', '/':
		return 3
	default:
		return 0
	}
}

func pop(s string) (uint8, string) {
	return s[len(s)-1], s[:len(s)-1]
}

func top(s string) uint8 {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1]
}

func Convert(str string) (string, error) {
	was_op := false
	np := 0
	stack := ""
	strIn := str
	strOut := ""
	var v uint8

	for i := range strIn {
		c := strIn[i]
		if utils.Is_digit(c) {
			strOut += string(c)
			was_op = false
			continue
		} else if c != ' ' {
			strOut += " "
		}
		switch c {
		case ' ':
			{
			}
		case '(':
			{
				stack += string(c)
				np++
				was_op = false
			}
		case '*', '/', '+', '-':
			{
				if i == len(strIn)-1 {
					return "", fmt.Errorf("syntax error")
				}
				if !was_op {
					was_op = true
					for prior(c) <= prior(top(stack)) {
						v, stack = pop(stack)
						strOut += string(v)
					}
					if prior(c) > prior(top(stack)) {
						stack += string(c)
					}
					break
				} else {
					return "", fmt.Errorf("syntax error")
				}
			}
		case ')':
			{
				if was_op {
					return "", fmt.Errorf("syntax error")
				} else {
					v, stack = pop(stack)
					for v != '(' && np > 0 {
						strOut += string(v)
						v, stack = pop(stack)
					}
				}
				np--
			}
		default:
			return "", fmt.Errorf("syntax error")
		}
	}
	for i := range stack {
		strOut += " " + string(stack[i])
	}
	if np > 0 {
		return "", fmt.Errorf("syntax error")
	}

	return strOut, nil

}
