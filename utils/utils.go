package utils

import (
	"fmt"
	"strconv"
)

func IsDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func IsFloat(s string) bool {
	const bitSize = 64 // Don't think about it to much. It's just 64 bits.
	_, err := strconv.ParseFloat(s, bitSize)
	if err == nil {
		return true
	}
	return false
}

func Prior(c string) int {
	switch c {
	case "(":
		return 1
	case "+", "-":
		return 2
	case "*", "/":
		return 3
	default:
		return 0
	}
}

func IsContainsFunc(x string) bool {
	arr := []string{"*", "/", "+", "-"}
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}

func DebugLog(currentSymbol string, stack []string, queue []string) {
	fmt.Println("current symbol")
	fmt.Println(currentSymbol)
	fmt.Println("stack")
	fmt.Println(stack)
	fmt.Println("queue")
	fmt.Println(queue)
}
