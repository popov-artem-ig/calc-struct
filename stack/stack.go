package stack

func Push(s string, stack string) string {
	res := stack + s
	return res
}

func Pop(stack string) (uint8, string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func Peek(stack string) uint8 {
	if len(stack) == 0 {
		return 0
	}
	return stack[len(stack)-1]
}

func Count(stack string) int {
	return len(stack)
}