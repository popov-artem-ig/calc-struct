package converter

import (
	"calc-struct/queue"
	"calc-struct/stack"
	"calc-struct/utils"
	"fmt"
)

func pop(s []string) (string, []string) {
	return s[len(s)-1], s[:len(s)-1]
}

func top(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func Convert(str string) ([]string, error) {
	np := 0                   // Инкремент открытых скобок
	var stackOp []string      // Стек операторов
	var resQueue []string     // Результирующая очередь
	strIn := str	          // Выражение для разбора
	hoarder := ""   		  // Накопитель для сборки/обработки дробных и больших чисел
	var v string

	for i := range strIn {
		c := strIn[i]
		//utils.DebugLog(string(c), stackOp, resQueue)

		if utils.IsDigit(c) || c == '.' {
			hoarder += string(c)
			continue
		}

		if len(hoarder) > 0 {
			resQueue = queue.Enqueue(hoarder, resQueue)
			hoarder = ""
		}

		switch c {
		case '(':
			{
				if i == len(strIn)-1 {
					return []string{}, fmt.Errorf("syntax error. opening parenthesis at the end of an expression")
				}
				stackOp = stack.Push(string(c), stackOp)
				np++
			}
		case '*', '/', '+', '-':
			{
				if i == len(strIn)-1 {
					return []string{}, fmt.Errorf("syntax error. operator at the end of an expression")
				}

					if stack.Peek(stackOp) == "" { //Если стек операторов пуст, алгоритм помещает входящий оператор в стек.
						stackOp = stack.Push(string(c), stackOp)
					} else {
						//Если приоритет входящего оператора ниже,
						//верхний оператор извлекается из стека и выводится в очередь,
						//после чего входящий оператор сравнивается с новой вершиной стека.
						for utils.Prior(string(c)) < utils.Prior(stack.Peek(stackOp)) {
							v, stackOp = stack.Pop(stackOp)
							resQueue = queue.Enqueue(v, resQueue)
						}
						//Если входящий оператор имеет более высокий приоритет,
						//чем тот оператор, что в настоящее время находится на вершине стека,
						//входящий оператор помещается на вершину стека.
						if utils.Prior(string(c)) > utils.Prior(stack.Peek(stackOp)) {
							stackOp = stack.Push(string(c), stackOp)
						} else if utils.Prior(string(c)) == utils.Prior(stack.Peek(stackOp)) {
							//Если входящий оператор имеет такой же приоритет,
							//верхний оператор извлекается из стека и выводится в очередь,
							//а входящий оператор помещается в стек.
							v, stackOp = stack.Pop(stackOp)
							resQueue = queue.Enqueue(v, resQueue)
							stackOp = stack.Push(string(c), stackOp)
						}
					}
			}
		case ')':
			{
				// До тех пор, пока верхним элементом стека не станет открывающая скобка, выталкиваем элементы из стека в выходную строку.
				for stack.Peek(stackOp) != "(" && stack.Peek(stackOp) != "" {
					v, stackOp = stack.Pop(stackOp)
					resQueue = queue.Enqueue(v, resQueue)
				}
				// Если стек закончился раньше, чем мы встретили открывающую скобку, это означает, что в выражении либо неверно поставлен разделитель,
				// либо не согласованы скобки.
				if stack.Peek(stackOp) == "" {
					return []string{}, fmt.Errorf("syntax error. inconsistent parentheses")
				}
				// При этом открывающая скобка удаляется из стека, но в выходную строку не добавляется.
				if stack.Peek(stackOp) == "(" {
					v, stackOp = stack.Pop(stackOp)
					np--
				}
				// Если после этого шага на вершине стека оказывается символ функции, выталкиваем его в выходную строку.
				if utils.IsContainsFunc(stack.Peek(stackOp)) {
					v, stackOp = stack.Pop(stackOp)
					resQueue = queue.Enqueue(v, resQueue)
				}
			}
		/*default:
			return "", fmt.Errorf("syntax error")*/
		}
	}
	for stack.Peek(stackOp) != "" {
		v, stackOp = stack.Pop(stackOp)
		resQueue = queue.Enqueue(v, resQueue)
	}
	/*if np > 0 {
		return "", fmt.Errorf("syntax error")
	}*/

	return resQueue, nil
}
